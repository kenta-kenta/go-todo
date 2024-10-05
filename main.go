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
	r.HandleFunc("/todo", handler.PostTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todo/list", handler.TodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo/{id:[0-9]+}", handler.UpdateTodoHandler).Methods(http.MethodPatch)
	r.HandleFunc("/todo/{id:[0-9]+}", handler.DeleteTodoHandler).Methods(http.MethodDelete)

	// サーバー起動時のログ
	log.Println("server start at port 8080")
	// サーバー起動(ルータを明示的に指定)
	log.Fatal(http.ListenAndServe(":8080", r))
}
