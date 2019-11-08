package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"
)

func getUUID() (string, error) {
	return uuid.New().String(), nil
}

func hoge() string {
	return "hoge"
}

const (
	testName = "test"
)

func main() {
	piyo()
	db := sync.Map{}
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var resJSON map[string]interface{}
		err = json.Unmarshal(body, &resJSON)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pp.Print(resJSON)
		log.Println("res json")
		pp.Print(resJSON)
		log.Println("data ")
		pp.Print(resJSON["data"])

		_, ok := resJSON["data"].([]interface{})
		if !ok {
			log.Println("d json id is bad")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "ok")
		return
	})

	http.HandleFunc("/pay", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		user := query.Get("user")

		_, ok := db.Load(user)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "user is not exist")
			return
		}
		//pp.Print(resJSON)
		//d, ok := resJSON["data"].([]interface{})
		//if !ok {
		//	log.Println("d json id is bad")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

		// data, ok := d[0].(map[string]interface{})
		// if !ok {
		// 	log.Println("data json id is bad")
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		// fmt.Fprintf(w, "you paied %f", data["amount"])
		return
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		panic(err)
	}
}

func piyo() error {
	return nil
}
