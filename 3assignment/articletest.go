package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
)
type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article

func allarticles (w http.Responcewrites,r * http.Request){

	articles := articles
	{
        Article{Title: "Hello", Desc: "sathya", Content: "practise golang"},
        Article{Title: "Hello", Desc: "sathyabama", Content: "practise golang api testing"},	
	}
	fmt.Println("Endpoint Hit: homePage")
	json.NewEncoder(w).Encode(articles)
}


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, " hai sathya welcome to the HomePage!")
    
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles",allarticles)
    log.Fatal(http.ListenAndServe(":80000", nil))
}
func main() {
    
    handleRequests()
}
