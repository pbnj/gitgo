package main

import (
	"fmt"
	"os"
	"strings"

	colour "github.com/fatih/color"
	flag "github.com/ogier/pflag"
)

var (
	user  string
	emoji string
)

func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	fmt.Println("")
	for _, u := range users {
		result := getUsers(u)
		colour.Cyan(`Username:	%s`, result.Login)
		colour.Blue(`Name:		%s`, result.Name)
		colour.Green(`Email:		%s`, result.Email)
		colour.HiMagenta(`Bio:		%s`, result.Bio)
		fmt.Println("")
	}

}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
