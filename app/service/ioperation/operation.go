package ioperation

import (
	"gofly/app/model"
	"gofly/app/service/ioperation/internal/dml"
	"gofly/app/service/ioperation/operation_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *operation_def.OperationQueryForm) (int64, error) {
	return dml.OperationDml.Count(f)
}

func (s *srv) Query(f *operation_def.OperationQueryForm) (*operation_def.OperationQueryRes, error) {
	res, err := dml.OperationDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *operation_def.OperationQueryForm) ([]model.Operation, error) {
	res, err := dml.OperationDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.Operation, error) {
	return dml.OperationDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.Operation, error) {
	res, err := dml.OperationDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.Operation, error) {
	res, err := dml.OperationDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.Operation) (*model.Operation, error) {
	res, err := dml.OperationDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.Operation, error) {
	res, err := dml.OperationDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.Operation) error {
	err := dml.OperationDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.Operation, batchSize int) (int64, error) {
	return dml.OperationDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.Operation) (int64, error) {
	return dml.OperationDml.Update(m)
}

func (s *srv) UpdateBy(f *model.Operation, data map[string]any) (int64, error) {
	return dml.OperationDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.OperationDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.OperationDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) error {
	return dml.OperationDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.Operation, error) {
	return dml.OperationDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.Operation, error) {
	return dml.OperationDml.RawFind(sql)
}
