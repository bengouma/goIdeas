package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
)

func main() {
	wordFile, errOne := os.Open("dictionary.txt")
	operatorFile, errTwo := os.Open("operations.txt")

	if errOne != nil {
		fmt.Println("Error!")
	}

	if errTwo != nil {
		fmt.Println("Error Two!!!")
	}

	scannerOne := bufio.NewScanner(wordFile)
	scannerTwo := bufio.NewScanner(operatorFile)
	var wordArrayOne[58103] string
	var wordArrayTwo[10] string
	indexOne := 0
	indexTwo := 0

	for scannerOne.Scan() {
		wordArrayOne[indexOne] = scannerOne.Text()
		indexOne++
	}

	for scannerTwo.Scan() {
		wordArrayTwo[indexTwo] = scannerTwo.Text()
		indexTwo++
	}

	fmt.Println(wordArrayOne[rand.Intn(58103)], wordArrayTwo[rand.Intn(10)])
}





