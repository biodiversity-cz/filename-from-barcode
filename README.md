# filename-from-barcode
Rename specimen Archive Master scans' (TIF images) according to barcode detected inside. 
1) put the file into folder where tif files are present
2) run it
3) all .tif files in this folder will be renamed, info is in terminal. In case the name already exists (there is more than one scan of the same specimen ID). th enew name will include the original filename as a suffix.

## compilation (Ubuntu 24)
Sadly, due the dependencies, it is not possible to build the script for different OS from scratch. Use GO on your platform and build it.. :/

### run only for the first time
```shell
sudo apt update
sudo apt install golang-go libzbar-dev

go mod init barcode_rename  
go mod tidy
go get github.com/mordfustang21/gozbar
go get golang.org/x/image/tiff

 
```

```shell
go build -o barcode_rename #build for actual OS
```

 