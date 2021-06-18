package model

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programing-tour-book/blog-service/pkg/app"
	"github.com/go-programing-tour-book/blog-service/pkg/errcode"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint   `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	// return
}
