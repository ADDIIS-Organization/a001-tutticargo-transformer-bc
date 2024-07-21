package main

import (
	"fmt"      // Paquete para formateo de strings
	"log"      // Paquete para manejo de logs
	"math/big" // Paquete para manejo de números grandes
	"net/http" // Paquete para manejo de peticiones HTTP
	"os"       // Paquete para manejo de archivos y sistema operativo
	"strings"  // Paquete para manejo de strings
	"sync"
	"time" // Paquete para manejo de tiempo

	"github.com/360EntSecGroup-Skylar/excelize/v2" // Librería para leer archivos Excel
	"github.com/gin-gonic/gin"                     // Framework web para Go
)

const maxConcurrency = 80 // Aumenta el número de gorutinas concurrentes
//const batchSize = 100     // Tamaño del lote para inserciones en batch

func test(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Server is running",
	})
}

func uploadFile(c *gin.Context) {
	logFile, err := os.OpenFile("insertion.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open log file"})
		return
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println("Nueva inserción a las", time.Now())

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please upload a valid file."})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file: " + err.Error()})
		return
	}
	defer f.Close()

	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file: " + err.Error()})
		return
	}

	sheetMap := excelFile.GetSheetMap()
	if len(sheetMap) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No sheets found in the Excel file"})
		return
	}

	sheetName := sheetMap[1]
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading rows: " + err.Error()})
		return
	}

	if len(rows) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No rows found in the sheet"})
		return
	}

	fmt.Println("Total rows found:", len(rows))

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)
	insertedProductsCount := 0
	var mu sync.Mutex

	for i, row := range rows {
		if i == 0 {
			continue // Skip header row
		}

		wg.Add(1)
		semaphore <- struct{}{}

		go func(row []string, i int) {
			defer wg.Done()
			defer func() { <-semaphore }()

			if len(row) < 14 || strings.TrimSpace(row[0]) == "" || strings.TrimSpace(row[5]) == "" {
				fmt.Println("Skipping incomplete or invalid row:", row, "fila", i)
				logger.Printf("Skipping incomplete or invalid row: %v fila %d\n", row, i)
				return
			}

			if strings.TrimSpace(row[12]) == "" && strings.TrimSpace(row[13]) != "" {
				fmt.Println("Skipping row with empty reference and non-empty provider:", row)
				return
			}

			productID, storeID, quantity, err := processRowForBatch(row)
			if err != nil {
				logger.Printf("Error processing row: %s", err.Error())
				return
			}

			orderNumber := new(big.Int)
			orderNumber.SetString(row[5], 10)

			order, err := getOrderByOrderNumber(orderNumber)
			if err != nil {
				if err == ErrOrderNotFound {
					order = &Order{
						OrderNumber: orderNumber.String(),
						Detra:       row[6],
						Date:        time.Now().Format(time.RFC3339),
						StoreID:     storeID,
					}
					err = createOrder(order)
					if err != nil {
						logger.Printf("Error creating order: %s", err.Error())
						return
					}
				} else {
					logger.Printf("Error checking if order exists: %s", err.Error())
					return
				}
			}

			var exists bool
			checkQuery := "SELECT EXISTS(SELECT 1 FROM order_product WHERE orders_id = $1 AND products_id = $2 LIMIT 1)"
			err = db.QueryRow(checkQuery, order.ID, productID).Scan(&exists)
			if err != nil {
				logger.Printf("Error checking if order_product exists: %s", err.Error())
				return
			}

			if exists {
				return
			}

			orderProduct := &OrderProduct{
				OrderID:   order.ID,
				ProductID: productID,
				Quantity:  quantity,
			}
			err = createOrderProduct(orderProduct)
			if err != nil {
				logger.Printf("Error creating order_product: %s", err.Error())
				return
			}

			mu.Lock()
			insertedProductsCount++
			mu.Unlock()
		}(row, i)
	}

	wg.Wait()
	logger.Printf("Inserción completada. Total de productos insertados: %d\n", insertedProductsCount)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d products inserted", insertedProductsCount)})
}

func processRowForBatch(row []string) (int64, int64, string, error) {
	ean := new(big.Int)
	ean.SetString(row[0], 10)
	storeCode := parseInt(row[1])
	orderNumber := new(big.Int)
	orderNumber.SetString(row[5], 10)
	detra := new(big.Int)
	detra.SetString(row[6], 10)
	quantity := new(big.Int)
	quantity.SetString(row[8], 10)
	reference := new(big.Int)
	reference.SetString(row[12], 10)

	if reference.String() == "0" {
		return 0, 0, "", nil
	}

	product, err := getProductByEAN(ean)
	if err != nil || product == nil {
		return 0, 0, "", fmt.Errorf("error obteniendo el producto de EAN %s: %w", ean.String(), err)
	}

	store, err := getStoreByCode(storeCode)
	if err != nil || store == nil {
		return 0, 0, "", fmt.Errorf("error obteniendo la tienda de código %d: %w", storeCode, err)
	}

	return product.ID, store.ID, quantity.String(), nil
}

func parseInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
