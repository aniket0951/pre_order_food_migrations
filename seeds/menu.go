package seeds

import (
	"pre_order_migration/models"

	"gorm.io/gorm"
)

/*
INSERT INTO categories (name, description, type) VALUES ('Appetizers', 'Starters to begin your meal', 'Appetizer');
INSERT INTO categories (name, description, type) VALUES ('Main Courses', 'The main part of your meal', 'Main Course');
INSERT INTO categories (name, description, type) VALUES ('Desserts', 'Sweet dishes to end your meal', 'Dessert');
INSERT INTO categories (name, description, type) VALUES ('Beverages', 'Drinks and refreshment options', 'Beverage');
INSERT INTO categories (name, description, type) VALUES ('Daily Specials', 'Special dishes available daily', 'Special');

*/

func Menu(db *gorm.DB) error {
	err := db.Create(&models.Category{
		Name:        "Appetizers",
		Description: "Starters to begin your meal",
		Type:        "Appetizer",
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&models.Category{
		Name:        "Main Courses",
		Description: "The main part of your meal",
		Type:        "Main Course",
	}).Error

	if err != nil {
		return err
	}
	err = db.Create(&models.Category{
		Name:        "Desserts",
		Description: "Sweet dishes to end your meal",
		Type:        "Dessert",
	}).Error

	if err != nil {
		return err
	}
	err = db.Create(&models.Category{
		Name:        "Beverages",
		Description: "Drinks and refreshment options",
		Type:        "Beverage",
	}).Error

	if err != nil {
		return err
	}
	err = db.Create(&models.Category{
		Name:        "Daily Specials",
		Description: "Special dishes available daily",
		Type:        "Special",
	}).Error

	return err
}
