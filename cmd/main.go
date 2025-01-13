package main

import (
	"fetch-rewards/receipt-processor-challenge/handlers"
	"fetch-rewards/receipt-processor-challenge/memstore"
	"fetch-rewards/receipt-processor-challenge/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	// register custom validators for pattern matching and precision checking
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("regex", utils.ValidateRegex)
		v.RegisterValidation("precision", utils.ValidatePrecision)
	}

	// create a new in-memory receipt repository to store and manage receipt data
	receiptRepo, err := memstore.NewReceiptRepository()
	if err != nil {
		// log a fatal error and terminate if repository creation fails
		log.Fatalln(err)
	}

	// create a new receipt handler with the repository
	receiptHandler := handlers.NewReceiptHandler(receiptRepo)

	// register API routes
	r.POST("/receipts/process", receiptHandler.ProcessReceipt)
	r.GET("/receipts/:id/points", receiptHandler.GetPoints)

	// get the HTTP port from environment variables
	httpPort, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		// if the environment variable is not set, use the default port 8080
		httpPort = "8080"
	}

	// run the server on the specified port
	r.Run(":" + httpPort)
}
