/*
 * Copyright (c) 2018 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package logger

import (
	"os"
)

var (
	globalLogger = NewLogger(Config{
		TraceWriter:   os.Stderr,
		DebugWriter:   os.Stderr,
		InfoWriter:    os.Stderr,
		WarningWriter: os.Stderr,
		ErrorWriter:   os.Stderr,
		PrintWriter:   os.Stdout,
		NoColor:       false,
	})
)

func Tracef(format string, v ...interface{}) {
	globalLogger.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	globalLogger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	globalLogger.Infof(format, v...)
}

func Warningf(format string, v ...interface{}) {
	globalLogger.Warningf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	globalLogger.Errorf(format, v...)
}

func Printf(format string, v ...interface{}) {
	globalLogger.Printf(format, v...)
}

func Trace(v ...interface{}) {
	globalLogger.Trace(v...)
}

func Debug(v ...interface{}) {
	globalLogger.Debug(v...)
}

func Info(v ...interface{}) {
	globalLogger.Info(v...)
}

func Warning(v ...interface{}) {
	globalLogger.Warning(v...)
}

func Error(v ...interface{}) {
	globalLogger.Error(v...)
}

func Print(v ...interface{}) {
	globalLogger.Print(v...)
}

func SetGlobalLogger(logger Logger) {
	globalLogger = logger
}
