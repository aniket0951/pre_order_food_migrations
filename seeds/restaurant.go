package seeds

import (
	"pre_order_migration/models"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func Restaurant(db *gorm.DB) error {
	err := db.Create(&models.Restaurant{
		Name:         "The Gourmet Bistro",
		CuisineTypes: []string{"Italian", "French", "Mediterranean"},
		OpenTime:     "11:00",
		CloseTime:    "23:00",
		IsVerified:   true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err.Error != nil {
		return err.Error
	}

	err = db.Create(&models.Restaurant{
		Name:         "The Vegies",
		CuisineTypes: []string{"Italian", "French", "Mediterranean"},
		OpenTime:     "11:00",
		CloseTime:    "23:00",
		IsVerified:   true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err.Error != nil {
		return err.Error
	}

	/* --------------------------------- PAYMENT DETAILS -------------------------------- */
	err = db.Create(&models.PaymentDetails{
		RestaurantID: uuid.New(),
		UpiCode:      "abc@upi",
		IsVerified:   true,
		IsActive:     true,
		QRCodePath:   "/path/to/qr_code.png",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err.Error != nil {
		return err.Error
	}

	/* -------------------------- REGISTRATION DETAILS -------------------------- */

	err = db.Create(&models.RegistrationDetails{
		RestaurantID:    uuid.New(),
		GstnNumber:      "29ABCDE1234F1Z5",
		CstnNumber:      "29ABCDE1234F1Z6",
		EstablishedDate: time.Now(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

	if err.Error != nil {
		return err.Error
	}

	/* -------------------------------- CONTACTS -------------------------------- */

	err = db.Create(&models.Contact{
		RestaurantID: uuid.New(),
		MobileNumber: pq.StringArray{"+1234567890", "+0987654321"},
		EmailId:      "contact@example.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err.Error != nil {
		return err.Error
	}

	/* --------------------------------- ADDRESS -------------------------------- */
	err = db.Create(&models.Address{
		RestaurantID: uuid.New(),
		AddressLine1: "123 Gourmet Street",
		City:         "Foodville",
		State:        "Culinary",
		PinCode:      "12345",
		Latitude:     12.345678,
		Longitude:    98.765432,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	return err.Error
}
