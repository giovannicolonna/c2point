# c2point
c2point

# Install Go: 
If you haven't already, download and install Go from the official website: https://golang.org/dl/. Follow the installation instructions for your operating system.

# Install dependencies: 
The program depends on several third-party packages, which can be installed using the following command:


```

go get github.com/koltyakov/gosip
go get github.com/koltyakov/gosip/api
go get github.com/koltyakov/gosip/auth/saml
go get github.com/tealeg/xlsx

```

Save the program: Save the program code to a file named c2point.go.

# Configure authentication: 
Replace the placeholder values in the authCnfg struct with your SharePoint auth info:

Here is an example of how to initialize AuthCnfg with the necessary fields for SharePoint Online authentication using OAuth2:

```
auth := &gosip.AuthCnfg{
    Strategy:    "saml",
    LoginURL:    "https://login.microsoftonline.com/<tenant-id>/saml2",
    LogoutURL:   "https://<tenant-name>.sharepoint.com/_layouts/15/SignOut.aspx",
    RelyingParty: "<relying-party>",
    Username:    "<username>",
    Password:    "<password>",
    ClientID:    "<client-id>",
    ClientSecret: "<client-secret>",
    Realm:       "<realm>",
    SiteURL:     "<site-url>",
}
```
Make sure to replace the placeholders (tenant-id, tenant-name, relying-party, username, password, client-id, client-secret, realm, and site-url) with the actual values for your SharePoint environment.

To retrieve the values for RelyingParty, ClientID, ClientSecret, and Realm, you will need to register an application in Azure Active Directory (AD) and grant it permissions to access SharePoint Online resources.

Here are the steps to retrieve these values:

1)    Go to the Azure portal (https://portal.azure.com/) and sign in with your Microsoft account.

2)    Click on the "Azure Active Directory" service and select the "App registrations" option.

3)    Click on the "New registration" button to create a new application.

4)    Provide a name for the application and select "Accounts in this organizational directory only" as the supported account type.

5)    For the "Redirect URI" field, enter a placeholder URL such as "https://localhost".

6)   After creating the application, you will see its "Application (client) ID" on the Overview page. This is the value you will use for ClientID.

7)    Under the "Certificates & secrets" tab, click on the "New client secret" button to generate a new secret. This will display the secret value which you will use for ClientSecret.

8)    Under the "API permissions" tab, click on the "Add a permission" button and select "Microsoft Graph" as the API.

9)    Select the "Application permissions" option and search for the "Sites.FullControl.All" permission.

10)    Click on the "Add permission" button to grant the permission to the application.

11)    Finally, under the "Endpoints" tab, you will find the values for RelyingParty and Realm.

Note that the value for RelyingParty is the URL that is displayed under the "Federation Metadata Document" endpoint.

The value for Realm is the value that appears after https://sts.windows.net/ in the "Issuer URL" endpoint.

Make sure to copy these values exactly as they appear, including any special characters such as forward slashes or colons.

Once you have retrieved these values, you can use them to initialize the AuthCnfg struct for SharePoint Online authentication in your Go program.






# Update file path
  Update the file path in the endpoint variable to point to your SharePoint file.

# Compile and run the program:
  Open a terminal or command prompt and navigate to the directory containing the c2point.go file. Run the following commands to compile and execute the program:

```

go build
./c2point

```

If you're on Windows, the compiled executable will be named c2point.exe. In that case, run the following command instead:

```
c2point.exe
```

The program will read the instructions from the Excel file and execute them as command-line commands on your machine. On Windows, the commands will be executed using the cmd.exe shell, while on Unix-based systems they will be executed using the /bin/bash shell.
