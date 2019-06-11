package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"io/ioutil"
)

//Article Structure defined
type Article struct {
	ID    int  `json:"ID"`
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
	Price int32  `json:"Price"`
}

//Articles is an array
type Articles []Article

//articles is defined
var articles = Articles{
	Article{ID: 101, Title: "Article1", Desc: "This is article 1", Price: 10000},
	Article{ID: 102, Title: "Article2", Desc: "This is article 2", Price: 20000},
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the homepage")
}
func varticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}
func varticlesid(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["ID"]
	// fmt.Fprintf(w,"Key: "+mux.Vars(r)["ID"])
	flag := false
	for _, article := range articles {
		if strconv.Itoa(article.ID)== key{
			json.NewEncoder(w).Encode(article)
			flag = true
		}
	}
	if flag == false {
		fmt.Fprint(w, "Sorry, no article found with that id")
	}
}
func carticles(w http.ResponseWriter,r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle Article
	json.Unmarshal(reqBody,&newArticle)
	articles = append(articles,newArticle)
	json.NewEncoder(w).Encode(newArticle)
	// fmt.Fprintf(w,"Test worked")
}
func delarticles(w http.ResponseWriter,r *http.Request){
	key:=mux.Vars(r)["Title"]
	flag:=false
	for index,article:=range articles{
		if article.Title==key{
			articles=append(articles[:index],articles[index+1:]...)
			flag=true
			fmt.Fprint(w,"Article Deleted")
		}
	}
	if flag==false{
		fmt.Fprint(w,"No article found with that title")
	}
}
func updarticles(w http.ResponseWriter,r *http.Request){
	key:=mux.Vars(r)["ID"]
	Value:=mux.Vars(r)["Desc"]
	flag:=false
	for index,article:=range articles{
		if strconv.Itoa(article.ID)==key{
			articles=append(articles[:index],articles[index+1:]...)
			flag=true
			article.Title=Value
			articles = append(articles,article)
			fmt.Fprint(w,"Article updated")
		}
	}
	if flag==false{
		fmt.Fprint(w,"No article found with that id")
	}
}
func handleFunction() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homeFunc)
	myRouter.HandleFunc("/varticles", varticles).Methods("GET")
	myRouter.HandleFunc("/varticles", carticles).Methods("POST")
	myRouter.HandleFunc("/varticles/{ID}", varticlesid).Methods("GET")
	myRouter.HandleFunc("/darticles/{Title}",delarticles).Methods("DELETE")
	myRouter.HandleFunc("/uarticles/{ID}/{Desc}",updarticles).Methods("PUT")
	log.Fatal(http.ListenAndServe(":4000", myRouter))
}
func main() {
	handleFunction()
}
