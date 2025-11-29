package user_admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iuser"
	"donkey-admin/app/service/iuser/user_def"
	"donkey-admin/lib/libutils"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *user_def.UserQueryForm) (*user_def.UserQueryResEx, error) {
	res, err := iuser.Srv.Query(f)
	// biz process
	result := new(user_def.UserQueryResEx)
	result.Total = res.Total
	result.List = make([]user_def.UserExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.User) user_def.UserExDTO {
	return user_def.UserExDTO{
		// TODO
		User: v,
	}
}

func (aSrv *adminSrv) GetList(f *user_def.UserQueryForm) ([]model.User, error) {
	res, err := iuser.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.User, error) {
	res, err := iuser.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.User, error) {
	res, err := iuser.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.User, error) {
	res, err := iuser.Srv.GetMulti(pkList)
	return res, err
}

func (aSrv *adminSrv) Add(v *model.User) (*model.User, error) {
	v.Password = libutils.EncryptWord(v.Password)

	err := iuser.Srv.Insert(v)
	return v, err
}

func (aSrv *adminSrv) Update(v *model.User) (int64, error) {
	if v.Password != "" {
		v.Password = libutils.EncryptWord(v.Password)
	}

	return iuser.Srv.Update(v)
}
