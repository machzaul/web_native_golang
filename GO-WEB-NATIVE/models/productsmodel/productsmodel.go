package productsmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(
		"select products.id,products.name,categories.name as category_name,products.stock,products.description,products.created_at,products.updated_at from products join categories on products.category_id = categories.id")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products

}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec("insert into products(name,category_id,stock,description,created_at,updated_at)values(?,?,?,?,?,?)",
		product.Name, product.Category.Id, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		panic(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow("select products.id,products.name,categories.name as category_name,products.stock,products.description,products.created_at,products.updated_at from products join categories on products.category_id = categories.id where products.id = ?", id)

	var product entities.Product
	err := row.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	return product

}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec("update products set name = ?,category_id = ?,stock = ?,description = ?, updated_at = ? where id = ?", product.Name, product.Category.Id, product.Stock, product.Description, product.UpdatedAt, id)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0

}

func Delete(id int) error {
	_, err := config.DB.Exec("delete from products where id =?", id)
	return err
}
