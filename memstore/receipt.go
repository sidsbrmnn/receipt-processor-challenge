package memstore

import (
	"fetch-rewards/receipt-processor-challenge/models"

	"github.com/google/uuid"
)

type inMemoryReceiptRepository struct {
	receipts map[uuid.UUID]*models.Receipt
}

// creates a new in-memory receipt repository
func NewReceiptRepository() (models.ReceiptRepository, error) {
	return &inMemoryReceiptRepository{
		receipts: make(map[uuid.UUID]*models.Receipt),
	}, nil
}

// creates a new receipt in the repository
func (r *inMemoryReceiptRepository) Create(receipt *models.Receipt) (string, error) {
	// checks if the receipt already exists
	// this is only an edge case for the in-memory repository
	if _, ok := r.receipts[receipt.ID]; ok {
		return "", models.ErrAlreadyExists
	}

	// generates a new id for the receipt and stores it in the repository
	receipt.ID = uuid.New()
	r.receipts[receipt.ID] = receipt

	// returns the id of the created receipt
	return receipt.ID.String(), nil
}

// retrieves a receipt from the repository by id
func (r *inMemoryReceiptRepository) GetByID(id string) (*models.Receipt, error) {
	// parses the id string to a UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		// returns an error if the id is invalid
		return nil, models.ErrNotFound
	}

	// retrieves the receipt from the repository by id
	receipt, ok := r.receipts[uid]
	if !ok {
		// returns an error if the receipt does not exist
		return nil, models.ErrNotFound
	}

	// returns the receipt
	return receipt, nil
}
