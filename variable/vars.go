package variable

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strconv"
	"strings"
)

var (
	Log SLoger
)

type SLoger struct {
	log  *log.Logger
	isIO bool
}

func (this *SLoger) IsIO(is bool) {
	this.isIO = is
}

func (this *SLoger) New(out io.Writer, prefix string, flag int, isio bool) {
	this.log = log.New(out, "", flag)
	this.isIO = isio
}

func (this *SLoger) GetFile() string {
	_, file, line, _ := runtime.Caller(2)
	file = file[strings.LastIndex(file, "/")+1:]
	str := file + ":" + strconv.Itoa(line) + ":"
	return str
}

func (this *SLoger) Println(v ...interface{}) {
	if !this.isIO {
		return
	}

	list := [](interface{}){}
	list = append(list, this.GetFile())
	for index := 0; index < len(v); index++ {
		list = append(list, v[index])
	}
	this.log.Println(list)
}

func (this *SLoger) PrintlnShow(v ...interface{}) {
	list := [](interface{}){}
	list = append(list, this.GetFile())
	for index := 0; index < len(v); index++ {
		list = append(list, v[index])
	}
	this.log.Println(list)
}

func (this *SLoger) Printf(format string, v ...interface{}) {
	if !this.isIO {
		return
	}
	format = this.GetFile() + format
	this.log.Printf(format, v...)
}

func (this *SLoger) Fatalf(format string, v ...interface{}) {
	if !this.isIO {
		return
	}
	format = this.GetFile() + format
	this.log.Fatalf(format, v...)
}

func (this *SLoger) Fatal(v ...interface{}) {
	if !this.isIO {
		return
	}
	list := [](interface{}){}
	list = append(list, this.GetFile())
	for index := 0; index < len(v); index++ {
		list = append(list, v[index])
	}
	this.log.Fatal(list)
}

// PanicRecoverError 统一处理panic, 并更新error
func PanicRecoverError(l SLoger) {
	r := recover()
	if r == nil {
		return
	}
	l.Println(fmt.Errorf(`[panic-recover] "%s" %v`, panicIdentify(), r))

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
