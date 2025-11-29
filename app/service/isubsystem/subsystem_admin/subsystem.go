package subsystem_admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/isubsystem"
	"donkey-admin/app/service/isubsystem/subsystem_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *subsystem_def.SubsystemQueryForm) (*subsystem_def.SubsystemQueryResEx, error) {
	res, err := isubsystem.Srv.Query(f)
	// biz process
	result := new(subsystem_def.SubsystemQueryResEx)
	result.Total = res.Total
	result.List = make([]subsystem_def.SubsystemExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.Subsystem) subsystem_def.SubsystemExDTO {
	return subsystem_def.SubsystemExDTO{
		// TODO
		Subsystem: v,
	}
}

func (aSrv *adminSrv) GetList(f *subsystem_def.SubsystemQueryForm) ([]model.Subsystem, error) {
	res, err := isubsystem.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.Subsystem, error) {
	res, err := isubsystem.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.Subsystem, error) {
	res, err := isubsystem.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.Subsystem, error) {
	res, err := isubsystem.Srv.GetMulti(pkList)
	return res, err
}

func (aSrv *adminSrv) Add(v *model.Subsystem) (*model.Subsystem, error) {
	err := isubsystem.Srv.Insert(v)
	return v, err
}

func (aSrv *adminSrv) Update(v *model.Subsystem) (int64, error) {
	affected, err := isubsystem.Srv.Update(v)
	return affected, err
}
