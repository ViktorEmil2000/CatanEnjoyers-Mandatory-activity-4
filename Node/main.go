package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter port:")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')
	
}
