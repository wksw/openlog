package openlog

import (
	"log"
)

const (
	i = "INFO: "
	d = "DEBUG: "
	e = "ERROR: "
	w = "WARN: "
	f = "FATAL: "
)

type golog struct {
}

func (l golog) Debug(message string, opts ...Option) {
	log.Println(d + message)
}

func (l golog) Debugf(format string, args ...interface{}) {
	log.SetPrefix(d)
	log.Printf(format, args...)
}

func (l golog) Info(message string, opts ...Option) {
	log.Println(i + message)
}

func (l golog) Infof(format string, args ...interface{}) {
	log.SetPrefix(i)
	log.Printf(format, args...)
}

func (l golog) Warn(message string, opts ...Option) {
	log.Println(w + message)
}

func (l golog) Warnf(format string, args ...interface{}) {
	log.SetPrefix(w)
	log.Printf(format, args...)
}

func (l golog) Error(message string, opts ...Option) {
	log.Println(e + message)
}

func (l golog) Errorf(format string, args ...interface{}) {
	log.SetPrefix(e)
	log.Printf(format, args...)
}

func (l golog) Fatal(message string, opts ...Option) {
	log.Panic(f + message)
}

func (l golog) Fatalf(format string, args ...interface{}) {
	log.SetPrefix(e)
	log.Fatalf(format, args...)
}
