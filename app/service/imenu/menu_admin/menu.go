package menu_admin

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/imenu"
	"gofly/app/service/imenu/menu_def"
	"gofly/app/service/tree/menutree"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *menu_def.MenuQueryForm) (*menu_def.MenuQueryResEx, error) {
	res, err := imenu.Srv.Query(f)
	if err != nil {
		return nil, err
	}
	// biz process
	result := new(menu_def.MenuQueryResEx)
	result.Total = res.Total
	result.List = make([]menu_def.MenuExDTO, len(res.List))

	tree := menutree.ToTree(res.List)
	menuList := menutree.TreeToMenuListDFS(tree)

	for i, v := range menuList {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.Menu) menu_def.MenuExDTO {
	return menu_def.MenuExDTO{
		// TODO
		Menu: v,
	}
}

func (aSrv *adminSrv) GetList(f *menu_def.MenuQueryForm) ([]model.Menu, error) {
	res, err := imenu.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.Menu, error) {
	res, err := imenu.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.Menu, error) {
	res, err := imenu.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.Menu, error) {
	res, err := imenu.Srv.GetMulti(pkList)
	return res, err
}

func (aSrv *adminSrv) Add(v *model.Menu) (*model.Menu, error) {
	if v.ParentId == 0 {
		v.NodePath = "/"

	} else {
		parent, err := aSrv.Get(v.ParentId)
		if err != nil {
			return nil, err
		}
		v.NodePath = fmt.Sprintf("%s%v/", parent.NodePath, parent.Id)
		v.Level = parent.Level + 1
	}

	err := imenu.Srv.Insert(v)
	return v, err
}

func (aSrv *adminSrv) Update(v *model.Menu) (int64, error) {
	if v.ParentId == 0 {
		v.NodePath = "/"
	} else {
		parent, err := aSrv.Get(v.ParentId)
		if err != nil {
			return 0, err
		}
		v.NodePath = fmt.Sprintf("%s%v/", parent.NodePath, parent.Id)
		v.Level = parent.Level + 1
	}

	affected, err := imenu.Srv.Update(v)
	return affected, err
}
