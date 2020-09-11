# Scribe CLI

## Tech Used
* Go
* Cobra
* Survey

ToDo Features
- [ ] create a config/setup command
- [ ] Ability to read from google sheets / csv files
- [ ] Allow for taking in objects for fields to be edited
- [x] Add ability to Delete files from drive
- [x] Ability to specify template.
- [x] Download file after filling out template
- [ ] Set download output path dynamically, flag?
- [x] Ability to parse a template and return the fields needed to fill out the template
- [ ] Create utils.go for helper functions
- [x] interactive command line ability with survey
- [ ] break cobra commands into subcommands
- [ ] list all available functions at root command with survey
- [ ] Research how to install and setup on end user computer without go installed. 


## Useful links

strings to map
https://stackoverflow.com/questions/35663892/parse-string-into-map-golang/35664442

tutorial  
https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html#adding-flags

promptui package  
https://hackernoon.com/improve-your-command-line-go-application-with-promptui-258ebe9eed1  

https://github.com/manifoldco/promptui  

https://manifold.co/blog/improve-your-command-line-go-application-with-promptui-6c4e6fb5a1bc

https://gist.github.com/jdaily/ae9640750d1d73170312963e720075d6

Survey
https://github.com/AlecAivazis/survey

https://godoc.org/gopkg.in/AlecAivazis/survey.v2

testings

cobra package  
https://godoc.org/github.com/spf13/cobra

google devleoper playground  
https://developers.google.com/oauthplayground/?code=4/2wGesxCBwoxpcCougQptTbjXydo7P1XOPG-iNt9niV4eDOLn9ierxwXvDME_aLfEpnhqzIuhttCv8arvqzDVYvA&scope=https://www.googleapis.com/auth/documents%20https://www.googleapis.com/auth/drive

google docs API    
https://developers.google.com/docs/api/how-tos/documents


google drive API  
https://developers.google.com/drive/api/v3/reference/files/copy

## godocs references
drive  
https://godoc.org/google.golang.org/api/drive/v3
