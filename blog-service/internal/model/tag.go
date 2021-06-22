package model

import "github.com/go-programing-tour-book/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint   `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Paper app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
