package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

const port = 3000

func read() int {
	data, err := ioutil.ReadFile("count")
	if err != nil {
		panic(err)
	}
	count, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}
	return count
}

func main() {

	http.HandleFunc("/add",
		func(w http.ResponseWriter, r *http.Request) {
			count := read()
			count++
			ioutil.WriteFile("count", []byte(strconv.Itoa(count)), 0644)
			http.Redirect(w, r, "/", http.StatusFound)
		})

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			var tmpl = template.Must(template.ParseFiles("index.tmpl"))
			tmpl.Execute(w, read())
		})

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

}
