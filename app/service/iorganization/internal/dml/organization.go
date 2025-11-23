package dml

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/iorganization/internal/dml/internal/dal"
	"gofly/app/service/iorganization/organization_def"
	"gofly/lib/lru"
)

/*  */

const (
	organizationDmlLruKey = "lru:organizationDml:"
)

type organizationDml struct{}

var (
	OrganizationDml            = &organizationDml{}
	organizationDmlLruCache, _ = lru.New(200)
)

func (dm *organizationDml) Count(f *organization_def.OrganizationQueryForm) (int64, error) {
	return dal.OrganizationDal.Count(f)
}

func (dm *organizationDml) Query(f *organization_def.OrganizationQueryForm) (*organization_def.OrganizationQueryRes, error) {
	return dal.OrganizationDal.Query(f)
}

func (dm *organizationDml) GetList(f *organization_def.OrganizationQueryForm) ([]model.Organization, error) {
	return dal.OrganizationDal.GetList(f)
}

func (dm *organizationDml) GetLisByPkList(pkList []int) ([]model.Organization, error) {
	return dal.OrganizationDal.GetLisByPkList(pkList)
}

func (dm *organizationDml) GetAll() ([]model.Organization, error) {
	return dal.OrganizationDal.GetAll()
}

func (dm *organizationDml) Get(pk int) (*model.Organization, error) {
	return dal.OrganizationDal.Get(pk)
}

func (dm *organizationDml) GetBy(f *model.Organization) (*model.Organization, error) {
	return dal.OrganizationDal.GetBy(f)
}

func (dm *organizationDml) GetMulti(pkList []int) (map[int]model.Organization, error) {
	return dal.OrganizationDal.GetMulti(pkList)
}

func (dm *organizationDml) Insert(m *model.Organization) error {
	return dal.OrganizationDal.Insert(m)
}

func (dm *organizationDml) BatchInsert(bm []*model.Organization, batchSize int) (int64, error) {
	return dal.OrganizationDal.BatchInsert(bm, batchSize)
}

func (dm *organizationDml) Update(m *model.Organization) (int64, error) {
	return dal.OrganizationDal.Update(m)
}

func (dm *organizationDml) UpdateBy(f *model.Organization, data map[string]any) (int64, error) {
	return dal.OrganizationDal.UpdateBy(f, data)
}

func (dm *organizationDml) SetInfo(data map[string]any) (int64, error) {
	return dal.OrganizationDal.SetInfo(data)
}

func (dm *organizationDml) Delete(pk int) error {
	return dal.OrganizationDal.Delete(pk)
}

func (dm *organizationDml) Exec(sql string, values ...interface{}) error {
	return dal.OrganizationDal.Exec(sql, values...)
}

func (dm *organizationDml) RawGet(sql string) (*model.Organization, error) {
	return dal.OrganizationDal.RawGet(sql)
}

func (dm *organizationDml) RawFind(sql string) ([]model.Organization, error) {
	return dal.OrganizationDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *organizationDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", organizationDmlLruKey, key)
}

func (dm *organizationDml) LruRemove(key string) bool {
	return organizationDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *organizationDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
