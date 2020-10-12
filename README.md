# Scribe CLI

## ToDo Features
- [ ] Allow for setting document type on download. 
- [ ] support being able to use multiple sheets inside a spreadsheet
- [ ] create a list all files cmd that where you can select a file then are prompted with another select list of what you would like to do with the file.
- [ ] Allow for taking in objects for fields to be edited
- [ ] Create interactive mode for createfromsheet cmd.
- [x] Add createfromsheet to scribe interactive mode.
- [ ] Add descriptions to all commands
- [x] Add config file instructions to README.md
- [x] Rename NewFileName variable to Name
- [x] create function for folderGeneration with files.
- [x] Ability to read from google sheets / csv files
- [x] configure google sheets api
- [x] Add create from file command
- [x] implement feature of cutting out the string "Template" from a template name when saving a file
- [x] create test files
- [x] list all available functions at root command with survey
- [x] create a config file
- [x] Add ability to Delete files from drive
- [x] Ability to specify template.
- [x] Download cmd clean up.
- [x] Download cmd interactive mode.
- [x] Download cmd set output path with flag.
- [x] Download cmd remove -d flag check.
- [x] Download file after filling out template
- [x] Ability to parse a template and return the fields needed to fill out the template
- [x] Create utils.go for helper functions
- [x] interactive command line ability with survey

## Install
```bash
  go get github.com/fomiller/scribe

  cd $GOPATH/src/github.com/fomiller/scribe
```
Your first time running ```scribe``` you will be asked to create credentials from the google developer console. Follow the instructions given by scribe to complete this process.  
Once you have created and saved your credentials run ```go install```  
Test to make sure your application is working correctly by running ```scribe``` in your command line.  
To run Scribe as a command, make sure that your $GOPATH/bin is added to your system $PATH

## Configure
Use the ```config.yaml``` file to set configuration settings for Scribe.  
Scribe automatically downloads files to the users ```$HOME\scribe``` path on Mac/Linux systems, and ```%USERPROFILE%\scribe``` path on Windows systems.  

You can define a custom default download path by setting ```download.path``` to your desired full path.  

If you make changes to your ```config.yaml``` file make sure to run ``` go install``` after you save your changes.  

## Creating and using templates
In Google Docs create a new template file.  
Create a template field within your document, indicated by ```{{}}```.  
Inside the ```{{}}``` add a Uppercase string describing the data you would like to insert into the template field.  
Example.
```
{{DATE}}
{{NAME}}
{{ADDRESS}}
```

## godocs references
drive  
https://godoc.org/google.golang.org/api/drive/v3  

surveyV2  
https://godoc.org/gopkg.in/AlecAivazis/survey.v2

cobra package  
https://godoc.org/github.com/spf13/cobra

## Useful links

* strings to map
  * https://stackoverflow.com/questions/35663892/parse-string-into-map-golang/35664442

* promptui package  
  * https://hackernoon.com/improve-your-command-line-go-application-with-promptui-258ebe9eed1  

  * https://github.com/manifoldco/promptui  

  * https://manifold.co/blog/improve-your-command-line-go-application-with-promptui-6c4e6fb5a1bc

  * https://gist.github.com/jdaily/ae9640750d1d73170312963e720075d6

* Survey
  * https://github.com/AlecAivazis/survey

* google devleoper playground  
  * https://developers.google.com/oauthplayground/?code=4/2wGesxCBwoxpcCougQptTbjXydo7P1XOPG-iNt9niV4eDOLn9ierxwXvDME_aLfEpnhqzIuhttCv8arvqzDVYvA&scope=https://www.googleapis.com/auth/documents%20https://www.googleapis.com/auth/drive

* google docs API    
  * https://developers.google.com/docs/api/how-tos/documents

* google drive API  
  * https://developers.google.com/drive/api/v3/reference/files/copy