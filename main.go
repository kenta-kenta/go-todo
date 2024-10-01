package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenta-kenta/go-todo/handler"
)

func main() {
	// ハンドラ：HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む関数
	// ルータを明示的に宣言
	r := mux.NewRouter()
	// ハンドラをサーバーで使用するように登録する
	r.HandleFunc("/hello", handler.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handler.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handler.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handler.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handler.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handler.PostCommentHandler).Methods(http.MethodPost)

	// サーバー起動時のログ
	log.Println("server start at port 8080")
	// サーバー起動(ルータを明示的に指定)
	log.Fatal(http.ListenAndServe(":8080", r))
}
