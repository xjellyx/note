package main

import (
	"github.com/olefen/note/log"

	"fmt"
	"runtime"
	"strings"
)

var (
	// LogPanic 输出panic错误
	LogPanic = new(log.Logger)
)

type aa struct {
	L string
}

func (a *aa) Log() {
	fmt.Println(a.L)
}

func main() {
	var l = log.NewLogFile("sss.log")
	var err error
	defer PanicRecoverError(l, &err)
	ad()

}

func ad() {
	panic("sdwqdqw")
}

// PanicRecoverError 统一处理panic, 并更新error
func PanicRecoverError(l *log.Logger, err *error) {
	r := recover()
	if r == nil {
		return
	}
	*err = fmt.Errorf(`[panic-recover] "%s" %v`, panicIdentify(), r)
	if err != nil {
		l.Error(*err)
		return
	}
	return
}

//
func panicIdentify() string {
	var (
		pc [16]uintptr
		n  = runtime.Callers(3, pc[:])
	)

	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		_fnName := fn.Name()
		if strings.HasPrefix(_fnName, "runtime.") {
			continue
		}
		file, line := fn.FileLine(pc)
		//
		var (
			_fnNameDir = strings.Split(_fnName, "/")
			_fnNameLis = strings.Split(_fnName, ".")
			_fnNameSrc string
		)

		if len(_fnNameDir) > 1 {
			_fnNameSrc = _fnNameDir[0] + "/" + _fnNameDir[1] + "/"
		} else {
			_fnNameSrc = _fnNameDir[0]
		}
		fnName := _fnNameLis[len(_fnNameLis)-1]

		// file
		_pcLis := strings.Split(file, _fnNameSrc)
		filePath := strings.Join(_pcLis[:], "")
		return fmt.Sprintf("%s:%d|%s", filePath, line, fnName)
	}

	return "unknown"
}
