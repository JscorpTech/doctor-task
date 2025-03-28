package main

import (
	"log"

	"github.com/JscorpTech/jst-go/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "manage",
	}
	rootCmd.AddCommand(commands.RunServer())
	rootCmd.AddCommand(commands.Migrate())
	rootCmd.AddCommand(commands.Doctor())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
