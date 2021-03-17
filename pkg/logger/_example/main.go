package main

import (
	"errors"
	"time"

	"github.com/JohnnyWei188/gwframe/pkg/logger"
)

func main() {

	opts := logger.Filename("./log/server.log")
	log := logger.New(opts)
	log.Errorf("this is a error: %s ", errors.New("error"))
	log.Debugf("this is a debug: %s ", errors.New("debug"))

	//for i := 0; i < 1000; i++ {
	//	go log.With("i"+strconv.Itoa(i), i).Info("bingfa")
	//}

	l := log.With("x1", "x111", "x2", "x222")

	l = l.With("x3", "x333")
	l.Info("info success")
	//{"level":"INFO","time":"2020-11-27T10:46:02.866+0800","caller":"_example/main.go:30","msg":"info success","x1":"x111","x2":"x222","x3":"x333"}
	l.Error("info error")
	//{"level":"ERROR","time":"2020-11-27T10:46:02.866+0800","caller":"_example/main.go:30","msg":"info error","x1":"x111","x2":"x222","x3":"x333"}

	ll := logger.Instance()
	ll.Info("Instance")
	ll.Error("Error")

	time.Sleep(10 * time.Second)

}
