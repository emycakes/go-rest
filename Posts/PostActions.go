package Posts

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	App "gorest/Application"
	"strconv"
	"time"
)

func GetAllPosts(request App.Request, response App.Response) {
	response.SendAsJson(200, GetPosts())
}

func FindPost(request App.Request, response App.Response) {
	id, _ := strconv.Atoi(request.Get("id"))

	/// todo: should never fetch all posts before searching
	post, _ := GetPosts().FindById(id)

	if post.IsBlank() {
		response.SendAsJson(404, nil)
		return
	}

	response.SendAsJson(200, post)
}

func CreatePost(request App.Request, response App.Response) {
	var post Post
	json.Unmarshal(request.GetBody(), &post)

	post.Slug = slug.Make(post.Title)
	post.UserUuid = uuid.New().String()
	post.CreatedAt = time.Now().UTC()

	response.SendAsJson(200, append(GetPosts(), post))
}

func UpdatePost(request App.Request, response App.Response) {
	id, _ := strconv.Atoi(request.Get("id"))

	posts := GetPosts()
	post, index := posts.FindById(id)

	if post.IsBlank() {
		response.SendAsJson(404, nil)
	}

	json.Unmarshal(request.GetBody(), &post)
	post.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	// to be replaced with actual storage code
	posts[index] = post

	response.SendAsJson(200, posts)
}

func DeletePost(request App.Request, response App.Response) {
	id, _ := strconv.Atoi(request.Get("id"))

	posts := GetPosts()
	_, index := posts.FindById(id)

	if index == -1 {
		response.SendAsJson(404, nil)
	}

	posts = append(posts[:index], posts[index + 1:]...)

	response.SendAsJson(200, posts)
}
