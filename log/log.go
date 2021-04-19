/*
 * Tencent is pleased to support the open source community by making BK-CI 蓝鲸持续集成平台 available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company.  All rights reserved.
 *
 * BK-CI 蓝鲸持续集成平台 is licensed under the MIT license.
 *
 * A copy of the MIT License is included in this file.
 *
 *
 * Terms of the MIT License:
 * ---------------------------------------------------
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation
 * files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy,
 * modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT
 * LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
 * NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

 package log

 import (
	 "fmt"
	 "sync"
 )
 
 var lock = new(sync.Mutex)
 
 // Info 日志
 func Info(v ...interface{}) {
	 lock.Lock()
	 fmt.Print("##[info]")
	 fmt.Println(v...)
	 lock.Unlock()
 }
 
 // Warn 日志
 func Warn(v ...interface{}) {
	 lock.Lock()
	 fmt.Print("##[warning]")
	 fmt.Println(v...)
	 lock.Unlock()
 }
 
 // Error 日志
 func Error(v ...interface{}) {
	 lock.Lock()
	 fmt.Print("##[error]")
	 fmt.Println(v...)
	 lock.Unlock()
 }
 
 // Debug 日志
 func Debug(v ...interface{}) {
	 lock.Lock()
	 fmt.Print("##[debug]")
	 fmt.Println(v...)
	 lock.Unlock()
 }
 
 // Command 日志
 func Command(v ...interface{}) {
	 lock.Lock()
	 fmt.Print("##[command]")
	 fmt.Println(v...)
	 lock.Unlock()
 }
 
 // Group 分组日志开始
 func Group(v ...interface{}) {
	 lock.Lock()
	 fmt.Print("##[group]")
	 fmt.Println(v...)
	 lock.Unlock()
 }
 
 // EndGroup 分组日志结束
 func EndGroup(v ...interface{}) {
	 lock.Lock()
	 fmt.Print("##[endgroup]")
	 fmt.Println(v)
	 lock.Unlock()
 }
 