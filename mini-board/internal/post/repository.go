package post

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(p *Post) error {
	result, err := r.db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", p.Title, p.Content)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = id
	return nil
}

func (r *Repository) FindAll() ([]Post, error) {
	rows, err := r.db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *Repository) FindByID(id int64) (*Post, error) {
	row := r.db.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id)
	var p Post
	if err := row.Scan(&p.ID, &p.Title, &p.Content); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Update(p *Post) error {
	_, err := r.db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", p.Title, p.Content, p.ID)
	return err
}

func (r *Repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}
