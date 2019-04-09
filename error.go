package note

import "errors"

var (
	ErrKeyIsNull = errors.New("input Key arguments error") //
	ErrPubKeyRsa = errors.New("public rasKey error")
	ErrPriKerRsa = errors.New("private rasKey error")
)
