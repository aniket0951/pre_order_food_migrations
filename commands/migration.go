package commands

import (
	"log"
	"pre_order_migration/connections"
	"pre_order_migration/models"

	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := connections.DB()

			if err := db.AutoMigrate(&models.Restaurant{}); err != nil {
				log.Println("Migration Error models.Restaurant : ", err)
				return err
			}

			if err := db.AutoMigrate(&models.PaymentDetails{}); err != nil {
				log.Println("Migration Error models.PaymentDetails : ", err)
				return err
			}

			if err := db.AutoMigrate(&models.RegistrationDetails{}); err != nil {
				log.Println("Migration Error models.RegistrationDetails : ", err)
				return err
			}
			if err := db.AutoMigrate(&models.Contact{}); err != nil {
				log.Println("Migration Error models.Contact : ", err)
				return err
			}
			if err := db.AutoMigrate(&models.Address{}); err != nil {
				log.Println("Migration Error models.Address : ", err)
				return err
			}

			if err := db.AutoMigrate(&models.Category{}); err != nil {
				log.Println("Migration Error models.Category : ", err)
				return err
			}

			if err := db.AutoMigrate(&models.Item{}); err != nil {
				log.Println("Migration Error models.Item : ", err)
				return err
			}

			if err := db.AutoMigrate(&models.ItemPrice{}); err != nil {
				log.Println("Migration Error models.ItemPrice : ", err)
				return err
			}

			if err := db.AutoMigrate(&models.MenuCard{}); err != nil {
				log.Println("Migration Error models.MenuCard : ", err)
				return err
			}

			// if err := db.AutoMigrate(&models.OrderPayment{}); err != nil {
			// 	log.Println("Migration Error models.OrderPayment : ", err)
			// 	return err
			// }

			if err := db.AutoMigrate(&models.PreOrders{}); err != nil {
				log.Println("Migration Error models.PreOrders : ", err)
				return err
			}

			return nil
		},
	}
}
