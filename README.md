# filename-from-barcode
Rename specimen scan images according to barcode detected inside

## Windows 

### var A
Download the exe file from releases

### var B
build the file on Linux by yourself
```shell
docker run -v "$(pwd):/src/" cdrx/pyinstaller-windows:python3 
```


## Linux 

Download only the [script](barcode_rename.py) and run
```shell
python3 -m venv myenv
source myenv/bin/activate
pip install -r requirements.txt

python barcode_rename.py

deactivate
```