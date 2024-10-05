package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ハンドラの定義は関数宣言する
// ハンドラ名は大文字から始める

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world\n")
}

func PostTodoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "posting todo")
}

func TodoListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "todo list ...")
}

func UpdateTodoHandler(w http.ResponseWriter, req *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Todo No.%d\n", todoID)
	io.WriteString(w, resString)
}

func DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "delete todo...")
}
