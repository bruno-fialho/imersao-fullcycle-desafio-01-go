package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Data struct {
	Name  string
	Age   int
	Score int
}

type DataSet []Data

func sortByNames(dataSet DataSet) DataSet {
	sort.Slice(dataSet, func(i, j int) bool {
		return dataSet[i].Name < dataSet[j].Name
	})
	return dataSet
}

func sortByAge(dataSet DataSet) DataSet {
	sort.Slice(dataSet, func(i, j int) bool {
		return dataSet[i].Age < dataSet[j].Age
	})
	return dataSet
}

func readFile(filePath string) (DataSet, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var dataSet DataSet
	for i, line := range lines {
		if i == 0 {
			continue
		}

		age, _ := strconv.Atoi(line[1])
		score, _ := strconv.Atoi(line[2])

		data := Data{
			Name:  line[0],
			Age:   age,
			Score: score,
		}
		dataSet = append(dataSet, data)
	}

	return dataSet, nil
}

func writeFile(filePath string, dataSet DataSet) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	header := []string{"Nome", "Idade", "Pontuação"}
	if err := csvWriter.Write(header); err != nil {
		return err
	}

	for _, data := range dataSet {
		line := []string{data.Name, strconv.Itoa(data.Age), strconv.Itoa(data.Score)}
		if err := csvWriter.Write(line); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <arquivo-origem.csv> <arquivo-destino.csv>")
		os.Exit(1)
	}

	inputFilePath := os.Args[1]
	outputFilePathByName := "ordenado_por_nome.csv"
	outputFilePathByAge := "ordenado_por_idade.csv"

	dataSet, err := readFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	dataSetSortedByName := sortByNames(dataSet)
	err = writeFile(outputFilePathByName, dataSetSortedByName)
	if err != nil {
		fmt.Println("Error writing the file sorted by name:", err)
		os.Exit(1)
	}

	dataSetSortedByAge := sortByAge(dataSet)
	err = writeFile(outputFilePathByAge, dataSetSortedByAge)
	if err != nil {
		fmt.Println("Error writing the file sorted by age:", err)
		os.Exit(1)
	}

	fmt.Println("Processing completed successfully.")
}
