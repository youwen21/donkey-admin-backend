package dml

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/iuser_permission/internal/dml/internal/dal"
	"gofly/app/service/iuser_permission/user_permission_def"
	"gofly/lib/lru"
)

/*  */

const (
	userPermissionDmlLruKey = "lru:userPermissionDml:"
)

type userPermissionDml struct{}

var (
	UserPermissionDml            = &userPermissionDml{}
	userPermissionDmlLruCache, _ = lru.New(200)
)

func (dm *userPermissionDml) Count(f *user_permission_def.UserPermissionQueryForm) (int64, error) {
	return dal.UserPermissionDal.Count(f)
}

func (dm *userPermissionDml) Query(f *user_permission_def.UserPermissionQueryForm) (*user_permission_def.UserPermissionQueryRes, error) {
	return dal.UserPermissionDal.Query(f)
}

func (dm *userPermissionDml) GetList(f *user_permission_def.UserPermissionQueryForm) ([]model.UserPermission, error) {
	return dal.UserPermissionDal.GetList(f)
}

func (dm *userPermissionDml) GetLisByPkList(pkList []int) ([]model.UserPermission, error) {
	return dal.UserPermissionDal.GetLisByPkList(pkList)
}

func (dm *userPermissionDml) GetAll() ([]model.UserPermission, error) {
	return dal.UserPermissionDal.GetAll()
}

func (dm *userPermissionDml) Get(pk int) (*model.UserPermission, error) {
	return dal.UserPermissionDal.Get(pk)
}

func (dm *userPermissionDml) GetBy(f *model.UserPermission) (*model.UserPermission, error) {
	return dal.UserPermissionDal.GetBy(f)
}

func (dm *userPermissionDml) GetMulti(pkList []int) (map[int]model.UserPermission, error) {
	return dal.UserPermissionDal.GetMulti(pkList)
}

func (dm *userPermissionDml) Insert(m *model.UserPermission) error {
	return dal.UserPermissionDal.Insert(m)
}

func (dm *userPermissionDml) BatchInsert(bm []*model.UserPermission, batchSize int) (int64, error) {
	return dal.UserPermissionDal.BatchInsert(bm, batchSize)
}

func (dm *userPermissionDml) Update(m *model.UserPermission) (int64, error) {
	return dal.UserPermissionDal.Update(m)
}

func (dm *userPermissionDml) UpdateBy(f *model.UserPermission, data map[string]any) (int64, error) {
	return dal.UserPermissionDal.UpdateBy(f, data)
}

func (dm *userPermissionDml) SetInfo(data map[string]any) (int64, error) {
	return dal.UserPermissionDal.SetInfo(data)
}

func (dm *userPermissionDml) Delete(pk int) error {
	return dal.UserPermissionDal.Delete(pk)
}

func (dm *userPermissionDml) Exec(sql string, values ...interface{}) (int64, error) {
	return dal.UserPermissionDal.Exec(sql, values...)
}

func (dm *userPermissionDml) RawGet(sql string) (*model.UserPermission, error) {
	return dal.UserPermissionDal.RawGet(sql)
}

func (dm *userPermissionDml) RawFind(sql string) ([]model.UserPermission, error) {
	return dal.UserPermissionDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *userPermissionDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", userPermissionDmlLruKey, key)
}

func (dm *userPermissionDml) LruRemove(key string) bool {
	return userPermissionDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *userPermissionDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
