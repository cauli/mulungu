package tree

type Node struct {
	ID       string   `json:"id"`
	Children []*Node  `json:"children,omitempty"`
	Depth    int      `json:"depth,omitEmpty"`
	Data     MetaData `json:"metadata,omitempty"`
	ParentID string   `json:"parentId,omitEmpty"`
}

type MetaData struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
}

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
		Subordinates: SubordinatesInfo{
			Count: SubordinatesCount{
				Direct: directCount,
				Total:  totalCount,
			},
			Hierarchy: node.Children,
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
