package variable

import (
	"fmt"
	"testing"
)

type aa struct {
	A string
}

func (a *aa) Test() {
	fmt.Println(a.A)
}

func TestPanicRecoverError(t *testing.T) {
	var (
		a *aa
	)
	PanicRecoverError(Log)
	a.Test()

}
