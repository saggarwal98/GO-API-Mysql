package main
import(
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)
type tag struct{
	Name string `json:"name"`
}
func main(){
	db,err:=sql.Open("mysql","saggarwal98:shubham@tcp(127.0.0.1:3306)/testdb")
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Database connected")
	insert,err:=db.Query("INSERT INTO users VALUES('shubham')")
	if err!=nil{
		log.Fatal(err.Error())
	}
	defer insert.Close()
	fmt.Println("Values inserted")
	results,err:=db.Query("Select * from users")
	for results.Next(){
		var t tag
		err=results.Scan(&t.Name)
		if err != nil {
            panic(err.Error())
		}
		log.Printf(t.Name)
	}
}