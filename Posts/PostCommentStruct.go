package Posts

import "time"

type PostComment struct{
	Content string `json:"content"`
	Uuid string `json:"uuid"`
	PostUuid string `json:"post_uuid"`
	UserUuid string `json:"user_uuid"`
	Likes int `json:"likes"`
	Dislikes int `json:"dislikes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

type PostCommentCollection []PostComment