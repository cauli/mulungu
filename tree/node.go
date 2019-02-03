package tree

func (node *Node) RemoveChildren(childToRemove *Node) {
	indexToRemove := -1

	for index, child := range node.Children {
		if child.ID == childToRemove.ID {
			indexToRemove = index
			childToRemove.ParentID = ""
			break
		}
	}

	if indexToRemove != -1 {
		node.Children = append(node.Children[:indexToRemove], node.Children[indexToRemove+1:]...)

		if len(node.Children) == 0 {
			node.Children = nil
		}
	}
}

func (node *Node) GetDescendants() (SubordinatesResponse, error) {
	directCount := len(node.Children)

	totalCount := 0
	for _, descendant := range node.Children {
		totalCount += descendant.countAllDescendants() + 1
	}

	response := SubordinatesResponse{
		subordinates: SubordinatesInfo{
			count: SubordinatesCount{
				direct: directCount,
				total:  totalCount,
			},
			hierarchy: node.Children,
		},
	}

	return response, nil
}

func (node *Node) countAllDescendants() int {
	var totalCount int

	for _, descendant := range node.Children {
		totalCount += descendant.countAllDescendants() + 1
	}

	return totalCount
}
