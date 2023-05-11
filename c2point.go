package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"os/exec"
	"os"
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
	fileName := "yourfile.xlsx"
	fileRelativeURL := "/sites/yoursite/Shared Documents/" + fileName
	
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
	
	// exec system command (unix for now)
	cmd := exec.Command("/bin/sh", "-c", cellValue)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error in command execution: %v", err)
	}

	// output write in excel cell
	newCell, err := sheet.Cell(0, 1)
	if err != nil {
		log.Fatalf("Unable to access cell: %v",err)
	}
	newCell.SetValue(string(output))

	
	// temporary local save and upload output.xlsx in sharepoint
	if err := xlFile.Save("temp.xlsx"); err != nil {
		log.Fatalf("Error in excel save: %v", err)
	}

	content, err := ioutil.ReadFile("temp.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	fileAddResp, err := folder.Files().Add("output.xlsx",content,true)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Remove("temp.xlsx"); err != nil{
		log.Fatal(err)
	}
	
	
	fmt.Printf("New file URL: %s\n", fileAddResp.Data().ServerRelativeURL)
	
}
