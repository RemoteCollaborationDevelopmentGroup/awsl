package main

import (
	"strings"
	"strconv"
	"net/http"
	"encoding/binary"
	"github.com/syndtr/goleveldb/leveldb"
)

var database_user *leveldb.DB
var database_task *leveldb.DB

func init() {
	var err error
	database_user, err = leveldb.OpenFile("data/database_user", nil)
	if err != nil { panic(err) }
	database_task, err = leveldb.OpenFile("data/database_task", nil)
	if err != nil { panic(err) }
}

func main() {
	mux := http.NewServeMux()
	//mux.Handle("/", &session{})
	mux.Handle("/user/", &session{database_user})
	mux.Handle("/task/", &session{database_task})
	http.ListenAndServe(":8888", mux)
}

type session struct {
	database *leveldb.DB
}

func (s *session) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Authorization")
		w.WriteHeader(200)
	case "GET":
		data, err := s.database.Get(URLToID(r.URL.Path), nil)
		if err != nil {
			http.Error(w, "目标资源不存在", 404)
			return
		}
		w.Write(data)
	case "PUT":
		w.Write([]byte("leveldb data"))
	}
}

//func (s *session) authentication() {}
//func (s *session) authorization() {}

func URLToID(url string) []byte {
	sp := strings.Split(url, "/")
	if len(sp) != 3 {
		return []byte{}
	}
	if num, err := strconv.Atoi(sp[2]); err != nil {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(num))
		return buf
	}
	return []byte{}
}

