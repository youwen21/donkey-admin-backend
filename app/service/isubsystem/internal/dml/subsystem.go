package dml

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/isubsystem/internal/dml/internal/dal"
	"donkey-admin/app/service/isubsystem/subsystem_def"
	"donkey-admin/lib/lru"
	"fmt"
)

/*  */

const (
	subsystemDmlLruKey = "lru:subsystemDml:"
)

type subsystemDml struct{}

var (
	SubsystemDml            = &subsystemDml{}
	subsystemDmlLruCache, _ = lru.New(200)
)

func (dm *subsystemDml) Count(f *subsystem_def.SubsystemQueryForm) (int64, error) {
	return dal.SubsystemDal.Count(f)
}

func (dm *subsystemDml) Query(f *subsystem_def.SubsystemQueryForm) (*subsystem_def.SubsystemQueryRes, error) {
	return dal.SubsystemDal.Query(f)
}

func (dm *subsystemDml) GetList(f *subsystem_def.SubsystemQueryForm) ([]model.Subsystem, error) {
	return dal.SubsystemDal.GetList(f)
}

func (dm *subsystemDml) GetLisByPkList(pkList []int) ([]model.Subsystem, error) {
	return dal.SubsystemDal.GetLisByPkList(pkList)
}

func (dm *subsystemDml) GetAll() ([]model.Subsystem, error) {
	return dal.SubsystemDal.GetAll()
}

func (dm *subsystemDml) Get(pk int) (*model.Subsystem, error) {
	return dal.SubsystemDal.Get(pk)
}

func (dm *subsystemDml) GetBy(f *model.Subsystem) (*model.Subsystem, error) {
	return dal.SubsystemDal.GetBy(f)
}

func (dm *subsystemDml) GetMulti(pkList []int) (map[int]model.Subsystem, error) {
	return dal.SubsystemDal.GetMulti(pkList)
}

func (dm *subsystemDml) Insert(m *model.Subsystem) error {
	return dal.SubsystemDal.Insert(m)
}

func (dm *subsystemDml) BatchInsert(bm []*model.Subsystem, batchSize int) (int64, error) {
	return dal.SubsystemDal.BatchInsert(bm, batchSize)
}

func (dm *subsystemDml) Update(m *model.Subsystem) (int64, error) {
	return dal.SubsystemDal.Update(m)
}

func (dm *subsystemDml) UpdateBy(f *model.Subsystem, data map[string]any) (int64, error) {
	return dal.SubsystemDal.UpdateBy(f, data)
}

func (dm *subsystemDml) SetInfo(data map[string]any) (int64, error) {
	return dal.SubsystemDal.SetInfo(data)
}

func (dm *subsystemDml) Delete(pk int) error {
	return dal.SubsystemDal.Delete(pk)
}

func (dm *subsystemDml) Exec(sql string, values ...interface{}) error {
	return dal.SubsystemDal.Exec(sql, values...)
}

func (dm *subsystemDml) RawGet(sql string) (*model.Subsystem, error) {
	return dal.SubsystemDal.RawGet(sql)
}

func (dm *subsystemDml) RawFind(sql string) ([]model.Subsystem, error) {
	return dal.SubsystemDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *subsystemDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", subsystemDmlLruKey, key)
}

func (dm *subsystemDml) LruRemove(key string) bool {
	return subsystemDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *subsystemDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
