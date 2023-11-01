package model

type HeaderData struct {
	Token   string `header:"token"`
	Version string `header:"version" binding:"required"`
}

type PageRequest struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

type PageData struct {
	Total       int         `json:"total"`
	TotalPage   int         `json:"total_page"`
	CurrentPage uint8       `json:"current_page"`
	PageSize    uint8       `json:"page_size"`
	List        interface{} `json:"list"`
}
