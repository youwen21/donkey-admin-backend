package iuser_role

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iuser_role/internal/dml"
	"donkey-admin/app/service/iuser_role/user_role_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *user_role_def.UserRoleQueryForm) (int64, error) {
	return dml.UserRoleDml.Count(f)
}

func (s *srv) Query(f *user_role_def.UserRoleQueryForm) (*user_role_def.UserRoleQueryRes, error) {
	res, err := dml.UserRoleDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *user_role_def.UserRoleQueryForm) ([]model.UserRole, error) {
	res, err := dml.UserRoleDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.UserRole, error) {
	return dml.UserRoleDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.UserRole, error) {
	res, err := dml.UserRoleDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.UserRole, error) {
	res, err := dml.UserRoleDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.UserRole) (*model.UserRole, error) {
	res, err := dml.UserRoleDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.UserRole, error) {
	res, err := dml.UserRoleDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.UserRole) error {
	err := dml.UserRoleDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.UserRole, batchSize int) (int64, error) {
	return dml.UserRoleDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.UserRole) (int64, error) {
	return dml.UserRoleDml.Update(m)
}

func (s *srv) UpdateBy(f *model.UserRole, data map[string]any) (int64, error) {
	return dml.UserRoleDml.UpdateBy(f, data)
}

func (s *srv) Delete(pk int) error {
	err := dml.UserRoleDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) error {
	return dml.UserRoleDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.UserRole, error) {
	return dml.UserRoleDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.UserRole, error) {
	return dml.UserRoleDml.RawFind(sql)
}

func (s *srv) GetListByUserId(userId int) ([]model.UserRole, error) {
	form := new(user_role_def.UserRoleQueryForm)
	form.UserId = userId
	form.Page = 1
	form.PageSize = 1000
	return s.GetList(form)
}
