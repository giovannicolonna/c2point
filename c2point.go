package main

import (
	"context"
	"log"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"os"
	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
	"github.com/koltyakov/gosip/auth/saml"
	"github.com/koltyakov/gosip/data"
	"github.com/koltyakov/gosip/data/sites"
	"github.com/tealeg/xlsx"
)

func main() {

	// Auth config
	authCnfg := &saml.AuthCnfg{
		SiteURL:      "https://yourtenant.sharepoint.com/sites/SiteName", //use your sharepoint tenant URL
		Username:     "youremail@yourtenant.onmicrosoft.com", //use your tenant userid
		Password:     "yourpassword",
		SecurityToken: "",
	}
	// Auth
	ctx := context.Background()
	client := &gosip.SPClient{
		AuthCnfg: authCnfg,
	}

	// HTTP request configuration
	endpoint := client.AuthCnfg.GetSiteURL() + "/_api/web/getfilebyserverrelativeurl('/sites/SiteName/DocumentLibraryName/FileName.xlsx')/$value" //edit this with the path and filename
	api := api.NewHTTPClient(&client.AuthCnfg)
	req, err := api.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalf("Error in http request creation: %v", err)
	}


	// HTTP request execution
	resp, err := api.Do(ctx, req)
	if err != nil {
		log.Fatalf("Error in excel file request: %v", err)
	}
	defer resp.Body.Close()

	// Excel decode
	file, err := xlsx.OpenReader(resp.Body)
	if err != nil {
		log.Fatalf("Error in excel file opening: %v", err)
	}

	// Excel file instruction reading
	sheet := file.Sheets[0]
	for _, row := range sheet.Rows[1:] {
		command := row.Cells[0].Value
		args := strings.Split(row.Cells[1].Value, " ")

		// command execution os-based
		if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd.exe", "/C", command)
			for _, arg := range args {
				cmd.Args = append(cmd.Args, arg)
			}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				log.Fatalf("Error in command execution: %v", err)
			}
		} else {
			cmd := exec.Command("/bin/bash", "-c", command+" "+strings.Join(args, " "))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				log.Fatalf("Error in command execution: %v", err)
			}
		}
	}

	fmt.Println("Command executed")
}

}
