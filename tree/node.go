package tree

type Node struct {
	ID       string   `json:"id"`
	Children []*Node  `json:"children,omitempty"`
	Height   int      `json:"height,omitempty"`
	Data     MetaData `json:"metadata,omitempty"`
	ParentID string   `json:"parentId,omitempty"`
	RootID   *string  `json:"rootId,omitempty"`
}

type MetaData struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
}

// GetDescendants returns a count and a hierarchy
// of the subtree (excluding the current node)
func (node *Node) GetDescendants() SubordinatesResponse {
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

	return response
}

func (node *Node) updateHeight(newHeight int) {
	node.Height = newHeight

	for _, children := range node.Children {
		children.updateHeight(node.Height + 1)
	}
}

func (node *Node) countAllDescendants() int {
	var totalCount int

	for _, descendant := range node.Children {
		totalCount += descendant.countAllDescendants() + 1
	}

	return totalCount
}

func (node *Node) removeChildren(childToRemove *Node) {
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
