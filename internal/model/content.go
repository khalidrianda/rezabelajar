package model

type Core struct {
	ID     uint
	Author string
	Title  string
	Body   string
}

type Article struct {
	ID     uint   `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
