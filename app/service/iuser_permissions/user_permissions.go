package iuser_permissions

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/iuser_permissions/internal/dml"
	"gofly/app/service/iuser_permissions/user_permissions_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *user_permissions_def.UserPermissionsQueryForm) (int64, error) {
	return dml.UserPermissionsDml.Count(f)
}

func (s *srv) Query(f *user_permissions_def.UserPermissionsQueryForm) (*user_permissions_def.UserPermissionsQueryRes, error) {
	res, err := dml.UserPermissionsDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *user_permissions_def.UserPermissionsQueryForm) ([]model.UserPermissions, error) {
	res, err := dml.UserPermissionsDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.UserPermissions, error) {
	return dml.UserPermissionsDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.UserPermissions, error) {
	res, err := dml.UserPermissionsDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.UserPermissions, error) {
	res, err := dml.UserPermissionsDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.UserPermissions) (*model.UserPermissions, error) {
	res, err := dml.UserPermissionsDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.UserPermissions, error) {
	res, err := dml.UserPermissionsDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.UserPermissions) error {
	err := dml.UserPermissionsDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.UserPermissions, batchSize int) (int64, error) {
	return dml.UserPermissionsDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.UserPermissions) (int64, error) {
	return dml.UserPermissionsDml.Update(m)
}

func (s *srv) UpdateBy(f *model.UserPermissions, data map[string]any) (int64, error) {
	return dml.UserPermissionsDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.UserPermissionsDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.UserPermissionsDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) (int64, error) {
	return dml.UserPermissionsDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.UserPermissions, error) {
	return dml.UserPermissionsDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.UserPermissions, error) {
	return dml.UserPermissionsDml.RawFind(sql)
}

func (s *srv) ClearUserPermissions(userId int, systemId int) (int64, error) {
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `user_id` = ? AND `system_id` = ? ", (&model.UserPermissions{}).TableName())
	return s.Exec(sql, userId, systemId)
}
