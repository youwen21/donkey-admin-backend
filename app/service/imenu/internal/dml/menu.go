package dml

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/imenu/internal/dml/internal/dal"
	"gofly/app/service/imenu/menu_def"
	"gofly/lib/lru"
)

/*  */

const (
	menuDmlLruKey = "lru:menuDml:"
)

type menuDml struct{}

var (
	MenuDml            = &menuDml{}
	menuDmlLruCache, _ = lru.New(200)
)

func (dm *menuDml) Count(f *menu_def.MenuQueryForm) (int64, error) {
	return dal.MenuDal.Count(f)
}

func (dm *menuDml) Query(f *menu_def.MenuQueryForm) (*menu_def.MenuQueryRes, error) {
	return dal.MenuDal.Query(f)
}

func (dm *menuDml) GetList(f *menu_def.MenuQueryForm) ([]model.Menu, error) {
	return dal.MenuDal.GetList(f)
}

func (dm *menuDml) GetLisByPkList(pkList []int) ([]model.Menu, error) {
	return dal.MenuDal.GetLisByPkList(pkList)
}

func (dm *menuDml) GetAll() ([]model.Menu, error) {
	return dal.MenuDal.GetAll()
}

func (dm *menuDml) Get(pk int) (*model.Menu, error) {
	return dal.MenuDal.Get(pk)
}

func (dm *menuDml) GetBy(f *model.Menu) (*model.Menu, error) {
	return dal.MenuDal.GetBy(f)
}

func (dm *menuDml) GetMulti(pkList []int) (map[int]model.Menu, error) {
	return dal.MenuDal.GetMulti(pkList)
}

func (dm *menuDml) Insert(m *model.Menu) error {
	return dal.MenuDal.Insert(m)
}

func (dm *menuDml) BatchInsert(bm []*model.Menu, batchSize int) (int64, error) {
	return dal.MenuDal.BatchInsert(bm, batchSize)
}

func (dm *menuDml) Update(m *model.Menu) (int64, error) {
	return dal.MenuDal.Update(m)
}

func (dm *menuDml) UpdateBy(f *model.Menu, data map[string]any) (int64, error) {
	return dal.MenuDal.UpdateBy(f, data)
}

func (dm *menuDml) SetInfo(data map[string]any) (int64, error) {
	return dal.MenuDal.SetInfo(data)
}

func (dm *menuDml) Delete(pk int) error {
	return dal.MenuDal.Delete(pk)
}

func (dm *menuDml) Exec(sql string, values ...interface{}) error {
	return dal.MenuDal.Exec(sql, values...)
}

func (dm *menuDml) RawGet(sql string) (*model.Menu, error) {
	return dal.MenuDal.RawGet(sql)
}

func (dm *menuDml) RawFind(sql string) ([]model.Menu, error) {
	return dal.MenuDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *menuDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", menuDmlLruKey, key)
}

func (dm *menuDml) LruRemove(key string) bool {
	return menuDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *menuDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
