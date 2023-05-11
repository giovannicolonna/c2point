package main

import (
	"fmt"
	"log"
	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
	"github.com/tealeg/xlsx/v3"
	strategy "github.com/koltyakov/gosip/auth/azurecert"
)

func main() {
	authCnfg := &strategy.AuthCnfg{}
	configPath := "./config/private.json"
	if err := authCnfg.ReadConfig(configPath); err != nil {
		log.Fatalf("Unable to get config: %v", err)
	}
	
	
	client := &gosip.SPClient{AuthCnfg: authCnfg}
	sp := api.NewSP(client)
	
	res, err := sp.Web().Select("Title").Get()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Site title: %s\n", res.Data().Title)	
	
	// +++++ SET HERE YOUR FILENAME
	fileName := "gctesting.xlsx"
	fileRelativeURL := "/sites/gctesting/Shared Documents/" + fileName
	
	file, err := sp.Web().GetFile(fileRelativeURL).Download()
	if err != nil {
		log.Fatalf("Unable to get Excel file: %v", err)
	}
	
	// Open Excel file
	xlFile, err := xlsx.OpenBinary(file)
	if err != nil {
		log.Fatalf("Unable to open Excel file: %v", err)
	}
	sheet := xlFile.Sheets[0]

	// Cell A1 value
	cell, err := sheet.Cell(0, 0)
	if err != nil {
		log.Fatalf("Unable to read cell: %v", err)
	}
	cellValue := cell.Value
	fmt.Println(cellValue)
}
