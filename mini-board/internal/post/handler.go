package post

import (
	"html/template"
	"net/http"
	"strconv"
)

// Handler는 게시물 핸들러에 대한 종속성을 보유합니다.
type Handler struct {
	Repository *Repository
	Templates  *template.Template
}

// NewHandler는 새 핸들러를 만듭니다.
func NewHandler(repo *Repository, tpl *template.Template) *Handler {
	return &Handler{
		Repository: repo,
		Templates:  tpl,
	}
}

// ListPostsHandler는 모든 게시물을 나열하는 요청을 처리합니다.
func (h *Handler) ListPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Repository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Templates.ExecuteTemplate(w, "index.html", posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NewPostHandler는 새 게시물 양식을 표시하는 요청을 처리합니다.
func (h *Handler) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Templates.ExecuteTemplate(w, "new.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreatePostHandler는 새 게시물을 만드는 요청을 처리합니다.
func (h *Handler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	p := &Post{
		Title:   title,
		Content: content,
	}
	if err := h.Repository.Create(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// EditPostHandler는 게시물 수정 양식을 표시하는 요청을 처리합니다.
func (h *Handler) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	p, err := h.Repository.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if p == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	err = h.Templates.ExecuteTemplate(w, "edit.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdatePostHandler는 게시물을 업데이트하는 요청을 처리합니다.
func (h *Handler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	p := &Post{
		ID:      id,
		Title:   title,
		Content: content,
	}

	if err := h.Repository.Update(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// DeletePostHandler는 게시물을 삭제하는 요청을 처리합니다.
func (h *Handler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := h.Repository.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
