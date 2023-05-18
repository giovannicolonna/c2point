# c2point
c2point - experimental and basic c2 structure in Go that uses Attacker-owned Azure Tenant (Sharepoint) as Command and Control.
This is a very simple project that aims to make the "agent" execute a (one for now :) ) command on target server from instructions written into an Online Excel file.

# Prerequisites
Global Admin on an Azure Tenant (if you have a lab, give it a try, set up your Developer Account in Microsoft to try O365 for free)


# Install Go: 
If you haven't already, download and install Go from the official website: https://golang.org/dl/. Follow the installation instructions for your operating system.

# Install dependencies: 
The program depends on several third-party packages, which can be installed using the following command:


```

go get github.com/koltyakov/gosip
go get github.com/koltyakov/gosip/api
go get github.com/koltyakov/gosip/auth/azurecert
go get github.com/tealeg/xlsx/v3

```



# Configure authentication: 

The program works with the support of a valid Azure Tenant with Sharepoint Online (Business subscription). According to Microsoft Documentation, set up your Azure environment to accept App-only access to an Azure app. You will need to register a new AAD Application and grant it to access SharePoint Online. 

Create a new Azure app and set up grants and certificate access as follows:

https://learn.microsoft.com/en-us/sharepoint/dev/solution-guidance/security-apponly-azuread

Replace the placeholder values in the private.json files

****Note that Cert Auth is the only Auth actually working with SharePoint online

Here is an example of how to initialize private.json with the necessary fields for SharePoint Online authentication using Cert Auth as specified in GoSip Documentation (in this link you'll find script to generate self-signed cert): https://go.spflow.com/auth/strategies/azure-certificate-auth

```

{
	"siteUrl": "https://your-site.sharepoint.com/sites/your-site-name",
	"tenantId": "your-tenantid",
	"clientId": "your-clientid",
	"certPath": "cert.pfx",
	"certPass": "password"
}

```
Make sure to replace the placeholders (tenant-id, clientId, certPath and certPass) with the actual values for your SharePoint environment.
**Create a 'config' folder in same directory of your .go file and put inside it the private.json and the .pfx cert file**


# Update file path
  Update the file path in the endpoint variable to point to your SharePoint file.

# Compile and run the program:
  Open a terminal or command prompt and navigate to the directory containing the c2point.go file. Run the following commands to compile and execute the program:

```

go build c2point.go
./c2point

```

If you're on Windows, the compiled executable will be named c2point.exe. In that case, run the following command instead:

```
GOOS=windows GOARCH=amd64 go build -o ./bin/c2point.exe c2point.go 

c2point.exe
```

The program will read the instructions from the Excel file and execute them as command-line commands on your machine. On Windows, the commands will be executed using the cmd.exe shell, while on Unix-based systems they will be executed using the /bin/bash shell.

Note that - for now - cert.pfx file and private.json must be in same folder.

# Further improvements
Many, this is a sort of experiment:
1 - introduce infinite loop for keeping the program listening for update commands in excel file
2 - modify code in order to not have .pfx and .json in same folder of executable (embedded Auth in GOSip throws errors)
3 - Usage of Graph API (Go support is still too much in early stage as per May 2023)


