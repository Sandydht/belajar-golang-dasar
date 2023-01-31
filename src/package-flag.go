package src

import (
	"flag"
	"fmt"
)

/**
Package flag berisikan fungsionalitas untuk memparsing command line argument
*/

func PackageFlag() {
	host := flag.String("host", "localhost", "Put yout database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
