package model

import (
	project "github.com/olefen/note/iris/mvc"
	"github.com/suboat/sorm"
)

var (
	ModelUser orm.Model
)

func Init(modelUser orm.Model) (err error) {
	if err = modelUser.Ensure(&project.Admin{}); err != nil {
		return
	}
	ModelUser = modelUser
	return
}
