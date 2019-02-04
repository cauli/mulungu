package models

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Leader string `json:"leader"`
}
