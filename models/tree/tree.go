package models

type Node struct {
	data     MetaData `json:"metadata"`
	children *[]Node  `json:"children"`
	parent   *Node    `json:"parent"`
}

type MetaData struct {
	name  string `json:"name"`
	title string `json:"title"`
}

type Tree struct {
	id   string `json:"id"`
	root *Node  `json:"root"`
}

func Create(id string) (*Tree, error) {
	rootNode := Node{
		data: MetaData{
			name:  "#1",
			title: "Founder",
		},
	}

	tree := &Tree{
		root: &rootNode,
	}

	return tree, nil
}
