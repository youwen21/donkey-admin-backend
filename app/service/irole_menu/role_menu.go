package irole_menu

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/irole_menu/internal/dml"
	"donkey-admin/app/service/irole_menu/role_menu_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *role_menu_def.RoleMenuQueryForm) (int64, error) {
	return dml.RoleMenuDml.Count(f)
}

func (s *srv) Query(f *role_menu_def.RoleMenuQueryForm) (*role_menu_def.RoleMenuQueryRes, error) {
	res, err := dml.RoleMenuDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *role_menu_def.RoleMenuQueryForm) ([]model.RoleMenu, error) {
	res, err := dml.RoleMenuDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.RoleMenu, error) {
	return dml.RoleMenuDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.RoleMenu, error) {
	res, err := dml.RoleMenuDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.RoleMenu, error) {
	res, err := dml.RoleMenuDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.RoleMenu) (*model.RoleMenu, error) {
	res, err := dml.RoleMenuDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.RoleMenu, error) {
	res, err := dml.RoleMenuDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.RoleMenu) error {
	err := dml.RoleMenuDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.RoleMenu, batchSize int) (int64, error) {
	return dml.RoleMenuDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.RoleMenu) (int64, error) {
	return dml.RoleMenuDml.Update(m)
}

func (s *srv) UpdateBy(f *model.RoleMenu, data map[string]any) (int64, error) {
	return dml.RoleMenuDml.UpdateBy(f, data)
}

func (s *srv) Delete(pk int) error {
	err := dml.RoleMenuDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) error {
	return dml.RoleMenuDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.RoleMenu, error) {
	return dml.RoleMenuDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.RoleMenu, error) {
	return dml.RoleMenuDml.RawFind(sql)
}

func (s *srv) GetListBy(systemId int, roleIdList []int) ([]model.RoleMenu, error) {
	// 获取 角色授权的菜单
	rmForm := new(role_menu_def.RoleMenuQueryForm)
	rmForm.SystemId = systemId
	rmForm.RoleIdList = roleIdList
	rmForm.Page = 1
	rmForm.PageSize = 1000

	return s.GetList(rmForm)
}
