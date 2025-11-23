package organization_admin

import (
	"fmt"
	"gofly/app/model"
	"gofly/app/service/iorganization"
	"gofly/app/service/iorganization/organization_def"
	"gofly/app/service/tree/orgtree"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *organization_def.OrganizationQueryForm) (*organization_def.OrganizationQueryResEx, error) {
	res, err := iorganization.Srv.Query(f)
	if err != nil {
		return nil, err
	}
	// biz process
	result := new(organization_def.OrganizationQueryResEx)
	result.Total = res.Total

	tree := orgtree.ToTree(res.List)
	orgList := orgtree.TreeToOrgListDFS(tree)

	result.List = make([]organization_def.OrganizationExDTO, len(orgList))
	for i, v := range orgList {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.Organization) organization_def.OrganizationExDTO {
	return organization_def.OrganizationExDTO{
		// TODO
		Organization: v,
	}
}

func (aSrv *adminSrv) GetList(f *organization_def.OrganizationQueryForm) ([]model.Organization, error) {
	res, err := iorganization.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.Organization, error) {
	res, err := iorganization.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.Organization, error) {
	res, err := iorganization.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.Organization, error) {
	res, err := iorganization.Srv.GetMulti(pkList)
	return res, err
}

func (aSrv *adminSrv) Add(v *model.Organization) (*model.Organization, error) {
	if v.ParentId == 0 {
		v.NodePath = "/"
	} else {
		parentOrg, err := aSrv.Get(v.ParentId)
		if err != nil {
			return nil, err
		}
		v.NodePath = fmt.Sprintf("%s%v/", parentOrg.NodePath, parentOrg.Id)
		v.Level = parentOrg.Level + 1
	}

	err := iorganization.Srv.Insert(v)
	return v, err
}

func (aSrv *adminSrv) Update(v *model.Organization) (int64, error) {
	if v.ParentId == 0 {
		v.NodePath = "/"
	} else {
		parentOrg, err := aSrv.Get(v.ParentId)
		if err != nil {
			return 0, err
		}
		v.NodePath = fmt.Sprintf("%s%v/", parentOrg.NodePath, parentOrg.Id)
		v.Level = parentOrg.Level + 1
	}

	affected, err := iorganization.Srv.Update(v)
	return affected, err
}

// sortOrgTree 对组织列表进行树形排序
// 排序规则：
// 1. 按照 path 层级深度排序（浅的在前）
// 2. 同级节点按照 order_no 排序
// 3. 同一父节点下的子节点紧跟在父节点后面
//func (aSrv *adminSrv) sortOrgTree(list []org_def.OrgExDTO) []org_def.OrgExDTO {
//	if len(list) == 0 {
//		return list
//	}
//
//	// 构建 parent_id 到子节点列表的映射
//	childrenMap := make(map[int][]*org_def.OrgExDTO)
//	var rootNodes []*org_def.OrgExDTO
//
//	for i := range list {
//		node := &list[i]
//		if node.ParentId == 0 {
//			// 根节点（parent_id 为 0）
//			rootNodes = append(rootNodes, node)
//		} else {
//			// 子节点
//			childrenMap[node.ParentId] = append(childrenMap[node.ParentId], node)
//		}
//	}
//
//	// 对每个父节点下的子节点按 order_no 排序
//	for parentId := range childrenMap {
//		children := childrenMap[parentId]
//		sort.Slice(children, func(i, j int) bool {
//			// 先按 order_no 排序
//			if children[i].OrderNo != children[j].OrderNo {
//				return children[i].OrderNo < children[j].OrderNo
//			}
//			// order_no 相同时，按 path 字典序排序
//			return children[i].NodePath < children[j].NodePath
//		})
//	}
//
//	// 对根节点排序
//	sort.Slice(rootNodes, func(i, j int) bool {
//		// 先按 order_no 排序
//		if rootNodes[i].OrderNo != rootNodes[j].OrderNo {
//			return rootNodes[i].OrderNo < rootNodes[j].OrderNo
//		}
//		// order_no 相同时，按 path 字典序排序
//		return rootNodes[i].NodePath < rootNodes[j].NodePath
//	})
//
//	// 深度优先遍历，构建排序后的列表
//	var sortedList []org_def.OrgExDTO
//	var dfs func(node *org_def.OrgExDTO)
//	dfs = func(node *org_def.OrgExDTO) {
//		// 添加当前节点
//		sortedList = append(sortedList, *node)
//		// 递归处理子节点
//		if children, ok := childrenMap[node.Id]; ok {
//			for _, child := range children {
//				dfs(child)
//			}
//		}
//	}
//
//	// 从根节点开始遍历
//	for _, root := range rootNodes {
//		dfs(root)
//	}
//
//	return sortedList
//}
