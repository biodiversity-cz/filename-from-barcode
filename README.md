# filename-from-barcode
Rename specimen Archive Master scans' (TIF images) according to the barcode detected inside. 

1) put the file into folder where tif files are present
2) run
3) all .tif files in this folder will be renamed, info is in terminal. In case the name already exists (there is more than one scan of the same specimen ID), the new name will include the original filename as a suffix.


## GO

### Ubuntu 24

```shell
sudo apt update
sudo apt install golang-go # libzbar-dev

go mod init barcode_rename  
go mod tidy
```

```shell
go build -o barcode_rename 
```

## Python

### Ubuntu

Download only the [script](barcode_rename.py) with [reuirements](requirements.txt) and run
```shell
python3 -m venv myenv
source myenv/bin/activate
pip install -r requirements.txt

python barcode_rename.py

deactivate
```

### Windows 
[solve dependencies](https://ruvi-d.medium.com/getting-zbarlight-to-work-on-windows-a3dc643dba18)
```shell
docker run -v "$(pwd):/src/" cdrx/pyinstaller-windows:python3 "pyinstaller --onefile  --specpath /src --hidden-import=pyzbar.pyzbar barcode_rename.py"
```