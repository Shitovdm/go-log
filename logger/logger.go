package logger

const (
	timeFormatStr = "2006-01-02 15:04:05"
)

var (
	stdoutLogger = NewStdoutLogger()
	stderrLogger = NewStderrLogger()
)

type Logger struct {
	UUID string
}

type Interface interface {
	Tracef(category string, UUID string, format string, v ...interface{})
	Traceln(category string, UUID string, message string)

	Infof(category string, UUID string, format string, v ...interface{})
	Infoln(category string, UUID string, message string)

	Debugf(category string, UUID string, format string, v ...interface{})
	Debugln(category string, UUID string, v ...interface{})

	Warningf(category string, UUID string, format string, v ...interface{})
	Warningln(category string, UUID string, v ...interface{})

	Errorf(category string, UUID string, format string, v ...interface{})
	Errorln(category string, UUID string, v ...interface{})

	Close() error
}

func NewLoggerInstance() *Logger {
	l := &Logger{}
	l.UUID = GenerateNewSessionUuid()
	return l
}

func (l *Logger) Trace(message, category string) {
	stdoutLogger.Traceln(category, l.UUID, message)
}

func (l *Logger) Tracef(category, format string, v ...interface{}) {
	stdoutLogger.Tracef(category, l.UUID, format, v...)
}

func (l *Logger) Info(message, category string) {
	stdoutLogger.Infoln(category, l.UUID, message)
}

func (l *Logger) Infof(category, format string, v ...interface{}) {
	stdoutLogger.Infof(category, l.UUID, format, v...)
}

func (l *Logger) Debug(message, category string) {
	stdoutLogger.Debugln(category, l.UUID, message)
}

func (l *Logger) Debugf(category, format string, v ...interface{}) {
	stdoutLogger.Debugf(category, l.UUID, format, v...)
}

func (l *Logger) Warning(message, category string) {
	stdoutLogger.Warningln(category, l.UUID, message)
}

func (l *Logger) Warningf(category, format string, v ...interface{}) {
	stdoutLogger.Warningf(category, l.UUID, format, v...)
}

func (l *Logger) Error(message, category string) {
	stderrLogger.Errorln(category, l.UUID, message)
}

func (l *Logger) Errorf(category, format string, v ...interface{}) {
	stderrLogger.Errorf(category, l.UUID, format, v...)
}
