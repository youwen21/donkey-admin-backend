package menutree

import (
	"gofly/app/model"
	"sort"
)

// ToTree 严格模式：将菜单列表转换为树结构
// 从根节点（ParentId为0）开始构建单树
func ToTree(menus []model.Menu) []*TreeMenu {
	if len(menus) == 0 {
		return nil
	}

	// 创建父ID到子菜单列表的映射
	childrenMap := make(map[int][]model.Menu)
	for _, menu := range menus {
		parentId := menu.ParentId
		childrenMap[parentId] = append(childrenMap[parentId], menu)
	}

	// 递归构建树
	var buildTree func(parentId int) []*TreeMenu
	buildTree = func(parentId int) []*TreeMenu {
		children, exists := childrenMap[parentId]
		if !exists || len(children) == 0 {
			return nil
		}

		var treeNodes []*TreeMenu
		for _, child := range children {
			node := &TreeMenu{
				Menu:     child,
				Children: buildTree(int(child.Id)),
			}
			treeNodes = append(treeNodes, node)
		}

		// 排序
		sort.Slice(treeNodes, func(i, j int) bool {
			return treeNodes[i].OrderNo < treeNodes[j].OrderNo
		})

		return treeNodes
	}

	// 从根节点（ParentId为0）开始构建树
	return buildTree(0)
}

// ToTreeLoose 宽松模式：将菜单列表转换为树结构（支持多树）
// 自动识别所有根节点（ParentId在数据中不存在的节点），返回多棵树
func ToTreeLoose(menus []model.Menu) []*TreeMenu {
	if len(menus) == 0 {
		return nil
	}

	// 创建ID到菜单的映射，用于快速查找
	idMap := make(map[int]model.Menu)
	for _, menu := range menus {
		idMap[int(menu.Id)] = menu
	}

	// 创建父ID到子菜单列表的映射
	childrenMap := make(map[int][]model.Menu)
	for _, menu := range menus {
		parentId := menu.ParentId
		childrenMap[parentId] = append(childrenMap[parentId], menu)
	}

	// 找出所有根节点
	// 根节点：ParentId为0，或者ParentId不在数据中的节点
	var rootNodes []model.Menu
	for _, menu := range menus {
		parentId := menu.ParentId
		// ParentId为0，或者ParentId不在数据中，则为根节点
		if parentId == 0 {
			rootNodes = append(rootNodes, menu)
		} else if _, exists := idMap[parentId]; !exists {
			rootNodes = append(rootNodes, menu)
		}
	}

	// 递归构建树
	var buildTree func(parentId int) []*TreeMenu
	buildTree = func(parentId int) []*TreeMenu {
		children, exists := childrenMap[parentId]
		if !exists || len(children) == 0 {
			return nil
		}

		var treeNodes []*TreeMenu
		for _, child := range children {
			node := &TreeMenu{
				Menu:     child,
				Children: buildTree(int(child.Id)),
			}
			treeNodes = append(treeNodes, node)
		}

		// 排序
		sort.Slice(treeNodes, func(i, j int) bool {
			return treeNodes[i].OrderNo < treeNodes[j].OrderNo
		})

		return treeNodes
	}

	// 为每个根节点构建树
	var result []*TreeMenu
	for _, root := range rootNodes {
		node := &TreeMenu{
			Menu:     root,
			Children: buildTree(int(root.Id)),
		}
		result = append(result, node)
	}

	// 对根节点排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].OrderNo < result[j].OrderNo
	})

	return result
}

// TreeToMenuListDFS 将树结构转换为菜单列表（深度优先遍历）
// 优先遍历到最深层，再返回处理同级节点
func TreeToMenuListDFS(tree []*TreeMenu) []model.Menu {
	if len(tree) == 0 {
		return nil
	}

	var result []model.Menu

	// 递归函数：深度优先遍历
	var dfs func(nodes []*TreeMenu)
	dfs = func(nodes []*TreeMenu) {
		for _, node := range nodes {
			// 先添加当前节点
			result = append(result, node.Menu)
			// 然后递归处理子节点（深度优先）
			if len(node.Children) > 0 {
				dfs(node.Children)
			}
		}
	}

	dfs(tree)
	return result
}

// TreeToMenuListBFS 将树结构转换为菜单列表（广度优先遍历）
// 优先遍历同级节点，再遍历下一层
func TreeToMenuListBFS(tree []*TreeMenu) []model.Menu {
	if len(tree) == 0 {
		return nil
	}

	var result []model.Menu
	// 使用队列进行广度优先遍历
	queue := make([]*TreeMenu, 0)

	// 将根节点加入队列
	queue = append(queue, tree...)

	// 队列不为空时继续处理
	for len(queue) > 0 {
		// 取出队列头部节点
		node := queue[0]
		queue = queue[1:]

		// 添加当前节点到结果
		result = append(result, node.Menu)

		// 将当前节点的所有子节点加入队列尾部
		if len(node.Children) > 0 {
			queue = append(queue, node.Children...)
		}
	}

	return result
}
