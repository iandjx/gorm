// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// type User struct {
// 	gorm.Model
// 	Name  string
// 	Email string
// }

// func allUsers(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	var users []User
// 	db.Find(&users)
// 	fmt.Println("{}", users)

// 	json.NewEncoder(w).Encode(users)
// }

// func newUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("New User Endpoint Hit")

// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]
// 	email := vars["email"]

// 	fmt.Println(name)
// 	fmt.Println(email)

// 	db.Create(&User{Name: name, Email: email})
// 	fmt.Fprintf(w, "New User Successfully Created")
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]

// 	var user User
// 	db.Where("name = ?", name).Find(&user)
// 	db.Delete(&user)

// 	fmt.Fprintf(w, "Successfully Deleted User")
// }

// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]
// 	email := vars["email"]

// 	var user User
// 	db.Where("name = ?", name).Find(&user)

// 	user.Email = email

// 	db.Save(&user)
// 	fmt.Fprintf(w, "Successfully Updated User")
// }

// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/users", allUsers).Methods("GET")
// 	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
// 	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
// 	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8081", myRouter))
// }

// func initialMigration() {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	// Migrate the schema
// 	db.AutoMigrate(&User{})
// }

// func main() {
// 	fmt.Println("Go ORM Tutorial")

// 	initialMigration()
// 	// Handle Subsequent requests
// 	handleRequests()
// }

package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code      string
	Price     uint
	CompanyID string
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	ID   int
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100, Company: Company{ID: 1, Name: "ian"}})

	// Read
	var products []Product
	db.Find(&products) // find product with integer primary key
	fmt.Println(products)
	var product Product
	db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product)
	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
