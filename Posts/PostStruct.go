package Posts

import "time"

type Post struct {
	Id int `json:"id"`
	Slug string `json:"slug"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserUuid string `json:"user_uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

type PostCollection []Post

func (post Post) IsBlank() bool {
	return post.Title == "" && post.Content == ""
}

func (posts PostCollection) FindById(id int) (Post, int) {
	for index, post := range posts {
		if post.Id == id {
			return post, index
		}
	}

	return Post{}, -1
}

// GetPosts mocked up data for testing
func GetPosts() PostCollection {
	creationDate, _ := time.Parse(time.RFC3339, "2021-07-01T14:00:00Z")

	return []Post{
		{
			Id: 1,
			Slug: "hello",
			Title: "Hello",
			Content: "Post Content",
			UserUuid: "d4ed6815-d23b-4c8f-ba83-09acafdda573",
			CreatedAt: creationDate,
			UpdatedAt: nil,
			DeletedAt: nil,
		},
		{
			Id: 2,
			Slug: "hello-2",
			Title: "Hello 2",
			Content: "Post Content",
			UserUuid: "00bf7918-6a1e-41ce-ad76-0f91a11072a0",
			CreatedAt: creationDate,
			UpdatedAt: nil,
			DeletedAt: nil,
		},
	}
}
