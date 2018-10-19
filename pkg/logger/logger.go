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
	"fmt"
	"io"
	"log"
)

const (
	black  = "0;30"
	red    = "0;31"
	green  = "0;32"
	yellow = "0;33"
	cyan   = "0;36"
)

type Logger interface {
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Print(v ...interface{})
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	trace   *log.Logger
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	print   *log.Logger
}

type Config struct {
	TraceWriter   io.Writer
	DebugWriter   io.Writer
	InfoWriter    io.Writer
	WarningWriter io.Writer
	ErrorWriter   io.Writer
	PrintWriter   io.Writer
	NoColor       bool
}

func NewLogger(config Config) (Logger) {
	flag := log.Ldate | log.Lmicroseconds
	return &logger{
		trace:   log.New(config.TraceWriter, color(black, padRight("[TRACE]"), config.NoColor), flag),
		debug:   log.New(config.DebugWriter, color(cyan, padRight("[DEBUG]"), config.NoColor), flag),
		info:    log.New(config.InfoWriter, color(green, padRight("[INFO]"), config.NoColor), flag),
		warning: log.New(config.WarningWriter, color(yellow, padRight("[WARN]"), config.NoColor), flag),
		error:   log.New(config.ErrorWriter, color(red, padRight("[ERROR]"), config.NoColor), flag),
		print:   log.New(config.PrintWriter, "", 0),
	}
}

func (log *logger) Tracef(format string, v ...interface{}) {
	log.Trace(fmt.Sprintf(format, v...))
}

func (log *logger) Debugf(format string, v ...interface{}) {
	log.Debug(fmt.Sprintf(format, v...))
}

func (log *logger) Infof(format string, v ...interface{}) {
	log.Info(fmt.Sprintf(format, v...))
}

func (log *logger) Warningf(format string, v ...interface{}) {
	log.Warning(fmt.Sprintf(format, v...))
}

func (log *logger) Errorf(format string, v ...interface{}) {
	log.Error(fmt.Sprintf(format, v...))
}

func (log *logger) Printf(format string, v ...interface{}) {
	log.Print(fmt.Sprintf(format, v...))
}

func (log *logger) Trace(v ...interface{}) {
	log.trace.Println(v...)
}

func (log *logger) Debug(v ...interface{}) {
	log.debug.Println(v...)
}

func (log *logger) Info(v ...interface{}) {
	log.info.Println(v...)
}

func (log *logger) Warning(v ...interface{}) {
	log.warning.Println(v...)
}

func (log *logger) Error(v ...interface{}) {
	log.error.Println(v...)
}

func (log *logger) Print(v ...interface{}) {
	log.print.Println(v...)
}

func color(code string, str string, noColor bool) string {
	if noColor {
		return str
	}
	return fmt.Sprintf("\033[%sm%s\033[m", code, str)
}

func padRight(str string) string {
	return fmt.Sprintf("%-8s", str)
}
