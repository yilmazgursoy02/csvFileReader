package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type CsvText struct {
	firstName    string
	lastName     string
	workingHours int
}

const (
	SortOrderLastName = "lastname"
	SortOrderHours    = "hours"
)

func main() {

	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("You must provide arguments. Run the command as ./calculate <input_file> <output_file> <order>")
	}
	err := Calculate(args[0], args[1], args[2])
	CheckError("failed to calculate:", err)

}
func Calculate(inputFile string, resultFileName string, order string) error {
	data := make(map[CsvText]int)
	read, err := ReadCsv(inputFile)
	CheckError("File read error:", err)

	for _, value := range read {
		h := CsvText{lastName: value.lastName, firstName: value.firstName}
		data[h] += value.workingHours
	}
	var dataList []CsvText
	for e, hours := range data {
		d := CsvText{lastName: e.lastName, firstName: e.firstName, workingHours: hours}
		dataList = append(dataList, d)
	}

	sortData, err := SortData(dataList, order)
	CheckError("Sort data error:", err)

	if err := WriterCsv(sortData, resultFileName); err != nil {
		return err
	}

	return nil
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
			firstName:    r[0],
			workingHours: hour,
		}
		csvText = append(csvText, data)
	}
	return csvText, nil
}

func WriterCsv(dataList []CsvText, results string) error {
	file, err := os.OpenFile(results, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	CheckError("failed creating error:", err)

	fileWriter := bufio.NewWriter(file)

	for _, v := range dataList {
		str := fmt.Sprintf("%s,%s,%d\n", v.lastName, v.firstName, v.workingHours)
		_, err = fileWriter.WriteString(str)
		CheckError("Write error:", err)
	}
	fileWriter.Flush()
	defer file.Close()
	return nil
}

func SortData(data []CsvText, order string) ([]CsvText, error) {
	switch order {
	case SortOrderLastName:
		sort.Slice(data, func(i, j int) bool {
			return data[i].lastName < data[j].lastName
		})
	case SortOrderHours:
		sort.Slice(data, func(i, j int) bool {
			return data[i].workingHours > data[j].workingHours
		})
	default:
		return nil, fmt.Errorf("unknown sort order type:%s", order)
	}
	return data, nil
}
func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
