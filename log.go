// Provides a Log interface that can be used to wrap other logging
// interfaces to allow swapping out log interfaces.
//
// Log to stdout:
//
//    log := NewStdoutLog()
//    log.Error("My name is: %s", name)
//
// Log using syslog:
//
//    log, err := log.NewLog("myapp")
//    if err != nil {
//        fmt.Fprintln(os.Stderr, "Failure to setup syslog logging: %v", err)
//        os.Exit(1)
//    }
//    log.Error("My name is: %s", name)
//
package log

import (
	"fmt"
	"log/syslog"
	"os"
)

type Log interface {
	Debug(format string, a ...interface{}) error
	Info(format string, a ...interface{}) error
	Notice(format string, a ...interface{}) error
	Warning(format string, a ...interface{}) error
	Error(format string, a ...interface{}) error
	Close()
}

type SyslogLog struct {
	l *syslog.Writer
}

func NewLog(label string) (Log, error) {
	log, err := syslog.New(syslog.LOG_ERR, label)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error establishing syslog. %v", err)
		return nil, err
	}
	return &SyslogLog{l: log}, nil
}

func (l *SyslogLog) Close() {
	l.l.Close()
}

func (l *SyslogLog) Debug(format string, a ...interface{}) error {
	s := fmt.Sprintf("DEBUG: "+format, a...)
	return l.l.Debug(s)
}

func (l *SyslogLog) Info(format string, a ...interface{}) error {
	s := fmt.Sprintf("INFO: "+format, a...)
	return l.l.Info(s)
}

func (l *SyslogLog) Notice(format string, a ...interface{}) error {
	s := fmt.Sprintf("NOTICE: "+format, a...)
	return l.l.Notice(s)
}

func (l *SyslogLog) Warning(format string, a ...interface{}) error {
	s := fmt.Sprintf("WARNING: "+format, a...)
	return l.l.Warning(s)
}

func (l *SyslogLog) Error(format string, a ...interface{}) error {
	s := fmt.Sprintf("ERROR: "+format, a...)
	return l.l.Err(s)
}

func NewStdoutLog() Log {
	return &StdoutLog{ShowDebug: false}
}

func NewStdoutLogDebug() Log {
	return &StdoutLog{ShowDebug: true}
}

type StdoutLog struct {
	ShowDebug bool
}

func (l *StdoutLog) Close() {
}

func (l *StdoutLog) Debug(format string, a ...interface{}) error {
	if l.ShowDebug {
		fmt.Printf("DEBUG: "+format+"\n", a...)
	}
	return nil
}

func (l *StdoutLog) Info(format string, a ...interface{}) error {
	fmt.Printf("INFO: "+format+"\n", a...)
	return nil
}

func (l *StdoutLog) Notice(format string, a ...interface{}) error {
	fmt.Printf("NOTICE: "+format+"\n", a...)
	return nil
}

func (l *StdoutLog) Warning(format string, a ...interface{}) error {
	fmt.Printf("WARNING: "+format+"\n", a...)
	return nil
}

func (l *StdoutLog) Error(format string, a ...interface{}) error {
	fmt.Printf("ERROR: "+format+"\n", a...)
	return nil
}
