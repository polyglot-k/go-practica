package post

import (
	"html/template"
	"net/http"
)

// RegisterHandlers는 기본 ServeMux에 게시물 핸들러를 등록합니다.
func RegisterHandlers(repo *Repository, tpl *template.Template) {
	h := NewHandler(repo, tpl)

	http.HandleFunc("/", h.ListPostsHandler)
	http.HandleFunc("/posts/new", h.NewPostHandler)
	http.HandleFunc("/posts/create", h.CreatePostHandler)
	http.HandleFunc("/posts/edit", h.EditPostHandler)
	http.HandleFunc("/posts/update", h.UpdatePostHandler)
	http.HandleFunc("/posts/delete", h.DeletePostHandler)
}
