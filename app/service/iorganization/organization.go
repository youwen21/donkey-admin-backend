package iorganization

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iorganization/internal/dml"
	"donkey-admin/app/service/iorganization/organization_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *organization_def.OrganizationQueryForm) (int64, error) {
	return dml.OrganizationDml.Count(f)
}

func (s *srv) Query(f *organization_def.OrganizationQueryForm) (*organization_def.OrganizationQueryRes, error) {
	res, err := dml.OrganizationDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *organization_def.OrganizationQueryForm) ([]model.Organization, error) {
	res, err := dml.OrganizationDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.Organization, error) {
	return dml.OrganizationDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.Organization, error) {
	res, err := dml.OrganizationDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.Organization, error) {
	res, err := dml.OrganizationDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.Organization) (*model.Organization, error) {
	res, err := dml.OrganizationDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.Organization, error) {
	res, err := dml.OrganizationDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.Organization) error {
	err := dml.OrganizationDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.Organization, batchSize int) (int64, error) {
	return dml.OrganizationDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.Organization) (int64, error) {
	return dml.OrganizationDml.Update(m)
}

func (s *srv) UpdateBy(f *model.Organization, data map[string]any) (int64, error) {
	return dml.OrganizationDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.OrganizationDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.OrganizationDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) error {
	return dml.OrganizationDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.Organization, error) {
	return dml.OrganizationDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.Organization, error) {
	return dml.OrganizationDml.RawFind(sql)
}
