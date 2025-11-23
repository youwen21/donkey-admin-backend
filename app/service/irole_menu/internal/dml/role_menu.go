package dml

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/irole_menu/internal/dml/internal/dal"
	"gofly/app/service/irole_menu/role_menu_def"
	"gofly/lib/lru"
)

/*  */

const (
	roleMenuDmlLruKey = "lru:roleMenuDml:"
)

type roleMenuDml struct{}

var (
	RoleMenuDml            = &roleMenuDml{}
	roleMenuDmlLruCache, _ = lru.New(200)
)

func (dm *roleMenuDml) Count(f *role_menu_def.RoleMenuQueryForm) (int64, error) {
	return dal.RoleMenuDal.Count(f)
}

func (dm *roleMenuDml) Query(f *role_menu_def.RoleMenuQueryForm) (*role_menu_def.RoleMenuQueryRes, error) {
	return dal.RoleMenuDal.Query(f)
}

func (dm *roleMenuDml) GetList(f *role_menu_def.RoleMenuQueryForm) ([]model.RoleMenu, error) {
	return dal.RoleMenuDal.GetList(f)
}

func (dm *roleMenuDml) GetLisByPkList(pkList []int) ([]model.RoleMenu, error) {
	return dal.RoleMenuDal.GetLisByPkList(pkList)
}

func (dm *roleMenuDml) GetAll() ([]model.RoleMenu, error) {
	return dal.RoleMenuDal.GetAll()
}

func (dm *roleMenuDml) Get(pk int) (*model.RoleMenu, error) {
	return dal.RoleMenuDal.Get(pk)
}

func (dm *roleMenuDml) GetBy(f *model.RoleMenu) (*model.RoleMenu, error) {
	return dal.RoleMenuDal.GetBy(f)
}

func (dm *roleMenuDml) GetMulti(pkList []int) (map[int]model.RoleMenu, error) {
	return dal.RoleMenuDal.GetMulti(pkList)
}

func (dm *roleMenuDml) Insert(m *model.RoleMenu) error {
	return dal.RoleMenuDal.Insert(m)
}

func (dm *roleMenuDml) BatchInsert(bm []*model.RoleMenu, batchSize int) (int64, error) {
	return dal.RoleMenuDal.BatchInsert(bm, batchSize)
}

func (dm *roleMenuDml) Update(m *model.RoleMenu) (int64, error) {
	return dal.RoleMenuDal.Update(m)
}

func (dm *roleMenuDml) UpdateBy(f *model.RoleMenu, data map[string]any) (int64, error) {
	return dal.RoleMenuDal.UpdateBy(f, data)
}

func (dm *roleMenuDml) Delete(pk int) error {
	return dal.RoleMenuDal.Delete(pk)
}

func (dm *roleMenuDml) Exec(sql string, values ...interface{}) error {
	return dal.RoleMenuDal.Exec(sql, values...)
}

func (dm *roleMenuDml) RawGet(sql string) (*model.RoleMenu, error) {
	return dal.RoleMenuDal.RawGet(sql)
}

func (dm *roleMenuDml) RawFind(sql string) ([]model.RoleMenu, error) {
	return dal.RoleMenuDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *roleMenuDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", roleMenuDmlLruKey, key)
}

func (dm *roleMenuDml) LruRemove(key string) bool {
	return roleMenuDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *roleMenuDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
