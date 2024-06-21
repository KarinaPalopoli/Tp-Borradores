package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	data := [][]string{
		{"vegetables", "fruits"},
		{"carrot", "banana"},
		{"potato", "strawberry"},
	}

	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// initialize csv writer
	writer := csv.NewWriter(file)

	defer writer.Flush()

	// write all rows at once
	writer.WriteAll(data)

	// write single row
	extraData := []string{"lettuce", "raspberry"}
	writer.Write(extraData)

	fmt.Println("Elija una opcion: ")

	reader := bufio.NewReader(os.Stdin)

	opcion, _ := reader.ReadString('\n')

	fmt.Print(opcion)

}
