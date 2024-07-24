package main

import (
	"database/sql"
	"errors"
	"log"
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

type OrderPallet struct {
	BigPallets    int64 `db:"big_pallets"`
	LittlePallets int64 `db:"little_pallets"`
	DispoId       int64 `db:"dispo_id"`
	OrderID       int64 `db:"orders_id"`
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

var ErrOrderNotFound = errors.New("order not found")

func getOrderByOrderNumber(orderNumber *big.Int) (*Order, error) {
	var order Order
	err := db.Get(&order, "SELECT id, order_number, detra, date, stores_id FROM orders WHERE order_number = $1", orderNumber.String())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Order not found: %s", orderNumber.String())
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	return &order, nil
}

func createOrder(order *Order) error {
	err := db.QueryRow(`INSERT INTO orders (order_number, detra, date, stores_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		order.OrderNumber, order.Detra, order.Date, order.StoreID).Scan(&order.ID)
	return err
}

func createOrderProduct(orderProduct *OrderProduct) error {
	log.Printf("a punto de insertar orderProduct: %+v", orderProduct)
	_, err := db.Exec(`INSERT INTO order_product (orders_id, products_id, quantity) VALUES ($1, $2, $3)`, orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity)
	return err
}

func createOrderPallet(OrderPallet *OrderPallet) error {
	_, err := db.Exec(`INSERT INTO order_pallet (big_pallets, little_pallets, dispo_id, orders_id) VALUES ($1, $2, $3, $4)`, OrderPallet.BigPallets, OrderPallet.LittlePallets, OrderPallet.DispoId, OrderPallet.OrderID)
	return err
}
