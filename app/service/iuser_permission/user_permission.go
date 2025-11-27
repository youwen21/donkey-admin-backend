package iuser_permission

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/iuser_permission/internal/dml"
	"gofly/app/service/iuser_permission/user_permission_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *user_permission_def.UserPermissionQueryForm) (int64, error) {
	return dml.UserPermissionDml.Count(f)
}

func (s *srv) Query(f *user_permission_def.UserPermissionQueryForm) (*user_permission_def.UserPermissionQueryRes, error) {
	res, err := dml.UserPermissionDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *user_permission_def.UserPermissionQueryForm) ([]model.UserPermission, error) {
	res, err := dml.UserPermissionDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.UserPermission, error) {
	return dml.UserPermissionDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.UserPermission, error) {
	res, err := dml.UserPermissionDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.UserPermission, error) {
	res, err := dml.UserPermissionDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.UserPermission) (*model.UserPermission, error) {
	res, err := dml.UserPermissionDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.UserPermission, error) {
	res, err := dml.UserPermissionDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.UserPermission) error {
	err := dml.UserPermissionDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.UserPermission, batchSize int) (int64, error) {
	return dml.UserPermissionDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.UserPermission) (int64, error) {
	return dml.UserPermissionDml.Update(m)
}

func (s *srv) UpdateBy(f *model.UserPermission, data map[string]any) (int64, error) {
	return dml.UserPermissionDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.UserPermissionDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.UserPermissionDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) (int64, error) {
	return dml.UserPermissionDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.UserPermission, error) {
	return dml.UserPermissionDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.UserPermission, error) {
	return dml.UserPermissionDml.RawFind(sql)
}

func (s *srv) ClearUserPermission(userId int, systemId int) (int64, error) {
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `user_id` = ? AND `system_id` = ? ", (&model.UserPermission{}).TableName())
	return s.Exec(sql, userId, systemId)
}
