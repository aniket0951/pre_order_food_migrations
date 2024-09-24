package main

import (
	"log"
	"pre_order_migration/commands"
	"pre_order_migration/utils"

	"github.com/spf13/cobra"
)

func main() {
	utils.LoadConfig()

	cmd := &cobra.Command{}
	cmd.AddCommand(commands.Migrate())
	cmd.AddCommand(commands.Seeder())

	if err := cmd.Execute(); err != nil {
		log.Println("failed commands : ", err)
	}
}
