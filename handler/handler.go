package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/kenta-kenta/go-todo/models"
)

// ハンドラの定義は関数宣言する
// ハンドラ名は大文字から始める

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world\n")
}

func PostTodoHandler(w http.ResponseWriter, req *http.Request) {
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)

	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	var reqTodo models.Todo
	if err := json.Unmarshal(reqBodybuffer, &reqTodo); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	todo := reqTodo

	jsonData, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func TodoListHandler(w http.ResponseWriter, req *http.Request) {
	todo := []models.Todo{models.Todo1, models.Todo2}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func UpdateTodoHandler(w http.ResponseWriter, req *http.Request) {
	// todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	// if err != nil {
	// 	http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	// 	return
	// }

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)

	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	var reqTodo models.Todo
	if err := json.Unmarshal(reqBodybuffer, &reqTodo); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	todo := reqTodo
	jsonData, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "delete todo...")
}
