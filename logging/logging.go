package logging

import "go.uber.org/zap"

func GetSugar() (l *zap.SugaredLogger) {
	log := GetLog()
	l = log.Sugar()
	return
}

func GetLog() (l *zap.Logger) {
	l, _ = zap.NewDevelopment()
	return
}
