package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type streamLogger struct {
	stream  io.WriteCloser
	linebuf []byte
	lock    sync.Mutex
}

func NewStdoutLogger() Interface {
	return &streamLogger{
		linebuf: []byte{},
		stream:  noopCloser{os.Stdout},
	}
}

func NewStderrLogger() Interface {
	return &streamLogger{
		linebuf: []byte{},
		stream:  noopCloser{os.Stderr},
	}
}

func (w *streamLogger) output(timeStr, level, msg string, logCategory string, UUID string) {

	w.lock.Lock()
	w.linebuf = w.linebuf[:0]

	w.linebuf = append(w.linebuf, timeStr...)
	w.linebuf = append(w.linebuf, []byte(" [")...)
	w.linebuf = append(w.linebuf, UUID...)
	w.linebuf = append(w.linebuf, []byte("]")...)

	w.linebuf = append(w.linebuf, '[')
	w.linebuf = append(w.linebuf, level...)
	w.linebuf = append(w.linebuf, ']')

	w.linebuf = append(w.linebuf, '[')
	w.linebuf = append(w.linebuf, logCategory...)
	w.linebuf = append(w.linebuf, ']')

	w.linebuf = append(w.linebuf, ' ')
	w.linebuf = append(w.linebuf, msg...)

	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		w.linebuf = append(w.linebuf, '\n')
	}

	w.stream.Write(w.linebuf)
	w.lock.Unlock()
}

func (w *streamLogger) Tracef(logCategory string, UUID string, f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "trace", msg, logCategory, UUID)
}

func (w *streamLogger) Traceln(logCategory string, UUID string, message string) {
	msg := message
	now := time.Now().Format(timeFormatStr)
	w.output(now, "trace", msg, logCategory, UUID)
}

func (w *streamLogger) Infof(logCategory string, UUID string, f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "info", msg, logCategory, UUID)
}

func (w *streamLogger) Infoln(logCategory string, UUID string, message string) {
	msg := message
	now := time.Now().Format(timeFormatStr)
	w.output(now, "info", msg, logCategory, UUID)
}

func (w *streamLogger) Debugf(logCategory string, UUID string, f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "debug", msg, logCategory, UUID)
}

func (w *streamLogger) Debugln(logCategory string, UUID string, v ...interface{}) {
	msg := fmt.Sprintln(v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "debug", msg, logCategory, UUID)
}

func (w *streamLogger) Warningf(logCategory string, UUID string, f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "warning", msg, logCategory, UUID)
}

func (w *streamLogger) Warningln(logCategory string, UUID string, v ...interface{}) {
	msg := fmt.Sprintln(v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "warning", msg, logCategory, UUID)
}

func (w *streamLogger) Errorf(logCategory string, UUID string, f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "error", msg, logCategory, UUID)
}

func (w *streamLogger) Errorln(logCategory string, UUID string, v ...interface{}) {
	msg := fmt.Sprintln(v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "error", msg, logCategory, UUID)
}

func (w *streamLogger) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.stream.Close()
}
