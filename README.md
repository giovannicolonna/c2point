# c2point
c2point

Install Go: If you haven't already, download and install Go from the official website: https://golang.org/dl/. Follow the installation instructions for your operating system.

Install dependencies: The program depends on several third-party packages, which can be installed using the following command:



go get github.com/koltyakov/gosip
go get github.com/koltyakov/gosip/api
go get github.com/koltyakov/gosip/auth/saml
go get github.com/tealeg/xlsx

Save the program: Save the program code to a file named main.go.

Configure authentication: Replace the placeholder values in the authCnfg struct with your SharePoint site URL, username, and password.

Update file path: Update the file path in the endpoint variable to point to your SharePoint file.

Compile and run the program: Open a terminal or command prompt and navigate to the directory containing the c2point.go file. Run the following commands to compile and execute the program:

bash

go build
./c2point

If you're on Windows, the compiled executable will be named c2point.exe. In that case, run the following command instead:

c2point.exe

The program will read the instructions from the Excel file and execute them as command-line commands on your machine. On Windows, the commands will be executed using the cmd.exe shell, while on Unix-based systems they will be executed using the /bin/bash shell.
