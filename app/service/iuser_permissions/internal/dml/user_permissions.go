package dml

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/iuser_permissions/internal/dml/internal/dal"
	"gofly/app/service/iuser_permissions/user_permissions_def"
	"gofly/lib/lru"
)

/*  */

const (
	userPermissionsDmlLruKey = "lru:userPermissionsDml:"
)

type userPermissionsDml struct{}

var (
	UserPermissionsDml            = &userPermissionsDml{}
	userPermissionsDmlLruCache, _ = lru.New(200)
)

func (dm *userPermissionsDml) Count(f *user_permissions_def.UserPermissionsQueryForm) (int64, error) {
	return dal.UserPermissionsDal.Count(f)
}

func (dm *userPermissionsDml) Query(f *user_permissions_def.UserPermissionsQueryForm) (*user_permissions_def.UserPermissionsQueryRes, error) {
	return dal.UserPermissionsDal.Query(f)
}

func (dm *userPermissionsDml) GetList(f *user_permissions_def.UserPermissionsQueryForm) ([]model.UserPermissions, error) {
	return dal.UserPermissionsDal.GetList(f)
}

func (dm *userPermissionsDml) GetLisByPkList(pkList []int) ([]model.UserPermissions, error) {
	return dal.UserPermissionsDal.GetLisByPkList(pkList)
}

func (dm *userPermissionsDml) GetAll() ([]model.UserPermissions, error) {
	return dal.UserPermissionsDal.GetAll()
}

func (dm *userPermissionsDml) Get(pk int) (*model.UserPermissions, error) {
	return dal.UserPermissionsDal.Get(pk)
}

func (dm *userPermissionsDml) GetBy(f *model.UserPermissions) (*model.UserPermissions, error) {
	return dal.UserPermissionsDal.GetBy(f)
}

func (dm *userPermissionsDml) GetMulti(pkList []int) (map[int]model.UserPermissions, error) {
	return dal.UserPermissionsDal.GetMulti(pkList)
}

func (dm *userPermissionsDml) Insert(m *model.UserPermissions) error {
	return dal.UserPermissionsDal.Insert(m)
}

func (dm *userPermissionsDml) BatchInsert(bm []*model.UserPermissions, batchSize int) (int64, error) {
	return dal.UserPermissionsDal.BatchInsert(bm, batchSize)
}

func (dm *userPermissionsDml) Update(m *model.UserPermissions) (int64, error) {
	return dal.UserPermissionsDal.Update(m)
}

func (dm *userPermissionsDml) UpdateBy(f *model.UserPermissions, data map[string]any) (int64, error) {
	return dal.UserPermissionsDal.UpdateBy(f, data)
}

func (dm *userPermissionsDml) SetInfo(data map[string]any) (int64, error) {
	return dal.UserPermissionsDal.SetInfo(data)
}

func (dm *userPermissionsDml) Delete(pk int) error {
	return dal.UserPermissionsDal.Delete(pk)
}

func (dm *userPermissionsDml) Exec(sql string, values ...interface{}) error {
	return dal.UserPermissionsDal.Exec(sql, values...)
}

func (dm *userPermissionsDml) RawGet(sql string) (*model.UserPermissions, error) {
	return dal.UserPermissionsDal.RawGet(sql)
}

func (dm *userPermissionsDml) RawFind(sql string) ([]model.UserPermissions, error) {
	return dal.UserPermissionsDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *userPermissionsDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", userPermissionsDmlLruKey, key)
}

func (dm *userPermissionsDml) LruRemove(key string) bool {
	return userPermissionsDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *userPermissionsDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
