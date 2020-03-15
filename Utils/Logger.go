package Utils

import (
	"GoFly/Model"
	"fmt"
	"log"
	"time"
)

func Log(l *log.Logger, obj Model.NLPResponse) {
	// Print log from NLPResponse object
	msg := fmt.Sprintf("%#v\n", obj)
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [SERVER] ")
	l.Print(msg)
}
