package log

import (
	. "design/interfaces"
	e "design/myError"
	"design/util"
	"os"
	"sync"
)

var once sync.Once

type Log struct {
}

func (l *Log) Update(command Command) error {
	var callSelf string
	if command == nil {
		callSelf = "error"
	}
	callSelf = command.CallSelf()

	// global variable of logger
	f, err := os.OpenFile("./logFiles/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return e.NewMyError("open log error")
	}
	defer func() {
		_ = f.Close()
	}()

	once.Do(func() {
		_ = util.Output("session start at "+util.GetNow()+"\n", f)
	})

	_ = util.Output(util.GetNow()+" "+callSelf+"\n", f)
	return nil
}
