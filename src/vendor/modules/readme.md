# Modules folder

Modules folder for herb-go apps.

This folder can be used in both vendor mode and  gomod mode

## Modules location

### Vendor mode
In vendor mode,this folder should be put in "\<app\>/src/vendor" folder.

### Go mod mode
In go mod mode.this folder shoud be "\<app\>/src", and codes below shoud be put in go.mod file in  "\<app\>/src"

    replace modules => ./modules

    require (
    	modules v0.0.0
    )

## File "go.mod" in this folder

You should not edit this file.

The "go.mod" file is used by herb-go cli tool.

You should keep first line  of  "go.mod" file  as

    module modules

## Change mode.

###  Vendor mode to Go mod mode
* Move this folder under "\<app\>/src" folder
* Run cmd go mod init <appname> in "\<app\>/src" folder
* Edit go.mod.

### Go mod mode to Vendor mode
* Move this folder under "\<app\>/src/vendor" folder
* Remove go.mod file