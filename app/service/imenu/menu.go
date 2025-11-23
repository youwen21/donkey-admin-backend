package imenu

import (
	"gofly/app/model"
	"gofly/app/service/imenu/internal/dml"
	"gofly/app/service/imenu/menu_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *menu_def.MenuQueryForm) (int64, error) {
	return dml.MenuDml.Count(f)
}

func (s *srv) Query(f *menu_def.MenuQueryForm) (*menu_def.MenuQueryRes, error) {
	res, err := dml.MenuDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *menu_def.MenuQueryForm) ([]model.Menu, error) {
	res, err := dml.MenuDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.Menu, error) {
	return dml.MenuDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.Menu, error) {
	res, err := dml.MenuDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.Menu, error) {
	res, err := dml.MenuDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.Menu) (*model.Menu, error) {
	res, err := dml.MenuDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.Menu, error) {
	res, err := dml.MenuDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.Menu) error {
	err := dml.MenuDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.Menu, batchSize int) (int64, error) {
	return dml.MenuDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.Menu) (int64, error) {
	return dml.MenuDml.Update(m)
}

func (s *srv) UpdateBy(f *model.Menu, data map[string]any) (int64, error) {
	return dml.MenuDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.MenuDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.MenuDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) error {
	return dml.MenuDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.Menu, error) {
	return dml.MenuDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.Menu, error) {
	return dml.MenuDml.RawFind(sql)
}

func (s *srv) GetListBySystemId(systemId int) ([]model.Menu, error) {
	f := new(menu_def.MenuQueryForm)
	f.SystemId = systemId
	f.Page = 1
	f.PageSize = 1000
	res, err := dml.MenuDml.GetList(f)
	return res, err
}
