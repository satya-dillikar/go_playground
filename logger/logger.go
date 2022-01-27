package main

import (
	"context"
	"log"
	"os"

	klog "k8s.io/klog/v2"
)

var logger = log.New(os.Stdout, "TEST-LOG: ", 5)

func main() {
	/* 	logger.Println("Hello, playground")
	   	// or
	   	log.Println("Hello, playground")

	   	klog.Errorf("Hello with %s", "formatting")
	   	klog.Error("Hello with %s", "formatting")

	   	klog.Infof("Hello with %s", "formatting")
	   	klog.Info("Hello with %s", "formatting")

	   	var logger Logger = NewBuiltinLogger()
	   	logger.Debugf("Hello with %s", "formatting")
	   	logger.Debug("Hello with %s", "formatting")

	   	logger.Infof("Hello with %s", "formatting")
	   	logger.Info("Hello with %s", "formatting")

	   	var klogger Logger = NewBuiltinKlogger()
	   	klogger.Debugf("Hello with %s", "formatting")
	   	klogger.Debug("Hello with %s", "formatting")

	   	klogger.Infof("Hello with %s", "formatting")
	   	klogger.Info("Hello with %s", "formatting")
	*/
	ctx := context.WithValue(context.Background(), "logger", NewBuiltinKlogger())
	DoStuff(ctx)
}

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}

type BuiltinLogger struct {
	logger *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "TEST-LOG: ", 5)}
}

func (l *BuiltinLogger) Debug(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
func (l *BuiltinLogger) Info(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Infof(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

type BuiltinKlogger struct {
	//logger Logger
}

func NewBuiltinKlogger() *BuiltinKlogger {
	return nil
}

func (l *BuiltinKlogger) Debug(args ...interface{}) {
	klog.Info(args...)
}

func (l *BuiltinKlogger) Debugf(format string, args ...interface{}) {
	klog.Infof(format, args...)
}
func (l *BuiltinKlogger) Info(args ...interface{}) {
	klog.Info(args...)
}

func (l *BuiltinKlogger) Infof(format string, args ...interface{}) {
	klog.Infof(format, args...)
}

func DoStuff(ctx context.Context) {
	log := ctx.Value("logger").(Logger)
	log.Debugf("DoStuff: Hello with %s", "formatting")
}
