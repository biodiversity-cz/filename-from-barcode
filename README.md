# filename-from-barcode
Rename specimen Archive Master scans' (TIF images) according to the barcode detected inside. 

1) put the file into folder where tif files are present
2) run
3) all .tif files in this folder will be renamed, info is in terminal. In case the name already exists (there is more than one scan of the same specimen ID), the new name will include the original filename as a suffix.

Sadly, due the dependencies, it is not possible to build the script for different OS from scratch. Use GO on your platform and build it yourself.. :/

## Ubuntu 24

```shell
sudo apt update
sudo apt install golang-go libzbar-dev

go mod init barcode_rename  
go mod tidy
```

```shell
go build -o barcode_rename 
```

 