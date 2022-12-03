package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Struct Product
type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var (
	database = make(map[string]Product)
)

func main() {
	initDB()

	http.HandleFunc("/", home)
	http.HandleFunc("/get-products", getAllProducts)
	http.HandleFunc("/get-product", getProductByID)
	http.HandleFunc("/add-product", addProduct)
	http.HandleFunc("/update-product", updateProduct)
	http.HandleFunc("/delete-product", deleteProduct)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initDB() {
	database["001"] = Product{ID: "001", Name: "Kopi Excelso", Quantity: 10}
	database["002"] = Product{ID: "002", Name: "Kopi Aroma", Quantity: 5}
}

func SetJSONResp(res http.ResponseWriter, message []byte, httpCode int) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func home(res http.ResponseWriter, req *http.Request) {
	message := []byte(`{"message": "server is up"}`)
	SetJSONResp(res, message, http.StatusOK)
}

func getAllProducts(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		message := []byte(`{"message": "Invalid http method"}`)
		SetJSONResp(res, message, http.StatusMethodNotAllowed)
		return
	}
	//Untuk mengkonversi data map database ke list
	var products []Product

	for _, product := range database {
		products = append(products, product)
	}

	productJson, err := json.Marshal(products)
	if err != nil {
		message := []byte(`{"message": "Error when parsing data"}`)
		SetJSONResp(res, message, http.StatusInternalServerError)
		return
	}
	SetJSONResp(res, productJson, http.StatusOK)
}

func getProductByID(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		message := []byte(`{"message": "Invalid http method"}`)
		SetJSONResp(res, message, http.StatusMethodNotAllowed)
		return
	}

	//Validasi untuk manggil route harus pakai id
	if _, ok := req.URL.Query()["id"]; !ok {
		message := []byte(`{"message": "Required product id"}`)
		SetJSONResp(res, message, http.StatusBadRequest)
		return
	}
	id := req.URL.Query()["id"][0]

	//Validasi data product jika tidak sesuai id nya
	product, ok := database[id]
	if !ok {
		message := []byte(`{"message": "product not found"}`)
		SetJSONResp(res, message, http.StatusOK)
		return
	}

	productJSON, err := json.Marshal(&product)
	if err != nil {
		message := []byte(`{"message": "some error when parsing data"}`)
		SetJSONResp(res, message, http.StatusInternalServerError)
		return
	}

	SetJSONResp(res, productJSON, http.StatusOK)
}

func addProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		message := []byte(`{"message": "Invalid http method"}`)
		SetJSONResp(res, message, http.StatusMethodNotAllowed)
		return
	}
	//Menampung body dari depan
	var product Product

	payload := req.Body

	defer req.Body.Close()

	err := json.NewDecoder(payload).Decode(&product)
	if err != nil {
		message := []byte(`{"message": "Error Parsing Data"}`)
		SetJSONResp(res, message, http.StatusInternalServerError)
		return
	}
	database[product.ID] = product
	message := []byte(`{"message": "Success Create Product"}`)
	SetJSONResp(res, message, http.StatusCreated)
}

func updateProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		message := []byte(`{"message": "Invalid http method"}`)
		SetJSONResp(res, message, http.StatusMethodNotAllowed)
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		message := []byte(`{"message": "Required product id"}`)
		SetJSONResp(res, message, http.StatusBadRequest)
		return
	}

	id := req.URL.Query()["id"][0]
	product, ok := database[id]
	if !ok {
		message := []byte(`{"message": "product not found"}`)
		SetJSONResp(res, message, http.StatusOK)
		return
	}

	var newProduct Product

	payload := req.Body

	defer req.Body.Close()

	err := json.NewDecoder(payload).Decode(&newProduct)
	if err != nil {
		message := []byte(`{"message": "error when parsing data"}`)
		SetJSONResp(res, message, http.StatusInternalServerError)
		return
	}

	product.Name = newProduct.Name
	product.Quantity = newProduct.Quantity

	database[product.ID] = product

	productJSON, err := json.Marshal(&product)
	if err != nil {
		message := []byte(`{"message": "some error when parsing data"}`)
		SetJSONResp(res, message, http.StatusInternalServerError)
		return
	}

	SetJSONResp(res, productJSON, http.StatusOK)
}

func deleteProduct(res http.ResponseWriter, req *http.Request) {

	if req.Method != "DELETE" {
		message := []byte(`{"message": "Invalid http method"}`)
		SetJSONResp(res, message, http.StatusMethodNotAllowed)
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		message := []byte(`{"message": "Required product id"}`)
		SetJSONResp(res, message, http.StatusBadRequest)
		return
	}

	id := req.URL.Query()["id"][0]
	product, ok := database[id]
	if !ok {
		message := []byte(`{"message": "product not found"}`)
		SetJSONResp(res, message, http.StatusOK)
		return
	}

	delete(database, id)

	productJSON, err := json.Marshal(&product)
	if err != nil {
		message := []byte(`{"message": "some error when parsing data"}`)
		SetJSONResp(res, message, http.StatusInternalServerError)
		return
	}

	SetJSONResp(res, productJSON, http.StatusOK)

}
