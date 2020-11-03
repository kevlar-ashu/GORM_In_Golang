package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User is struct which contains User data
type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:Ashu88742@@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:Ashu88742@@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:Ashu88742@@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open("mysql", "root:Ashu88742@@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initialMigration() {

	db, err := gorm.Open("mysql", "root:Ashu88742@@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main() {

	fmt.Println("Go ORM Tutorial")

	initialMigration()
	// Handle Subsequent requests
	handleRequests()

}

/* // Student is a struct type which contains all fields related with Student data
type Student struct {
	ID int
	//Name string
	Address string `gorm:"type:varchar(30)"`
}

func getConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Ashu88742@@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return db
}

func main() {
	db := getConnection()
	fmt.Println(db)
	db.AutoMigrate(&Student{})

	defer db.Close()
} */
