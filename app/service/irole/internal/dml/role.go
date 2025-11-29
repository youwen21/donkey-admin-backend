package dml

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/irole/internal/dml/internal/dal"
	"donkey-admin/app/service/irole/role_def"
	"donkey-admin/lib/lru"
	"fmt"
)

/*  */

const (
	roleDmlLruKey = "lru:roleDml:"
)

type roleDml struct{}

var (
	RoleDml            = &roleDml{}
	roleDmlLruCache, _ = lru.New(200)
)

func (dm *roleDml) Count(f *role_def.RoleQueryForm) (int64, error) {
	return dal.RoleDal.Count(f)
}

func (dm *roleDml) Query(f *role_def.RoleQueryForm) (*role_def.RoleQueryRes, error) {
	return dal.RoleDal.Query(f)
}

func (dm *roleDml) GetList(f *role_def.RoleQueryForm) ([]model.Role, error) {
	return dal.RoleDal.GetList(f)
}

func (dm *roleDml) GetLisByPkList(pkList []int) ([]model.Role, error) {
	return dal.RoleDal.GetLisByPkList(pkList)
}

func (dm *roleDml) GetAll() ([]model.Role, error) {
	return dal.RoleDal.GetAll()
}

func (dm *roleDml) Get(pk int) (*model.Role, error) {
	return dal.RoleDal.Get(pk)
}

func (dm *roleDml) GetBy(f *model.Role) (*model.Role, error) {
	return dal.RoleDal.GetBy(f)
}

func (dm *roleDml) GetMulti(pkList []int) (map[int]model.Role, error) {
	return dal.RoleDal.GetMulti(pkList)
}

func (dm *roleDml) Insert(m *model.Role) error {
	return dal.RoleDal.Insert(m)
}

func (dm *roleDml) BatchInsert(bm []*model.Role, batchSize int) (int64, error) {
	return dal.RoleDal.BatchInsert(bm, batchSize)
}

func (dm *roleDml) Update(m *model.Role) (int64, error) {
	return dal.RoleDal.Update(m)
}

func (dm *roleDml) UpdateBy(f *model.Role, data map[string]any) (int64, error) {
	return dal.RoleDal.UpdateBy(f, data)
}

func (dm *roleDml) SetInfo(data map[string]any) (int64, error) {
	return dal.RoleDal.SetInfo(data)
}

func (dm *roleDml) Delete(pk int) error {
	return dal.RoleDal.Delete(pk)
}

func (dm *roleDml) Exec(sql string, values ...interface{}) error {
	return dal.RoleDal.Exec(sql, values...)
}

func (dm *roleDml) RawGet(sql string) (*model.Role, error) {
	return dal.RoleDal.RawGet(sql)
}

func (dm *roleDml) RawFind(sql string) ([]model.Role, error) {
	return dal.RoleDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *roleDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", roleDmlLruKey, key)
}

func (dm *roleDml) LruRemove(key string) bool {
	return roleDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *roleDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
