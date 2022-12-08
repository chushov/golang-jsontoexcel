package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type SelfemployedActivities struct {
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

	defer sourceFile.Close()

	var Naming []SelfemployedActivities
	if err := json.NewDecoder(sourceFile).Decode(&Naming); err != nil {
		return err
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

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
	if err := convertJSONToCSV("json/activities_prod.json", "xlsx/data.csv"); err != nil {
		log.Fatal(err)
	}
}
