package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1.homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//2.categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//3.products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
