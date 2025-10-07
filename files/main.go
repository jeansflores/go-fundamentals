package main

import (
	"fmt"
	"os"

	"example.com/files/utils"
)

func main() {
	directory, _ := os.Getwd()
	fmt.Println(utils.ReadFile(directory + "/data/text.txt"))
	utils.WriteFile(directory+"/data/new.txt", "I'm a new file!")
}
