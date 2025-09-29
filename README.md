# filename-from-barcode
Rename specimen Archive Master scans' (TIF images) according to the barcode detected inside. 

1) put the file into folder where tif files are present
2) run
3) all .tif files in this folder will be renamed, info is in terminal. In case the name already exists (there is more than one scan of the same specimen ID), the new name will include the original filename as a suffix.

## Run 
Download from the [Releases](https://github.com/biodiversity-cz/filename-from-barcode/releases)

## Build
### Ubuntu
```shell
sudo apt-get install libzbar0

python3 -m venv myenv
source myenv/bin/activate

pyinstaller --onefile --hidden-import=pyzbar.pyzbar --hidden-import=PIL --collect-binaries pyzbar barcode_rename.py

deactivate
```
for debugging:
```shell
python3 -m venv myenv
source myenv/bin/activate
pip install -r requirements.txt

python barcode_rename.py

deactivate
```

### Windows
Let's have a Windows with Python installed:
1) ```pip install pyinstaller pyzbar Pillow```
2) ```pyinstaller --onefile --hidden-import=pyzbar.pyzbar --hidden-import=PIL --collect-binaries pyzbar --add-binary "libiconv.dll;." barcode_rename.py```

There should be possible to build it via Docker, but I did not found a working solution
```shell
docker run -v "$(pwd):/src/" cdrx/pyinstaller-windows:python3  bash -c "pyinstaller --onefile --specpath /src --hidden-import=pyzbar.pyzbar --add-binary \"/src/libiconv.dll;_MEIPASS\" --add-binary \"/src/libzbar64-0.dll;_MEIPASS\" barcode_rename.py"
```


[//]: # (obligatory branding for EOSC.CZ)
<hr style="margin-top: 100px; margin-bottom: 20px">

<p style="text-align: left"> <img src="https://webcentrum.muni.cz/media/3831863/seda_eosc.png" alt="EOSC CZ Logo" height="90"> </p>
This project output was developed with financial contributions from the EOSC CZ initiative throught the project National Repository Platform for Research Data (CZ.02.01.01/00/23_014/0008787) funded by Programme Johannes Amos Comenius (P JAC) of the Ministry of Education, Youth and Sports of the Czech Republic (MEYS).

<p style="text-align: left"> <img src="https://webcentrum.muni.cz/media/3832168/seda_eu-msmt_eng.png" alt="EU and MŠMT Logos" height="90"> </p>
