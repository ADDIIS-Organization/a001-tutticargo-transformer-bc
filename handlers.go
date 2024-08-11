package main

import (
	"fmt" // Paquete para formateo de strings
	"log" // Paquete para manejo de logs

	// Paquete para manejo de números grandes
	"net/http" // Paquete para manejo de peticiones HTTP
	"os"       // Paquete para manejo de archivos y sistema operativo
	"strings"  // Paquete para manejo de strings
	"sync"
	"time" // Paquete para manejo de tiempo

	"github.com/360EntSecGroup-Skylar/excelize/v2" // Librería para leer archivos Excel
	"github.com/gin-gonic/gin"                     // Framework web para Go
)

const maxConcurrency = 1 // Aumenta el número de gorutinas concurrentes
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
	insertedStoresCount := 0
	var mu sync.Mutex

	// Iterar sobre las filas del archivo Excel
	for i, row := range rows {
		if i == 0 {
			continue // Omitir la fila de encabezado
		}

		// Incrementar el contador de WaitGroup
		wg.Add(1)
		// Enviar una señal al semáforo, esto para indicar que una gorutina está en ejecución
		semaphore <- struct{}{}

		// Crear una gorutina para procesar la fila actual
		go func(row []string, i int) {
			defer wg.Done() // Decrementar el contador de WaitGroup al finalizar la gorutina
			defer func() { <-semaphore }()

			// Validar que la fila tenga al menos 14 columnas y que las columnas 0 y 5 no estén vacías
			if len(row) < 14 || strings.TrimSpace(row[0]) == "" || strings.TrimSpace(row[5]) == "" {
				fmt.Println("Skipping incomplete or invalid row:", row, "fila", i)
				logger.Printf("Skipping incomplete or invalid row: %v fila %d\n", row, i)
				return
			}

			// Validar que la columna 12 no esté vacía si la columna 13 no está vacía
			if strings.TrimSpace(row[12]) == "" && strings.TrimSpace(row[13]) != "" {
				fmt.Println("Skipping row with empty reference and non-empty provider:", row)
				return
			}

			storeCode := parseInt(row[1])
			store, err := getStoreByCode(storeCode)
			if err != nil || store == nil {
				logger.Printf("Error getting store with code %d: %s", storeCode, err.Error())
				return
			}

			if storeInsertedToday(store.ID) {
				logger.Printf("Esta tienda ya ha sido insertada el día de hoy: %d, fecha: %s", store.ID, time.Now().Format("2006-01-02"))
				return
			}

			orderStore := &OrderStore{StoreID: store.ID}
			err = createOrderStore(orderStore)
			if err != nil {
				logger.Printf("Error creating Order Store: %s", err.Error())
				return
			}

			mu.Lock()             // Bloquear el acceso a la variable compartida
			insertedStoresCount++ // Incrementar el contador de tiendas insertadas
			mu.Unlock()           // Desbloquear el acceso a la variable compartida
		}(row, i)
	}

	wg.Wait() // Esperar a que todas las gorutinas terminen
	logger.Printf("Inserción completada. Total de tiendas insertadas: %d\n", insertedStoresCount)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Las órdenes de %d tiendas fueron insertadas correctamente. Se leyeron %d registros.", insertedStoresCount, len(rows)),
	})
}

func parseInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
