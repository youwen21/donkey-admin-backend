package irole

import (
	"gofly/app/model"
	"gofly/app/service/irole/internal/dml"
	"gofly/app/service/irole/role_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *role_def.RoleQueryForm) (int64, error) {
	return dml.RoleDml.Count(f)
}

func (s *srv) Query(f *role_def.RoleQueryForm) (*role_def.RoleQueryRes, error) {
	res, err := dml.RoleDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *role_def.RoleQueryForm) ([]model.Role, error) {
	res, err := dml.RoleDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.Role, error) {
	return dml.RoleDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.Role, error) {
	res, err := dml.RoleDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.Role, error) {
	res, err := dml.RoleDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.Role) (*model.Role, error) {
	res, err := dml.RoleDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.Role, error) {
	res, err := dml.RoleDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.Role) error {
	err := dml.RoleDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.Role, batchSize int) (int64, error) {
	return dml.RoleDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.Role) (int64, error) {
	return dml.RoleDml.Update(m)
}

func (s *srv) UpdateBy(f *model.Role, data map[string]any) (int64, error) {
	return dml.RoleDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.RoleDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.RoleDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) error {
	return dml.RoleDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.Role, error) {
	return dml.RoleDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.Role, error) {
	return dml.RoleDml.RawFind(sql)
}
