from PyInstaller.utils.hooks import collect_data_files

a = Analysis(
    ['barcode_rename.py'],
    pathex=['/src'],
    binaries=[],
    hookspath=[],
    excludes=[],
)

pyz = PYZ(a.pure, a.zipped_data)

exe = EXE(
    pyz,
    a.scripts,
    [],
    exclude_binaries=True,
    name='barcode_rename',
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    console=True,
)

coll = COLLECT(
    exe,
    a.datas,
    a.binaries,
    a.zipfiles,
    copy_dependents=True,
    upx=True,
    name='barcode_rename',
)
