package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
)

const maxConcurrency = 99 // Aumenta el número de gorutinas concurrentes
const batchSize = 1000    // Tamaño del lote para inserciones en transacciones

func test(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Server is running",
	})
}

func uploadFile(c *gin.Context) {
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
	sem := make(chan struct{}, maxConcurrency)
	rowChan := make(chan []string, len(rows)-1) // Canal para enviar las filas a las gorutinas

	insertedProductsCount := 0
	mu := &sync.Mutex{}

	// Inicia las gorutinas que procesarán las filas
	for i := 0; i < maxConcurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rowChan {
				processRow(row, &insertedProductsCount, mu, sem)
			}
		}()
	}

	// Envía las filas a las gorutinas
	for i, row := range rows {
		if i == 0 {
			continue // Skip header row
		}

		if len(row) < 14 {
			fmt.Println("Skipping incomplete row:", row)
			continue
		}

		rowChan <- row
	}
	close(rowChan) // Cierra el canal para que las gorutinas terminen

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d products inserted", insertedProductsCount)})
}

func processRow(row []string, insertedProductsCount *int, mu *sync.Mutex, sem chan struct{}) {
	sem <- struct{}{}
	defer func() { <-sem }()

	ean := new(big.Int)
	ean.SetString(row[0], 10)
	storeCode := parseInt(row[1])
	orderNumber := new(big.Int)
	orderNumber.SetString(row[5], 10)
	detra := new(big.Int)
	detra.SetString(row[6], 10)
	quantity := new(big.Int)
	quantity.SetString(row[8], 10)

	product, err := getProductByEAN(ean)
	if err != nil || product == nil {
		log.Printf("Error getting product for EAN %s: %s", ean.String(), err.Error())
		return
	}

	store, err := getStoreByCode(storeCode)
	if err != nil || store == nil {
		log.Printf("Error getting store for code %d: %s", storeCode, err.Error())
		return
	}

	order, err := getOrderByOrderNumber(orderNumber)
	if err != nil {
		log.Printf("Error getting order for order number %s: %s", orderNumber.String(), err.Error())
	}

	if order == nil {
		log.Printf("Inserting Order for OrderNumber %s and Detra %s", orderNumber.String(), detra.String())
		order = &Order{
			OrderNumber: orderNumber.String(),
			Detra:       detra.String(),
			Date:        time.Now().Format(time.RFC3339),
			StoreID:     store.ID,
		}
		err = createOrder(order)
		if err != nil {
			log.Printf("Error creating order: %s", err.Error())
			return
		}
	}

	if order != nil && product != nil {
		log.Printf("Inserting OrderProduct for Order %d and Product %d", order.ID, product.ID)
		orderProduct := &OrderProduct{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  quantity.String(),
		}
		err = createOrderProduct(orderProduct)
		if err != nil {
			log.Printf("Error creating order product: %s", err.Error())
			return
		}

		mu.Lock()
		*insertedProductsCount++
		mu.Unlock()
	}
}

func parseInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
