package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type LogWithCollector struct {
	*zap.SugaredLogger
	collectedLogs []string
}

func (l *LogWithCollector) AddLog(template string, args ...interface{}) {
	l.collectedLogs = append(l.collectedLogs, fmt.Sprintf(template, args...))
}

func (l *LogWithCollector) AddLogs(logs *[]string) {
	l.collectedLogs = append(l.collectedLogs, *logs...)
}

func (l *LogWithCollector) clearCollectedLogs() {
	l.collectedLogs = []string{}
}

func (l *LogWithCollector) GetCollectedLogs() *[]string {
	return &l.collectedLogs
}

// ---
func (l *LogWithCollector) Debugf(template string, args ...interface{}) {
	l.SugaredLogger.Debugf(template, args...)
	l.AddLog(template, args...)
}

func (l *LogWithCollector) Infof(template string, args ...interface{}) {
	l.SugaredLogger.Infof(template, args...)
	l.AddLog(template, args...)
}

func (l *LogWithCollector) Warnf(template string, args ...interface{}) {
	l.SugaredLogger.Warnf(template, args...)
	l.AddLog(template, args...)
}

func (l *LogWithCollector) Errorf(template string, args ...interface{}) {
	l.SugaredLogger.Errorf(template, args...)
	l.AddLog(template, args...)
}

func WithLogCollector[T any](logger *zap.SugaredLogger, block func(*LogWithCollector) T) (T, []string) {
	collector := NewLogWithCollector(logger)
	result := block(collector)
	logs := make([]string, len(collector.collectedLogs))
	copy(logs, collector.collectedLogs)
	collector.clearCollectedLogs()
	return result, logs
}

func NewLogWithCollector(logger *zap.SugaredLogger) *LogWithCollector {
	return &LogWithCollector{
		SugaredLogger: logger,
	}
}
