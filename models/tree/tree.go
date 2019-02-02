package tree

type Tree struct {
	id   string `json:"id"`
	root *Node  `json:"root"`
}

type Node struct {
	data     MetaData `json:"metadata"`
	children *[]Node  `json:"children"`
	parent   *Node    `json:"parent"`
}

type MetaData struct {
	name  string `json:"name"`
	title string `json:"title"`
}

func Create(treeId string) (*Tree, error) {
	rootNode := Node{
		data: MetaData{
			name:  "#1",
			title: "Founder",
		},
	}

	tree := &Tree{
		id:   treeId,
		root: &rootNode,
	}

	return tree, nil
}
