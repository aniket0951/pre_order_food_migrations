package models

import (
	"time"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                                  Menu Card                                 */
/* -------------------------------------------------------------------------- */
type MenuCard struct {
	ID           uuid.UUID  `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID uuid.UUID  `gorm:"column:restaurant_id;unique;not null"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
	IsActive     bool       `gorm:"column:is_active;default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

/* -------------------------------------------------------------------------- */
/*                  Category represents the categories table                  */
/* -------------------------------------------------------------------------- */
type Category struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Type        string // e.g., "Appetizer", "Dessert"
	IsDeleted   bool   `gorm:"column:is_deleted;deafult:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

/* -------------------------------------------------------------------------- */
/*                       Item represents the items table                      */
/* -------------------------------------------------------------------------- */
type Item struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null;"`
	Description string
	CategoryID  uuid.UUID `gorm:"not null"`
	MenuCardID  uuid.UUID `gorm:"not null"`
	DietaryType string    // e.g., "Vegetarian", "Non-Vegetarian"
	IsAvailable bool      `gorm:"default:true"`
	ImageURL    string
	Category    Category    `gorm:"foreignKey:CategoryID"`
	MenuCard    MenuCard    `gorm:"foreignKey:MenuCardID"`
	Prices      []ItemPrice `gorm:"foreignKey:ItemID"`
	IsDeleted   bool        `gorm:"default:false;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

/* -------------------------------------------------------------------------- */
/*                                 Item Prices                                */
/* -------------------------------------------------------------------------- */
type ItemPrice struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	ItemID    uuid.UUID `gorm:"not null"`
	Size      string    // e.g., "Small", "Large"
	Price     float64   `gorm:"not null"`
	Item      Item      `gorm:"foreignKey:ItemID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

/* -------------------------------------------------------------------------- */
/*                                  PreOrder                                  */
/* -------------------------------------------------------------------------- */
type PreOrder struct {
	ID          uuid.UUID   `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	ItemIDS     []uuid.UUID `gorm:"column:item_ids;not null"`
	TotalPrice  int64       `gorm:"column:total_price;not null"`
	PaymentID   uuid.UUID   `gorm:"column:payment_id;not null"`
	TimeToReach time.Time   `gorm:"column:time_to_reach;not null"`
	Item        []Item      `gorm:"foreignKey:ItemIDS"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderPayment struct {
	ID            uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	PaymentMethod string    `gorm:"column:payment_method;not null"`
	Status        string    `gorm:"column:status;not null"`
	TransactionID string    `gorm:"column:transaction_id;not null"`
	DebetorID     uuid.UUID `gorm:"column:debetor_id;not null"`
	CreditorID    uuid.UUID `gorm:"column:creditor_id;not null"`
	PreOrderID    uuid.UUID `gorm:"column:pre_order_id;not null"`
	PreOrder      PreOrder  `gorm:"foreignKey:PreOrderID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
