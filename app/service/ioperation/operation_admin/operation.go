package operation_admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/ioperation"
	"donkey-admin/app/service/ioperation/operation_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *operation_def.OperationQueryForm) (*operation_def.OperationQueryResEx, error) {
	res, err := ioperation.Srv.Query(f)
	// biz process
	result := new(operation_def.OperationQueryResEx)
	result.Total = res.Total
	result.List = make([]operation_def.OperationExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.Operation) operation_def.OperationExDTO {
	return operation_def.OperationExDTO{
		// TODO
		Operation: v,
	}
}

func (aSrv *adminSrv) GetList(f *operation_def.OperationQueryForm) ([]model.Operation, error) {
	res, err := ioperation.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.Operation, error) {
	res, err := ioperation.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.Operation, error) {
	res, err := ioperation.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.Operation, error) {
	res, err := ioperation.Srv.GetMulti(pkList)
	return res, err
}
