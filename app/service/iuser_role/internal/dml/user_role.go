package dml

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iuser_role/internal/dml/internal/dal"
	"donkey-admin/app/service/iuser_role/user_role_def"
	"donkey-admin/lib/lru"
	"fmt"
)

/*  */

const (
	userRoleDmlLruKey = "lru:userRoleDml:"
)

type userRoleDml struct{}

var (
	UserRoleDml            = &userRoleDml{}
	userRoleDmlLruCache, _ = lru.New(200)
)

func (dm *userRoleDml) Count(f *user_role_def.UserRoleQueryForm) (int64, error) {
	return dal.UserRoleDal.Count(f)
}

func (dm *userRoleDml) Query(f *user_role_def.UserRoleQueryForm) (*user_role_def.UserRoleQueryRes, error) {
	return dal.UserRoleDal.Query(f)
}

func (dm *userRoleDml) GetList(f *user_role_def.UserRoleQueryForm) ([]model.UserRole, error) {
	return dal.UserRoleDal.GetList(f)
}

func (dm *userRoleDml) GetLisByPkList(pkList []int) ([]model.UserRole, error) {
	return dal.UserRoleDal.GetLisByPkList(pkList)
}

func (dm *userRoleDml) GetAll() ([]model.UserRole, error) {
	return dal.UserRoleDal.GetAll()
}

func (dm *userRoleDml) Get(pk int) (*model.UserRole, error) {
	return dal.UserRoleDal.Get(pk)
}

func (dm *userRoleDml) GetBy(f *model.UserRole) (*model.UserRole, error) {
	return dal.UserRoleDal.GetBy(f)
}

func (dm *userRoleDml) GetMulti(pkList []int) (map[int]model.UserRole, error) {
	return dal.UserRoleDal.GetMulti(pkList)
}

func (dm *userRoleDml) Insert(m *model.UserRole) error {
	return dal.UserRoleDal.Insert(m)
}

func (dm *userRoleDml) BatchInsert(bm []*model.UserRole, batchSize int) (int64, error) {
	return dal.UserRoleDal.BatchInsert(bm, batchSize)
}

func (dm *userRoleDml) Update(m *model.UserRole) (int64, error) {
	return dal.UserRoleDal.Update(m)
}

func (dm *userRoleDml) UpdateBy(f *model.UserRole, data map[string]any) (int64, error) {
	return dal.UserRoleDal.UpdateBy(f, data)
}

func (dm *userRoleDml) Delete(pk int) error {
	return dal.UserRoleDal.Delete(pk)
}

func (dm *userRoleDml) Exec(sql string, values ...interface{}) error {
	return dal.UserRoleDal.Exec(sql, values...)
}

func (dm *userRoleDml) RawGet(sql string) (*model.UserRole, error) {
	return dal.UserRoleDal.RawGet(sql)
}

func (dm *userRoleDml) RawFind(sql string) ([]model.UserRole, error) {
	return dal.UserRoleDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *userRoleDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", userRoleDmlLruKey, key)
}

func (dm *userRoleDml) LruRemove(key string) bool {
	return userRoleDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *userRoleDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
