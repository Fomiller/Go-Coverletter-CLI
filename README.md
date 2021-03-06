# Scribe Command Line Tool

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
If you would like to use two words for a field such as ```Extracurricular Activity```, concatinate the two strings with an underscore ex: ```Extracurricular_Activity```
Example.
```
{{DATE}}
{{NAME}}
{{ADDRESS}}
{{EXTRACURRICULAR_ACTIVITY}}
```

## Populating templates from Google Sheets
Make a copy of this file ```https://docs.google.com/spreadsheets/d/16gT3TQ5BveYpaU56PZkoKfCoQNagb_fRHKd0NAxHmkg/edit?usp=sharing```.  
Make sure to keep the Unique File ID Column. This Column can be appended to your templates file name to ensure that each file has a unique file name. 

## Commands
- create
- create from sheet
- show template fields
- download
- download multiple
- delete

## Flags
- -d, --download
- -n, --name
- -p, --path
- -s, --sheet
- -t, --template

## ToDo List
- [ ] Allow for setting document type on download. 
- [ ] Add descriptions to all commands
- [ ] Remove unnecessary commands
- [ ] Parse command refactor
- [ ] list command refactor
- [ ] Add screenshots and gifs 
- [ ] Add example template for users, that works with example sheet.
- [ ] create a list all files cmd that where you can select a file then are prompted with another select list of what you would like to do with the file.
- [ ] support being able to use multiple sheets inside a spreadsheet
- [x] Create interactive mode for createfromsheet cmd.
- [x] Add createfromsheet to scribe interactive mode.
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

## Packages used  
drive  
https://godoc.org/google.golang.org/api/drive/v3  

surveyV2  
https://godoc.org/gopkg.in/AlecAivazis/survey.v2

cobra package  
https://godoc.org/github.com/spf13/cobra
