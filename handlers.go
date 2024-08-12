package main

import (
	"fmt" // Paquete para formateo de strings
	"log" // Paquete para manejo de logs

	// Paquete para manejo de números grandes
	"net/http" // Paquete para manejo de peticiones HTTP
	"os"       // Paquete para manejo de archivos y sistema operativo
	"strings"  // Paquete para manejo de strings
	"time"     // Paquete para manejo de tiempo

	"github.com/360EntSecGroup-Skylar/excelize/v2" // Librería para leer archivos Excel
	"github.com/gin-gonic/gin"                     // Framework web para Go
)

func test(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": fmt.Sprintf("Server is running and today's date is: %s", time.Now().Format("2006-01-02")),
	})
}

func uploadFile(c *gin.Context) {
	// 1. Abrir el archivo de log en modo de escritura, creación y apertura
	logFile, err := os.OpenFile("insertion.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// 2. Manejar el error si no se pudo abrir el archivo
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open log file"})
		return
	}
	// 3. Cerrar el archivo al finalizar la función
	defer logFile.Close()

	// 4. Crear un nuevo logger con el archivo de log
	logger := log.New(logFile, "", log.LstdFlags)
	// 5. Escribir un mensaje en el archivo de log para diferenciar cada inserción
	logger.Println("Nueva inserción a las", time.Now())

	// 6. Obtener el archivo subido. Aqui, el formato de archivo es un puntero de multipart.FileHeader
	file, err := c.FormFile("file")
	// 7. Manejar el error si no se pudo obtener el archivo
	if err != nil {
		// 7.1. Responder con un error 400 Bad Request si no se pudo obtener el archivo
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please upload a valid file."})
		// 7.2. Escribir un mensaje en el archivo de log si no se pudo obtener el archivo
		logger.Println("Error al obtener el archivo:", err)
		// 7.3. return para cortar la ejecución de la función
		return
	}

	// 8. Abrir el archivo subido, en est epunto obtenemos como tal el archivo subido no el puntero de multipart.FileHeader
	f, err := file.Open()
	// 9. Manejar el error si no se pudo abrir el archivo
	if err != nil {
		// 9.1. Responder con un error 500 Internal Server Error si no se pudo abrir el archivo
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file: " + err.Error()})
		// 9.2. Escribir un mensaje en el archivo de log si no se pudo abrir el archivo
		logger.Println("Error al abrir el archivo:", err)
		// 9.3. return para cortar la ejecución de la función
		return
	}
	// 10. Cerrar el archivo al finalizar la función
	defer f.Close()

	// 11. Crear un nuevo lector de archivos Excel, a este punto excelFile es un punto de acceso a las hojas y celdas del archivo Excel
	excelFile, err := excelize.OpenReader(f)
	// 12. Manejar el error si no se pudo abrir el archivo Excel
	if err != nil {
		// 12.1. Responder con un error 500 Internal Server Error si no se pudo abrir el archivo Excel
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file: " + err.Error()})
		// 12.2. Escribir un mensaje en el archivo de log si no se pudo abrir el archivo Excel
		logger.Println("Error al leer el archivo:", err)
		// 12.3. return para cortar la ejecución de la función
		return
	}

	// 13. Obtener un mapa de las hojas del archivo Excel
	sheetMap := excelFile.GetSheetMap()
	// 14. Validar que el archivo Excel tenga al menos una hoja
	if len(sheetMap) == 0 {
		// 14.1. Responder con un error 500 Internal Server Error si no se encontraron hojas en el archivo Excel
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No sheets found in the Excel file"})
		// 14.2. Escribir un mensaje en el archivo de log si no se encontraron hojas en el archivo Excel
		logger.Println("No se encontraron hojas en el archivo Excel")
		// 14.3. return para cortar la ejecución de la función
		return
	}

	// 15. Obtener el nombre de la primera hoja del archivo Excel
	sheetName := sheetMap[1]
	// 16. Obtener las filas de la hoja del archivo Excel
	rows, err := excelFile.GetRows(sheetName)
	// 17. Manejar el error si no se pudieron obtener las filas de la hoja
	if err != nil {
		// 17.1. Responder con un error 500 Internal Server Error si no se pudieron obtener las filas de la hoja
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading rows: " + err.Error()})
		// 17.2. Escribir un mensaje en el archivo de log si no se pudieron obtener las filas de la hoja
		logger.Println("Error al leer las filas:", err)
		// 17.3. return para cortar la ejecución de la función
		return
	}

	// 18. Validar que se encontraron filas en la hoja del archivo Excel
	if len(rows) == 0 {
		// 18.1. Responder con un error 500 Internal Server Error si no se encontraron filas en la hoja
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No rows found in the sheet"})
		// 18.2. Escribir un mensaje en el archivo de log si no se encontraron filas en la hoja
		logger.Println("No se encontraron filas en la hoja")
		// 18.3. return para cortar la ejecución de la función
		return
	}

	// 19. Imprimir el número total de filas encontradas en el archivo Excel
	fmt.Println("Total rows found:", len(rows))

	// 20. Contador para almacenar el número de tiendas insertadas para posterior uso en respuesta a cliente
	insertedStoresCount := 0

	// 21. Iterar sobre las filas del archivo Excel
	for i, row := range rows {
		// 22. Validar que la fila no sea la primera (encabezados)
		if i == 0 {
			continue
		}

		// 23. Validar que la fila tenga al menos 14 columnas y que las columnas 0 y 5 no estén vacías
		if len(row) < 14 || strings.TrimSpace(row[0]) == "" || strings.TrimSpace(row[5]) == "" {
			// 23.1. Imprimir un mensaje de advertencia en la consola y en el archivo de log
			fmt.Println("Skipping incomplete or invalid row:", row, "fila", i)
			// 23.2. Escribir un mensaje en el archivo de log
			logger.Printf("Skipping incomplete or invalid row: %v fila %d\n", row, i)
			continue
		}

		// 24. Validar que la columna 12 esté vacía y que la columna 13 no esté vacía
		if strings.TrimSpace(row[12]) == "" && strings.TrimSpace(row[13]) != "" {
			// 24.1. Imprimir un mensaje de advertencia en la consola y en el archivo de log
			fmt.Println("Skipping row with empty reference and non-empty provider:", row)
			// 24.2. Escribir un mensaje en el archivo de log
			logger.Printf("Skipping row with empty reference and non-empty provider: %v\n", row)
			// 24.3. return para cortar la ejecución de la función
			continue
		}

		// 25. Obtener el código de la tienda de la columna 1 de la fila
		storeCode := parseInt(row[1])
		// 26. Obtener la tienda con el código obtenido
		store, err := getStoreByCode(storeCode)
		// 27. Manejar el error si no se pudo obtener la tienda
		if err != nil || store == nil {
			// 27.1. Imprimir un mensaje de advertencia en la consola y en el archivo de log
			fmt.Println("Skipping row with invalid store code:", row)
			// 27.2. Escribir un mensaje en el archivo de log
			logger.Printf("Skipping row with invalid store code: %v\n", row)
			// 27.3. return para cortar la ejecución de la función
			continue
		}

		// 28. Validar que la tienda no haya sido insertada el día de hoy
		if storeInsertedToday(store.ID, logger) {
			// 28.1. Imprimir un mensaje de advertencia en la consola y en el archivo de log
			fmt.Println("Skipping store already inserted today:", store.ID)
			// 28.2. Escribir un mensaje en el archivo de log
			logger.Printf("Skipping store already inserted today: %d\n", store.ID)
			// 28.3. return para cortar la ejecución de la función
			continue
		}

		// 29. Crear una nueva orden con los datos de la fila
		orderStore := &OrderStore{StoreID: store.ID}
		// 30. Asignar los valores de la fila a la orden
		err = createOrderStore(orderStore)
		// 31. Manejar el error si no se pudo crear la orden
		if err != nil {
			// 31.1. Imprimir un mensaje de advertencia en la consola y en el archivo de log
			fmt.Println("Error creating Order Store:", err)
			// 31.2. Escribir un mensaje en el archivo de log
			logger.Printf("Error creating Order Store: %s", err.Error())
			// 31.3. return para cortar la ejecución de la función
			continue
		}

		// 32. Incrementar el contador de tiendas insertadas
		insertedStoresCount++ // Incrementar el contador de tiendas insertadas
	}

	// 33. Imprimir un mensaje en la consola y en el archivo de log con el total de tiendas insertadas
	logger.Printf("Inserción completada. Total de tiendas insertadas: %d\n", insertedStoresCount)

	// 34. Responder con un mensaje de éxito al cliente
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Las órdenes de %d tiendas fueron insertadas correctamente. Se leyeron %d registros.", insertedStoresCount, len(rows)),
	})
}

func parseInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
