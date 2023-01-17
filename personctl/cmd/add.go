/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type Person struct {
	firstName    string
	lastName     string
	phoneNumbers []string
}

var firstName string
var lastName string
var phone Nargs

type Nargs []string

// String representation of Nargs
func (nargs *Nargs) String() string {
	return strings.Join(*nargs, " ")
}

// Set appends a new argument  to instance of Nargs
func (nargs *Nargs) Set(arg string) error {
	// split by space
	allElems := strings.Split(arg, ",")
	h := *nargs
	for _, elem := range allElems {
		h = append(h, elem)
	}
	*nargs = h
	return nil
}

// Type is a no-op
func (nargs *Nargs) Type() string {
	return "Nargs"
}

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
		flag.Parse()
		phone := strings.Split(phone.String(), " ")
		fmt.Println(phone)

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
	addCmd.Flags().Var(&phone, "phone-number", "phone number of the person")
}
