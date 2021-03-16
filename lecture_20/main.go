package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/mateors/mcb"
)

var db *mcb.DB
var err error

// same as mysql table schema/structure
type RequestTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
}

func init() {
	// Open up our database connection.

	// Couchbase Database Connection Part

	db = mcb.Connect("localhost", "fahim047", "904977")

	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func features(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func request(w http.ResponseWriter, r *http.Request) {

	// for inserting data into couchbase bucket

	r.ParseForm()

	for key, val := range r.Form {

		fmt.Println(key, val)
	}
	var reqTable RequestTable

	r.Form.Add("bucket", "master_academy")
	r.Form.Add("aid", "request::5")
	r.Form.Add("type", "request")
	r.Form.Add("status", "1")
	pRes := db.Insert(r.Form, &reqTable)
	fmt.Println(pRes.Status, pRes.Errors)

	fmt.Fprintf(w, `OK`)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/request", request)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}
