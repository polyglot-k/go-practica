package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/polyglot-k/mini-board/internal/db"
	"github.com/polyglot-k/mini-board/internal/post"
)

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
)

func main() {
	// 데이터베이스 연결 초기화
	database := db.NewMySQL()
	defer database.Close()

	// 게시물 테이블이 없으면 생성
	_, err := database.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	postRepo := post.NewRepository(database)

	post.RegisterHandlers(postRepo, templates)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Starting server on http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}