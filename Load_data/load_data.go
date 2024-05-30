package main

import (
	"encoding/csv"
	"os"
	"pincodeapp/models"
)

func loadCSV() {
	file, err := os.Open("/Users/gautam/go/src/github.com/Gautam3767/pincodeapp/data/all-india-pincode-html-csv.csv") // Update the path to your CSV file
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records[1:] { // Skipping the header row
		pincode := models.Pincode{
			Officename:        record[0],
			Pincode:           record[1],
			OfficeType:        record[2],
			Deliverystatus:    record[3],
			Divisionname:      record[4],
			Regionname:        record[5],
			Circlename:        record[6],
			Taluk:             record[7],
			Districtname:      record[8],
			Statename:         record[9],
			Telephone:         record[10],
			RelatedSuboffice:  record[11],
			RelatedHeadoffice: record[12],
		}
		models.DB.Create(&pincode)
	}
}

func main() {
	models.InitDatabase()
	loadCSV()
}
