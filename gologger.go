package gologger

import (
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	Name string `json:"name"`
}

func NewLogger(name string) *Logger {
	return &Logger{Name: name}
}
func GetTime() string {
	currentTime := time.Now()
	return currentTime.Format("15:04:05.000000")
}

func (l *Logger) Green(msg string, a ...interface{}) {
	color.Green("["+GetTime()+"] "+"["+l.Name+"] [+] "+msg, a...)
}

func (l *Logger) Blue(msg string, a ...interface{}) {
	color.Blue("["+GetTime()+"] "+"["+l.Name+"] [$] "+msg, a...)
}

func (l *Logger) Red(msg string, a ...interface{}) {
	color.Red("["+GetTime()+"] "+"["+l.Name+"] [x] "+msg, a...)
}

func (l *Logger) Yellow(msg string, a ...interface{}) {
	color.Yellow("["+GetTime()+"] "+"["+l.Name+"] [?] "+msg, a...)
}

func (l *Logger) Normal(msg string, a ...interface{}) {
	color.White("["+GetTime()+"] "+"["+l.Name+"] [#] "+msg, a...)
}
