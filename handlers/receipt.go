package handlers

import (
	"errors"
	"fetch-rewards/receipt-processor-challenge/models"
	"fetch-rewards/receipt-processor-challenge/utils"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
	Repository models.ReceiptRepository
}

// creates a new receipt handler with the given repository
func NewReceiptHandler(r models.ReceiptRepository) *ReceiptHandler {
	return &ReceiptHandler{Repository: r}
}

// processes a receipt and stores it in the repository
func (h *ReceiptHandler) ProcessReceipt(c *gin.Context) {
	var r models.Receipt
	// binds the request body to the receipt model
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
			"message": "The receipt is invalid. Please verify input",
			"details": err.Error(),
		}})
		return
	}

	// creates a new receipt in the repository
	id, err := h.Repository.Create(r)
	if err != nil {
		// handles the case where the receipt already exists
		if errors.Is(err, models.ErrAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// handles other errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// returns the id of the created receipt
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// retrieves points for a receipt by id
func (h *ReceiptHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")

	// retrieves the receipt from the repository by id
	r, err := h.Repository.GetByID(id)
	if err != nil {
		// handles the case where the receipt is not found
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// handles other errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// calculates points based on the retailer name
	points := utils.CountAlphaNumeric(r.Retailer)

	// adds 50 points if the total is a whole number
	if int(r.Total.Float64()*100)%100 == 0 {
		points += 50
	}

	// adds 25 points if the total is a multiple of 0.25
	if int(r.Total.Float64()*100)%25 == 0 {
		points += 25
	}

	// adds 5 points for every two items
	points += len(r.Items) / 2 * 5

	// adds points for items with descriptions whose trimmed length is a multiple of 3
	for _, item := range r.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))

		if trimmedLength%3 == 0 {
			points += int(math.Ceil(item.Price.Float64() * 0.2))
		}
	}

	// adds 6 points if the purchase date is odd
	if r.PurchaseDate.Time().Day()%2 != 0 {
		points += 6
	}

	// checks if the purchase time falls between 2 pm and 4 pm
	purchaseTime := r.PurchaseTime.Time()
	time1 := time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)
	time2 := time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)

	if purchaseTime.After(time1) && purchaseTime.Before(time2) {
		points += 10
	}

	// returns the calculated points
	c.JSON(http.StatusOK, gin.H{"points": strconv.Itoa(points)})
}
