/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Person struct {
	firstName    string
	lastName     string
	phoneNumbers []string
}

var firstName string
var lastName string
var phone []string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		firstName, err := cmd.Flags().GetString("first-name")
		if err != nil {
			panic(err)
		}
		lastName, err := cmd.Flags().GetString("last-name")
		if err != nil {
			panic(err)
		}
		phone, err := cmd.Flags().GetStringSlice("phone-number")
		if err != nil {
			panic(err)
		}

		p := new(Person)
		p.firstName = firstName
		p.lastName = lastName
		p.phoneNumbers = phone
		fmt.Println(p)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&firstName, "first-name", "", "First name of the person")
	addCmd.Flags().StringVar(&lastName, "last-name", "", "Last name of the person")
	addCmd.Flags().StringSliceVar(&phone, "phone-number", []string{}, "phone number of the person")
}
