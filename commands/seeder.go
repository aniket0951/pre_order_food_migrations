package commands

import (
	"log"
	"pre_order_migration/connections"
	"pre_order_migration/seeds"

	"github.com/spf13/cobra"
)

func Seeder() *cobra.Command {
	return &cobra.Command{
		Use: "seed",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := connections.DB()
			if err := seeds.Restaurant(db); err != nil {
				log.Println("Seed Error Restaurant : ", err)
			}

			if err := seeds.Menu(db); err != nil {
				log.Println("Seed Error Menu : ", err)
			}

			return nil
		},
	}
}
