package main

import (
	"git.yichui.net/tudy/go-rest"
	"git.yichui.net/tudy/go-rest/log"
	"runtime"
)

var LogPanic rest.Logger

func main() {

	if err := test2(10, 2); err != nil {
		println(err.Error())
	}

	/*	if err := test2(11, 2); err != nil {
		panic(err)
	}*/
}

// PanicRecoverError 统一处理panic, 并更新error
func PanicRecoverError(logger rest.Logger, err *error) {
	buf := make([]byte, 64<<10)
	if logger == nil && LogPanic != nil {
		logger = LogPanic
	} else {
		//logger = log.Log
	}
	r := recover()
	if r != nil {
		buf = buf[:runtime.Stack(buf, false)]
		logger.Errorf(`[panic-recover] %s,%v`, string(buf), r)
	} else {
		return
	}

	return
}

func test1(a, b int) (err error) {
	defer PanicRecoverError(LogPanic, &err)
	var s *string
	println(len(*s))
	return
}

func test2(a, b int) (err error) {
	defer rest.PanicRecoverError(LogPanic, &err)
	var s *string
	println(len(*s))
	return
}

func init() {
	LogPanic, _ = log.NewLog(nil)
}
