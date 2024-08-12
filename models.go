/*
*
  - En este archivo se definen las estructuras de datos que se van a utilizar en la aplicación
  - Recordemos tambien que en Go, las estructuras de datos se definen con la palabra reservada type
  - y se pueden definir de la siguiente manera:
  - type NombreDeLaEstructura struct {
  - NombreDelCampo TipoDeDato `etiqueta:"nombre_de_la_columna_en_la_base_de_datos"`
  - }
    *
  - En este caso, se definen las estructuras Product, Store, Order, OrderProduct, OrderStore y OrderPallet
  - También se definen las funciones que se van a utilizar para interactuar con la base de datos
*/
package main

import (
	"fmt"
	"log"
)

type Store struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Code int    `db:"code"`
}

type OrderStore struct {
	StoreID int64  `db:"store_id"`
	Date    string `db:"date"`
}

func getStoreByCode(code int) (*Store, error) {
	var store Store
	err := db.Get(&store, "SELECT id, name, code FROM stores WHERE code = $1", code)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func storeInsertedToday(storeId int64, logger *log.Logger) bool {
	var count int
	query := `
        SELECT COUNT(*) 
        FROM order_store 
        WHERE store_id = $1 
        AND DATE(date) = CURRENT_DATE;`

	// Aquí se está construyendo la consulta final con el valor interpolado.
	fullQuery := fmt.Sprintf("SELECT COUNT(*) FROM order_store WHERE store_id = %d AND DATE(date) = CURRENT_DATE;", storeId)

	fmt.Println("Query:", fullQuery)
	logger.Printf("Query: %s", fullQuery)

	err := db.Get(&count, query, storeId)
	fmt.Println("Desde storeInsertedToday count:", count)
	logger.Printf("Desde storeInsertedToday count: %d", count)
	if err != nil {
		log.Printf("Error getting count of store inserted today: %s", err.Error())
		return false
	}
	return count > 0
}

func createOrderStore(orderStore *OrderStore) error {
	_, err := db.Exec(`
		INSERT INTO order_store (store_id, date) 
		VALUES ($1, CURRENT_TIMESTAMP)`, orderStore.StoreID)
	return err
}
