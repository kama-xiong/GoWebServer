package models

type Page struct {
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Keyword  string `json:"keyword"`
	Desc     bool   `json:"desc"`
}
