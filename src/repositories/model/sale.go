package model

import "time"

type Sale struct {
	ID                  string
	VehicleID           string
	BuyerDocumentNumber string
	Status              string
	SoldAt              time.Time
}
