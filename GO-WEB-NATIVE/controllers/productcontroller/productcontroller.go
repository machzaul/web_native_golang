package productcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productsmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productsmodel.GetAll()
	data := map[string]any{
		"products": products,
	}
	tempt, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	tempt.Execute(w, data)
}
func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := productsmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}
	temp, err := template.ParseFiles("views/Product/detail.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var product entities.Product
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}
		Stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = int64(Stock)
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productsmodel.Create(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}

}
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		idstring := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}
		product := productsmodel.Detail(id)

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product":    product,
		}

		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var product entities.Product

		idstring := r.FormValue("id")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}
		Stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = int64(Stock)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		if ok := productsmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		panic(err)
	}

	if err := productsmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
