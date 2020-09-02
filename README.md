# Scribe CLI

## Tech Used
* Go
* Cobra
* promptui

ToDo Features
- [ ] create a config/setup command
- [ ] Ability to read from google sheets / csv files
- [ ] Allow for taking in objects for fields to be edited
- [ ] Delete files from drive command
- [x] Download file after filling out template
- [ ] Set download output path dynamically, flag?
- [ ] Ability to parse a template and return the fields needed to fill out the template
- [ ] interactive command line ability with promptui
- [ ] Research how to install and setup on end user computer without go installed. 


## Useful links

tutorial  
https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html#adding-flags

promptui package  
https://hackernoon.com/improve-your-command-line-go-application-with-promptui-258ebe9eed1  

https://github.com/manifoldco/promptui  

https://manifold.co/blog/improve-your-command-line-go-application-with-promptui-6c4e6fb5a1bc

https://gist.github.com/jdaily/ae9640750d1d73170312963e720075d6

similar to inquirer js
https://github.com/AlecAivazis/survey

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
