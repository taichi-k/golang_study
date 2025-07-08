package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// User は users テーブルのモデル
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func main() {
	// 1) DB 接続
	// DSN: [user[:password]@][protocol[(address)]]/dbname[?param1=value1]
	dsn := "root:pass@tcp(127.0.0.1:3333)/gotest?parseTime=true"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB 接続失敗: %v", err)
	}
	defer db.Close()

	// 接続プール設定の例：最大オープン接続数 = CPUコア数 × 2
	cpu := runtime.NumCPU()
	db.SetMaxOpenConns(cpu * 2) // 例: 4コアなら 8
	db.SetMaxIdleConns(cpu)     // 例: 4

	// 2) ルーター準備
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")

	// 3) サーバ起動
	addr := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

// getUsers: 全ユーザ取得
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// createUser: ユーザ作成
func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", u.Name, u.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	u.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
