package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"` //评论人id
	User    User   `json:"user"`                    //评论人信息
	PostID  uint   `gorm:"not null" json:"post_id"` //文章id
	Post    Post   `json:"-"`                       //文章信息
}
