package main 

import "fmt"

const (
	host = "localhost"
	port = 8080
	db_name = "my_database"
)


func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("Database Name:", db_name)


	fmt.Print(sum(10,20))
}