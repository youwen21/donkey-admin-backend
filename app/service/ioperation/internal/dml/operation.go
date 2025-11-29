package dml

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/ioperation/internal/dml/internal/dal"
	"donkey-admin/app/service/ioperation/operation_def"
	"donkey-admin/lib/lru"
	"fmt"
)

/*  */

const (
	operationDmlLruKey = "lru:operationDml:"
)

type operationDml struct{}

var (
	OperationDml            = &operationDml{}
	operationDmlLruCache, _ = lru.New(200)
)

func (dm *operationDml) Count(f *operation_def.OperationQueryForm) (int64, error) {
	return dal.OperationDal.Count(f)
}

func (dm *operationDml) Query(f *operation_def.OperationQueryForm) (*operation_def.OperationQueryRes, error) {
	return dal.OperationDal.Query(f)
}

func (dm *operationDml) GetList(f *operation_def.OperationQueryForm) ([]model.Operation, error) {
	return dal.OperationDal.GetList(f)
}

func (dm *operationDml) GetLisByPkList(pkList []int) ([]model.Operation, error) {
	return dal.OperationDal.GetLisByPkList(pkList)
}

func (dm *operationDml) GetAll() ([]model.Operation, error) {
	return dal.OperationDal.GetAll()
}

func (dm *operationDml) Get(pk int) (*model.Operation, error) {
	return dal.OperationDal.Get(pk)
}

func (dm *operationDml) GetBy(f *model.Operation) (*model.Operation, error) {
	return dal.OperationDal.GetBy(f)
}

func (dm *operationDml) GetMulti(pkList []int) (map[int]model.Operation, error) {
	return dal.OperationDal.GetMulti(pkList)
}

func (dm *operationDml) Insert(m *model.Operation) error {
	return dal.OperationDal.Insert(m)
}

func (dm *operationDml) BatchInsert(bm []*model.Operation, batchSize int) (int64, error) {
	return dal.OperationDal.BatchInsert(bm, batchSize)
}

func (dm *operationDml) Update(m *model.Operation) (int64, error) {
	return dal.OperationDal.Update(m)
}

func (dm *operationDml) UpdateBy(f *model.Operation, data map[string]any) (int64, error) {
	return dal.OperationDal.UpdateBy(f, data)
}

func (dm *operationDml) SetInfo(data map[string]any) (int64, error) {
	return dal.OperationDal.SetInfo(data)
}

func (dm *operationDml) Delete(pk int) error {
	return dal.OperationDal.Delete(pk)
}

func (dm *operationDml) Exec(sql string, values ...interface{}) error {
	return dal.OperationDal.Exec(sql, values...)
}

func (dm *operationDml) RawGet(sql string) (*model.Operation, error) {
	return dal.OperationDal.RawGet(sql)
}

func (dm *operationDml) RawFind(sql string) ([]model.Operation, error) {
	return dal.OperationDal.RawFind(sql)
}

/* ----  lru  ---- */

func (dm *operationDml) LruGetKey(key interface{}) string {
	return fmt.Sprintf("%s%v", operationDmlLruKey, key)
}

func (dm *operationDml) LruRemove(key string) bool {
	return operationDmlLruCache.Remove(key)
}

// 非单台机器提供服务的情况下， 完善此处
func (dm *operationDml) LruPublishRemove(key string) error {
	dm.LruRemove(key)
	// TODO
	return nil
}

/* ----  lru  ---- */
