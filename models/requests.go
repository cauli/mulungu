package models

import "github.com/cauli/mulungu/tree"

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Leader string `json:"leader"`
}

func (employee *Employee) CreateNode() *tree.Node {
	return &tree.Node{
		ID: employee.ID,
		Data: tree.MetaData{
			Name:  employee.Name,
			Title: employee.Title,
		},
		ParentID: employee.Leader,
	}
}
