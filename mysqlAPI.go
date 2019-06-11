package main
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	// "io/ioutil"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)
var db,err=sql.Open("mysql","saggarwal98:shubham@tcp(127.0.0.1:3306)/mysqlAPI")
type Article struct {
	ID    int  `json:"ID"`
	Title string `json:"Title"`
	Description  string `json:"Description"`
	Price int32  `json:"Price"`
}
func homeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the homepage")
}
func varticles(w http.ResponseWriter, r *http.Request) {
	results,err:=db.Query("Select * from Articles")
	for results.Next(){
		var a Article
		err=results.Scan(&a.ID,&a.Title,&a.Description,&a.Price)
		if err != nil {
            log.Println(err.Error())
		}
		fmt.Fprintf(w,"ID:%v Title:%v Description:%v Price:%v \n",a.ID,a.Title,a.Description,a.Price)
	}
}
func varticlesid(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["ID"]
	flag := false
	results,err:=db.Query("Select * from Articles")
	for results.Next(){
		var a Article
		err=results.Scan(&a.ID,&a.Title,&a.Description,&a.Price)
		if err != nil {
            log.Println(err.Error())
		}
		if strconv.Itoa(a.ID)== key{
			fmt.Fprintf(w,"ID:%v Title:%v Description:%v Price:%v \n",a.ID,a.Title,a.Description,a.Price)
			flag = true
		}
	}
	if flag == false {
		fmt.Fprint(w, "Sorry, no article found with that id")
	}
}
func carticles(w http.ResponseWriter,r *http.Request){
	key:=mux.Vars(r)["ID"]
	key1:=mux.Vars(r)["Title"]
	key2:=mux.Vars(r)["Description"]
	key3:=mux.Vars(r)["Price"]
	_,err:=db.Query("INSERT INTO Articles VALUES('"+key+"','"+key1+"','"+key2+"','"+key3+"')")
	if err!=nil{
		fmt.Fprint(w,"Could not create Article")
		log.Fatal(err.Error())
	}
}
func delarticles(w http.ResponseWriter,r *http.Request){
	key:=mux.Vars(r)["Title"]
	flag:=false
	results,err:=db.Query("Select * from Articles")
	for results.Next(){
		var a Article
		err=results.Scan(&a.ID,&a.Title,&a.Description,&a.Price)
		if err != nil {
            log.Print(err.Error())
		}
		if a.Title== key{
			fmt.Fprintf(w,"ID:%v Title:%v Description:%v Price:%v \n",a.ID,a.Title,a.Description,a.Price)
			flag = true
			var str ="Delete from Articles where Title='"+a.Title+"'"
			fmt.Println(str)
			_,err:=db.Query(str)
			if err!=nil{
				log.Println(err.Error())
			}
			fmt.Fprintf(w,"Article Deleted")
		}
	}
	if flag==false{
		fmt.Fprint(w,"No article found with that title")
	}
}
func updarticles(w http.ResponseWriter,r *http.Request){
	key:=mux.Vars(r)["ID"]
	Value:=mux.Vars(r)["Description"]
	flag:=false
	results,err:=db.Query("Select * from Articles")
	for results.Next(){
		var a Article
		err=results.Scan(&a.ID,&a.Title,&a.Description,&a.Price)
		if err != nil {
            log.Print(err.Error())
		}
		if strconv.Itoa(a.ID)== key{
			flag = true
			var str ="update Articles set Title='"+Value+"' where ID='"+strconv.Itoa(a.ID)+"'"
			fmt.Println(str)
			_,err:=db.Query(str)
			if err!=nil{
				log.Println(err.Error())
			}
			fmt.Fprintf(w,"Article Updated")
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
	myRouter.HandleFunc("/varticles/{ID}/{Title}/{Description}/{Price}", carticles).Methods("POST")
	myRouter.HandleFunc("/varticles/{ID}", varticlesid).Methods("GET")
	myRouter.HandleFunc("/darticles/{Title}",delarticles).Methods("DELETE")
	myRouter.HandleFunc("/uarticles/{ID}/{Description}",updarticles).Methods("PUT")
	log.Fatal(http.ListenAndServe(":4000", myRouter))
}
func main() {
	handleFunction()
}