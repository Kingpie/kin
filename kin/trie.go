package kin

import "strings"

type node struct {
	pattern  string  //待匹配路由 如/user/:id
	part     string  //路由中的一部分 如 :id
	children []*node //子节点
	isWild   bool    //是否精准匹配
}

// 第一个匹配成功的节点
// matchChild 根据给定的部分路径匹配子节点
// part: 要匹配的路径部分
// 返回值: 匹配到的子节点指针，如果没有匹配到则返回nil
func (n *node) matchChild(part string) *node {
	// 遍历当前节点的所有子节点
	for _, child := range n.children {
		// 如果子节点的part完全匹配或者子节点是通配符节点，则返回该子节点
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 根据给定的part匹配节点的子节点
// part: 需要匹配的路径部分
// 返回值: 匹配到的子节点切片，包含直接匹配和通配符匹配的节点
func (n *node) matchChildren(part string) []*node {
	// 遍历当前节点的所有子节点，查找匹配的节点
	nodes := make([]*node, 0)
	for _, child := range n.children {
		// 如果子节点的part与给定part相等，或者子节点是通配符节点，则加入结果集
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 将路由模式插入到前缀树节点中
// pattern: 完整的路由模式字符串
// parts: 路由模式按分隔符分割后的字符串切片
func (n *node) insert(pattern string, parts []string) {
	current := n

	// 遍历路径的每一部分
	for i := 0; i < len(parts); i++ {
		part := parts[i]

		// 查找是否已存在匹配的子节点
		var child *node
		for _, c := range current.children {
			if c.part == part || c.isWild {
				child = c
				break
			}
		}

		// 如果不存在匹配的子节点，创建新节点
		if child == nil {
			child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
			current.children = append(current.children, child)
		}

		// 移动到下一个节点
		current = child
	}

	// 到达路径末尾，设置pattern
	current.pattern = pattern
}

// search 在路由树中搜索匹配的节点
// parts: 路由路径按分隔符分割后的字符串切片
// 返回值: 匹配的节点指针，如果未找到则返回nil
func (n *node) search(parts []string) *node {
	current := n

	// 遍历路径的每一部分
	for i := 0; i < len(parts); i++ {
		part := parts[i]

		// 如果当前节点是通配符节点，则直接返回（如果已有pattern）
		if strings.HasPrefix(current.part, "*") {
			if current.pattern != "" {
				return current
			}
			return nil
		}

		// 查找匹配的子节点
		child := current.matchChild(part)
		if child == nil {
			return nil
		}

		// 移动到下一个节点
		current = child
	}

	// 检查最后一个节点是否有完整的路由模式
	if current.pattern != "" {
		return current
	}

	return nil
}
