package main

import (
	"math/big"
)

type Product struct {
	ID   int64  `db:"id"`
	Code int    `db:"code"`
	EAN  string `db:"ean"`
}

type Store struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Code int    `db:"code"`
}

type Order struct {
	ID          int64  `db:"id"`
	OrderNumber string `db:"order_number"`
	Detra       string `db:"detra"`
	Date        string `db:"date"`
	StoreID     int64  `db:"stores_id"`
}

type OrderProduct struct {
	OrderID   int64  `db:"orders_id"`
	ProductID int64  `db:"products_id"`
	Quantity  string `db:"quantity"`
}

func getProductByEAN(ean *big.Int) (*Product, error) {
	var product Product
	err := db.Get(&product, "SELECT id, code, ean FROM products WHERE ean = $1", ean.String())
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func getStoreByCode(code int) (*Store, error) {
	var store Store
	err := db.Get(&store, "SELECT id, name, code FROM stores WHERE code = $1", code)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func getOrderByOrderNumber(orderNumber *big.Int) (*Order, error) {
	var order Order
	err := db.Get(&order, "SELECT id, order_number, detra, date, stores_id FROM orders WHERE order_number = $1", orderNumber.String())
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func createOrder(order *Order) error {
	_, err := db.NamedExec(`INSERT INTO orders (order_number, detra, date, stores_id) VALUES (:order_number, :detra, :date, :stores_id)`, order)
	return err
}

func createOrderProduct(orderProduct *OrderProduct) error {
	_, err := db.NamedExec(`INSERT INTO order_product (orders_id, products_id, quantity) VALUES (:orders_id, :products_id, :quantity)`, orderProduct)
	return err
}
