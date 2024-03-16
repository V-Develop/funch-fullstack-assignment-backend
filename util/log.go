package util

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

type Loger struct {
	CorrelationId string
	LoginId       string
	Service       string
}

const (
	InfoLevel  = "INFO"
	DebugLevel = "DEBUG"
	ErrorLevel = "ERROR"
)

type LogModel struct {
	CorrelationId string    `json:"correlation_id"`
	LoginId       string    `json:"login_id"`
	LogLevel      string    `json:"log_level"`
	Service       string    `json:"service"`
	Message       string    `json:"message"`
	DateTime      time.Time `json:"date_time"`
}

func NewLog(loginId string, service string) Loger {
	logControllerPointer := Loger{
		CorrelationId: uuid.New().String(),
		LoginId:       loginId,
		Service:       service,
	}
	return logControllerPointer
}

func (l Loger) LogInfo(message string, loginId string) {
	var log = LogModel{
		CorrelationId: l.CorrelationId,
		LoginId:       l.LoginId,
		LogLevel:      InfoLevel,
		Service:       l.Service,
		Message:       message,
		DateTime:      time.Now(),
	}
	if loginId != "0" {
		log.LoginId = loginId
	}
	line := fmt.Sprintf("%s %s %s %s %s %s", log.DateTime, log.CorrelationId, log.Service, log.LogLevel, log.Message, log.LoginId)
	l.AddLog(line)
}

func (l Loger) LogDebug(message string, loginId string) {
	var log = LogModel{
		CorrelationId: l.CorrelationId,
		LoginId:       l.LoginId,
		LogLevel:      DebugLevel,
		Service:       l.Service,
		Message:       message,
		DateTime:      time.Now(),
	}
	if loginId != "0" {
		log.LoginId = loginId
	}
	line := fmt.Sprintf("%s %s %s %s %s %s", log.DateTime, log.CorrelationId, log.Service, log.LogLevel, log.Message, log.LoginId)
	l.AddLog(line)
}

func (l Loger) LogError(message string, loginId string) {
	var log = LogModel{
		CorrelationId: l.CorrelationId,
		LoginId:       l.LoginId,
		LogLevel:      ErrorLevel,
		Service:       l.Service,
		Message:       message,
		DateTime:      time.Now(),
	}
	if loginId != "0" {
		log.LoginId = loginId
	}
	line := fmt.Sprintf("%s %s %s %s %s %s", log.DateTime, log.CorrelationId, log.Service, log.LogLevel, log.Message, log.LoginId)
	l.AddLog(line)
}

func (l Loger) LogInfof(format string, args interface{}, loginId string) {
	message := fmt.Sprintf(format, args)
	var log = LogModel{
		CorrelationId: l.CorrelationId,
		LoginId:       l.LoginId,
		LogLevel:      InfoLevel,
		Service:       l.Service,
		Message:       message,
		DateTime:      time.Now(),
	}
	if loginId != "0" {
		log.LoginId = loginId
	}
	line := fmt.Sprintf("%s %s %s %s %s %s", log.DateTime, log.CorrelationId, log.Service, log.LogLevel, log.Message, log.LoginId)
	l.AddLog(line)
}

func (l Loger) LogDebugf(format string, args interface{}, loginId string) {
	message := fmt.Sprintf(format, args)
	var log = LogModel{
		CorrelationId: l.CorrelationId,
		LoginId:       l.LoginId,
		LogLevel:      DebugLevel,
		Service:       l.Service,
		Message:       message,
		DateTime:      time.Now(),
	}
	if loginId != "0" {
		log.LoginId = loginId
	}
	line := fmt.Sprintf("%s %s %s %s %s %s", log.DateTime, log.CorrelationId, log.Service, log.LogLevel, log.Message, log.LoginId)
	l.AddLog(line)
}

func (l Loger) LogDebugMoreArgsf(loginId string, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	var log = LogModel{
		CorrelationId: l.CorrelationId,
		LoginId:       l.LoginId,
		LogLevel:      DebugLevel,
		Service:       l.Service,
		Message:       message,
		DateTime:      time.Now(),
	}
	if loginId != "0" {
		log.LoginId = loginId
	}
	line := fmt.Sprintf("%s %s %s %s %s %s", log.DateTime, log.CorrelationId, log.Service, log.LogLevel, log.Message, log.LoginId)
	l.AddLog(line)
}

func (l Loger) LogErrorf(format string, args interface{}, loginId string) {
	message := fmt.Sprintf(format, args)
	var log = LogModel{
		CorrelationId: l.CorrelationId,
		LoginId:       l.LoginId,
		LogLevel:      ErrorLevel,
		Service:       l.Service,
		Message:       message,
		DateTime:      time.Now(),
	}
	if loginId != "0" {
		log.LoginId = loginId
	}
	line := fmt.Sprintf("%s %s %s %s %s %s", log.DateTime, log.CorrelationId, log.Service, log.LogLevel, log.Message, log.LoginId)
	l.AddLog(line)
}

func (l Loger) AddLog(message string) {
	const layout = "2006-01-02"
	logFileName := "temp/" + time.Now().Format(layout) + ".log"
	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(message + "\n"); err != nil {
		panic(err)
	}
}
