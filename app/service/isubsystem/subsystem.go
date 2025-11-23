package isubsystem

import (
	"gofly/app/model"
	"gofly/app/service/isubsystem/internal/dml"
	"gofly/app/service/isubsystem/subsystem_def"
)

/*  */

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Count(f *subsystem_def.SubsystemQueryForm) (int64, error) {
	return dml.SubsystemDml.Count(f)
}

func (s *srv) Query(f *subsystem_def.SubsystemQueryForm) (*subsystem_def.SubsystemQueryRes, error) {
	res, err := dml.SubsystemDml.Query(f)
	// biz process
	return res, err
}

func (s *srv) GetList(f *subsystem_def.SubsystemQueryForm) ([]model.Subsystem, error) {
	res, err := dml.SubsystemDml.GetList(f)
	return res, err
}

func (s *srv) GetLisByPkList(pkList []int) ([]model.Subsystem, error) {
	return dml.SubsystemDml.GetLisByPkList(pkList)
}

func (s *srv) GetAll() ([]model.Subsystem, error) {
	res, err := dml.SubsystemDml.GetAll()
	return res, err
}

func (s *srv) Get(pk int) (*model.Subsystem, error) {
	res, err := dml.SubsystemDml.Get(pk)
	return res, err
}

func (s *srv) GetBy(f *model.Subsystem) (*model.Subsystem, error) {
	res, err := dml.SubsystemDml.GetBy(f)
	return res, err
}

func (s *srv) GetMulti(pkList []int) (map[int]model.Subsystem, error) {
	res, err := dml.SubsystemDml.GetMulti(pkList)
	return res, err
}

func (s *srv) Insert(m *model.Subsystem) error {
	err := dml.SubsystemDml.Insert(m)
	return err
}

func (s *srv) BatchInsert(bm []*model.Subsystem, batchSize int) (int64, error) {
	return dml.SubsystemDml.BatchInsert(bm, batchSize)
}

func (s *srv) Update(m *model.Subsystem) (int64, error) {
	return dml.SubsystemDml.Update(m)
}

func (s *srv) UpdateBy(f *model.Subsystem, data map[string]any) (int64, error) {
	return dml.SubsystemDml.UpdateBy(f, data)
}

func (s *srv) SetInfo(data map[string]any) (int64, error) {
	return dml.SubsystemDml.SetInfo(data)
}

func (s *srv) Delete(pk int) error {
	err := dml.SubsystemDml.Delete(pk)
	return err
}

func (s *srv) Exec(sql string, values ...interface{}) error {
	return dml.SubsystemDml.Exec(sql, values...)
}

func (s *srv) RawGet(sql string) (*model.Subsystem, error) {
	return dml.SubsystemDml.RawGet(sql)
}

func (s *srv) RawFind(sql string) ([]model.Subsystem, error) {
	return dml.SubsystemDml.RawFind(sql)
}
