package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)
var db *sql.DB


type User struct {
	Id int
	Name string
	Age int
	Country string
	City string
}

func connectToDb(){
	config := mysql.Config{
		User: "root",
		Passwd: "Pietranayra@3",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "simple_go_crud",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}

 pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("------ CONNECTED TO DB ------")
}

// func insertToDb(user User){
// 	result, err := db.Exec("INSERT INTO users (id, name, age, country, city) VALUES (?, ?, ?, ?, ?)", user.Id, user.Name, user.Age, user.Country, user.City)
// 	if err != nil{
// 		panic(err)
// 	}else {
// 		fmt.Println(result)
// 		fmt.Println("INSERTED user " + user.Name)
// 	}
// }

func getAllUsersFromDb()[]User{
	var users []User
	result, err := db.Query("SELECT id, name, age, country, city FROM users")
	if err != nil{
		panic(err.Error())
	}else {
		for result.Next(){
			var user User 
			err = result.Scan(&user.Id, &user.Name, &user.Age, &user.Country, &user.City)
			if err != nil{
				panic(err.Error())
			}else{
				fmt.Println("QUERY FROM users ==> ", user.Name)
				users = append(users, user)
			}
		}
	}
	return users
}

func getAllUsers(context *gin.Context){
		context.IndentedJSON(http.StatusOK, getAllUsersFromDb())
}

func getUserFromDbById(id string)User{
	var user User
	result := db.QueryRow("SELECT id, name, age, country, city FROM users WHERE id = ?", id);
	err := result.Scan(&user.Id, &user.Name, &user.Age, &user.Country, &user.City)
	if err != nil{
		panic(err.Error())
	}else{
		return user
	}
}

func getUserById(context *gin.Context){
	id := context.Param("id")
	context.IndentedJSON(http.StatusOK, getUserFromDbById(id))
}



func main(){
	connectToDb()
	router := gin.Default()
	router.GET("/users", getAllUsers)
	router.GET("/users/:id", getUserById)
	router.Run("localhost:8080")
}