package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type SelfEmployedActivities struct {
	IDCode   int    `json:"id"`
	ParentID int    `json:"parentId"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}

func convertJSONToCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			fmt.Println("Something went wrong", err)
			log.Fatal("Something went wrong", err)
		}
	}(sourceFile)

	var Naming []SelfEmployedActivities
	if err := json.NewDecoder(sourceFile).Decode(&Naming); err != nil {
		return err
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("Cant save output file", err)
			log.Fatal("Cant save output file", err)
		}
	}(outputFile)

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"IDCode", "ParentID", "Name", "Active"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range Naming {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(r.IDCode), fmt.Sprint(r.ParentID), fmt.Sprint(r.Name), fmt.Sprint(r.Active))
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := convertJSONToCSV("json/activities_prod.json", "xlsx/activities_prod.csv"); err != nil {
		log.Fatal(err)
	}
}
