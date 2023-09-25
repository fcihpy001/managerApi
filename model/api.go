package model

type HeaderData struct {
	Token   string `header:"token"`
	Version string `header:"version" binding:"required"`
}
