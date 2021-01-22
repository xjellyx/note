package main

import(
	"strconv"
	"github.com/olongfen/demo/app/model/user"
	"github.com/mitchellh/mapstructure"
)


















// AddUserReqForm
type AddUserReqForm struct {Name string `json:"name" form:"name"` // if required, add binding:"required" to tag by self
	Age int `json:"age" form:"age"` // if required, add binding:"required" to tag by self
	Class string `json:"class" form:"class"` // if required, add binding:"required" to tag by self
}

func (a *AddUserReqForm) Valid() (err error) {
	return
}


// EditUserReqForm
type EditUserReqForm struct {

	ID uint `json:"id" form:"id" binding:"required"`




	Name string `json:"name" form:"name"` // if required, add binding:"required" to tag by self

	Age int `json:"age" form:"age"` // if required, add binding:"required" to tag by self

	Class string `json:"class" form:"class"` // if required, add binding:"required" to tag by self
}

func (a *EditUserReqForm) Valid() (err error) {
	return
}

func (a *EditUserReqForm)ToMAP()(ret map[string]interface{}){
	ret= make(map[string]interface{},0)
	if a.Name!=nil{ ret["name"] = *a.Name}; if a.Age!=nil{ ret["age"] = *a.Age}; if a.Class!=nil{ ret["class"] = *a.Class};
	return
}

// AddUserOne add
func AddUserOne(req *AddUserReqForm)(ret *model_user.User, err error) {
	if err = req.Valid();err!=nil{
		return
	}
	var(
		data = new(model_user.User)
	)
	if err = mapstructure.Decode(req,data);err!=nil{
		return
	}
	// if needed todo add you business logic code

	if err = data.Add();err!=nil{
		return
	}

	//
	ret = data
	return
}

type UserBatchForm []*AddUserReqForm

// AddUserBatch add User
func AddUserBatch(req UserBatchForm)(ret []* model_user.User , err error) {
	var(
		datas []* model_user.User
	)
	if err = mapstructure.Decode(req,&datas);err!=nil{
		return
	}
	// if needed todo add you business logic code
	if err =model_user.AddUserBatch(datas);err!=nil{
		return
	}
	//
	ret = datas
	return
}

// EditUserOne edit
func EditUserOne(req *EditUserReqForm)(ret *model_user.User, err error) {
	if err = req.Valid();err!=nil{
		return
	}
	var(
		data =model_user.NewUser()
	)
	// if needed todo add you business logic code code
	if err = mapstructure.Decode(req, data); err != nil {
		return
	}
	if err = data.SetQueryBy(uint(req.)).Update();err!=nil{return}

	//
	ret = data
	return
}

// GetUserPage get page User data
func GetUserPage(req *model_user.QueryUserForm)(ret []*model_user.User, err error) {
	var(
		datas []*model_user.User
	)
	// if needed todo add you business logic code code

	if datas,err = model_user.GetUserList(req);err!=nil{return}

	//
	ret = datas
	return
}

// GetUserOne get User
func GetUserOne(in string)(ret *model_user.User, err error) {
	var(
		id int64
	)
	if 	id,err = strconv.ParseInt(in, 10, 64);err!=nil{return}
	var(
		d = model_user.NewUser().SetQueryBy(uint(id))
	)
	if err = d.GetByID();err!=nil{return}

	ret = d
	return
}

// DeleteUserOne delete User
func DeleteUserOne(in string)( err error) {
	var(
		id int64
	)
	if 	id,err = strconv.ParseInt(in, 10, 64);err!=nil{return}
	var(
		d = model_user.NewUser().SetQueryBy(uint(id))
	)
	// if needed todo add you business logic code
	return   d.DeleteByID()
}

// DeleteUserBatch delete User
func DeleteUserBatch(ids []string)( err error) {
	// if needed todo add you business logic code
	return   model_user.DeleteUserBatch(ids)
}
