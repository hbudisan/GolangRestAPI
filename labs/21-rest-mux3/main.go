package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func main() {
	// Initialize Database
	Connect("root:Welcome1@tcp(127.0.0.1:3306)/sales")

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Println("Starting Server on port 8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8080), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", GetProductById).Methods("GET")
	router.HandleFunc("/api/products", CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", UpdateProduct).Methods("PUT")

	router.HandleFunc("/api/products/{id}", DeleteProduct).Methods("DELETE")
}

func Connect(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	DB.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func checkIfProductExists(productId string) bool {
	var product Product
	DB.First(&product, productId)

	return product.ID != 0
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product Product
	DB.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product Product
	DB.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	DB.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product Product
	DB.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}
