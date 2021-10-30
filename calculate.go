package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CsvText struct {
	fistName     string
	lastName     string
	workingHours int
}

func main() {

	Calculate("data.csv", "results.csv")

}
func Calculate(inputFile string, outputFile string) {
	data := map[string]int{}
	read, err := ReadCsv(inputFile)
	CheckError("File read error:", err)

	for _, value := range read {
		hours := fmt.Sprintf("%s,%s", value.fistName, value.lastName)
		data[hours] += value.workingHours
	}
	WriterCsv(data, outputFile)
}

func ReadCsv(filename string) ([]CsvText, error) {
	file, err := os.Open(filename)
	CheckError("File open error:", err)

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	CheckError("File read error:", err)

	var csvText []CsvText
	for _, r := range records {
		hour, _ := strconv.Atoi(r[2])
		data := CsvText{
			lastName:     r[1],
			fistName:     r[0],
			workingHours: hour,
		}
		csvText = append(csvText, data)
	}
	return csvText, nil
}

func WriterCsv(dataList map[string]int, results string) {
	file, err := os.OpenFile(results, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	CheckError("failed creating error:", err)

	fileWriter := bufio.NewWriter(file)

	for key, value := range dataList {
		str := fmt.Sprintf("%s,%d \n", key, value)
		_, err = fileWriter.WriteString(str)
		CheckError("Write error:", err)
	}
	fileWriter.Flush()
	defer file.Close()
}

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
