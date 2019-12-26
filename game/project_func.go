package game

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

func NewUser(in *User) (out *User) {
	if in != nil {
		out = in
	} else {
		out = new(User)
	}
	out.Uid = uuid.NewV4().String()
	if out.CreatedAt.Unix() <= 0 {
		out.CreatedAt = time.Now()
	}
	if out.UpdatedAt.Unix() <= 0 {
		out.UpdatedAt = time.Now()
	}
	return
}
