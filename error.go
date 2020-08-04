package note

import "errors"

var (
	ErrKeyIsNull = errors.New("input Key arguments error") //   dssad
	ErrPubKeyRsa = errors.New("public rasKey error")
	ErrPriKerRsa = errors.New("private rasKey error")
)
