package model

type Actor struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Brief string `json:"brief"`
	Job   string `json:"job"`
	Work  string `json:"work"`
}
