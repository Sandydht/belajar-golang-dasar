package database

import "fmt"

var connection string

func init() { // Function yang secara otomatis akan dieksekusi
	connection = "MySQL"
	fmt.Println("Init dipanggil")
}

func GetDatabase() string {
	return connection
}
