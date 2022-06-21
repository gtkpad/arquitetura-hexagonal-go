package db

import (
	"database/sql"

	"github.com/gtkpad/arquitetura-hexagonal-go/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.IProduct, error) {
 var product application.Product
 stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
 if err != nil {
	 return nil, err
 }
 err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

 if err != nil {
	 return nil, err
 }

 return &product, nil
}

func (p *ProductDb) Save(product application.IProduct) (application.IProduct, error) {
	var rows int
	p.db.QueryRow("select count(*) from products where id = ?", product.GetID()).Scan(&rows)
	if rows == 0 {
		return p.create(product)
	}
	return p.update(product)
}

func (p *ProductDb) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare("insert into products (id, name, price, status) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(), 
		product.GetName(), 
		product.GetPrice(), 
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.IProduct) (application.IProduct, error) {
	_, err := p.db.Exec("update products set name = ?, price = ?, status = ? where id = ?", 
		product.GetName(), 
		product.GetPrice(), 
		product.GetStatus(),
		product.GetID(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}