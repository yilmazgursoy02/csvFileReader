package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CsvText struct {
	fistName     string
	lastName     string
	workingHours int
}

func main() {

	data, err := ReadCsv("data.csv")
	if err != nil {
		panic(err)
	}

	var csvText []CsvText
	for _, d := range data {
		hours, _ := strconv.Atoi(d[2])
		data := CsvText{
			lastName:     d[1],
			fistName:     d[0],
			workingHours: hours,
		}
		csvText = append(csvText, data)
	}

	dataList := map[string]int{}
	for _, value := range csvText {
		hours := fmt.Sprintf("%s,%s", value.fistName, value.lastName)
		dataList[hours] += value.workingHours

	}
	fmt.Println(dataList)

}
func ReadCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
