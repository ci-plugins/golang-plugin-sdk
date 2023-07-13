package api

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/ci-plugins/golang-plugin-sdk/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	DEFAULT_LANGUAGE_TYPE = "zh_CN"
	I18NFILE_PREFIX       = "message_"
	I18NFILE_SUFFIX       = ".properties"
	pattern               = `\{(\d+)\}`
)

var (
	localizer          *localizerType
	defaultLocalTag    = language.Make(DEFAULT_LANGUAGE_TYPE)
	runtimeLanguageEnv = "BK_CI_LOCALE_LANGUAGE"
)

type localizerType struct {
	nowLocalizer language.Tag
	rwLock       sync.RWMutex
	localizers   map[language.Tag]*i18n.Localizer
}

func (l *localizerType) getLocalizer() *i18n.Localizer {
	l.rwLock.RLock()
	defer l.rwLock.RUnlock()
	local, ok := l.localizers[l.nowLocalizer]
	if !ok {
		// 未找到对应的本地化时默认使用中文
		return l.localizers[defaultLocalTag]
	}
	return local
}

// GetRuntimeLanguage 获取当前运行时的插件语言
func GetRuntimeLanguage() string {
	l := os.Getenv(runtimeLanguageEnv)
	if l == "" {
		return DEFAULT_LANGUAGE_TYPE
	}
	return l
}

// Initi18n 初始化国际化，并切换为指定语言
// translations 生成的国际化数据，可配合 i18ngenerator 使用，机构为 [language]{key, value}
func InitI18n(translations map[string][][]string, nowLanguage string) error {
	localizers := map[language.Tag]*i18n.Localizer{}
	bundle := i18n.NewBundle(language.SimplifiedChinese)

	for languageTag, v := range translations {
		tag, err := language.Parse(languageTag)
		if err != nil {
			log.Warnf("parse translations tag %s param failed: %s", languageTag, err.Error())
			continue
		}

		var messages = []*i18n.Message{}
		for _, c := range v {
			messages = append(messages, &i18n.Message{
				ID:    c[0],
				Other: c[1],
			})
		}

		err = bundle.AddMessages(tag, messages...) 
		if err != nil {
			return err
		}
		localizers[tag] = i18n.NewLocalizer(bundle, tag.String())
	}

	localizer = &localizerType{
		// 初始化时默认为中文
		nowLocalizer: defaultLocalTag,
		rwLock:       sync.RWMutex{},
		localizers:   localizers,
	}

	ChangeLocalizer(nowLanguage)

	return nil
}

// ChangeLocalizer 切换国际化语言
func ChangeLocalizer(l string) {
	newLocal := language.Make(l)

	// 先用读锁看一眼，如果一样就不换了
	localizer.rwLock.RLock()
	if localizer.nowLocalizer == newLocal {
		localizer.rwLock.RUnlock()
		return
	}
	localizer.rwLock.RUnlock()

	localizer.rwLock.Lock()
	defer localizer.rwLock.Unlock()

	localizer.nowLocalizer = newLocal
}

// Localize 使用国际化内容
func Localize(messageId string, params ...interface{}) (string, error) {
	localizer.rwLock.RLock()
	defer localizer.rwLock.RUnlock()

	nowLocalizer := localizer.getLocalizer()
	if nowLocalizer == nil {
		return "", errors.New("no current internationalization language found")
	}

	translation, err := nowLocalizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageId,
	})
	if err != nil {
		return "", err
	}

	return format(translation, params...), nil
}

func format(message string, params ...interface{}) string {
	re := regexp.MustCompile(pattern)

	var keySet = make(map[string]int)
	matches := re.FindAllString(message, -1)
	for _, match := range matches {
		v, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(match, "{"), "}"))
		if err != nil {
			continue
		}
		keySet[match] = v
	}

	res := message
	for k, idx := range keySet {
		if idx > len(params)-1 || idx < 0 {
			continue
		}
		res = strings.ReplaceAll(res, k, fmt.Sprint(params[idx]))
	}

	return res
}
