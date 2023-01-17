# security-cli

## build commands

go build -o ./dist/security-cli

## run commands

./dist/app args

scan repo

./dist/app.exe scm scan https://github.com/laverdet/isolated-vm/tree/v4.3.3


## Prerequisits
go install github.com/spf13/cobra-cli@latest


### to add new command

cobra-cli add command-name

