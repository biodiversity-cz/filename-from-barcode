import os
import re
from pyzbar.pyzbar import decode
from PIL import Image

Image.MAX_IMAGE_PIXELS = None  # no restriction for image size (!)

def get_barcode_from_image(image_path):
    try:
        image = Image.open(image_path)
        barcodes = decode(image)

        if barcodes:
            return barcodes[0].data.decode('utf-8')
        else:
            print(f"Barcode not found in {image_path}")
            return None
    except Exception as e:
        print(f"Error reading barcode from {image_path}: {e}")
        return None

def sanitize_filename(filename):
    return re.sub(r'[^a-zA-Z0-9]', '-', filename.upper())

def rename_files_in_directory(directory):
    script_directory = os.path.dirname(os.path.realpath(__file__))

    for filename in os.listdir(script_directory):
        if filename.lower().endswith('.tif'):
            image_path = os.path.join(script_directory, filename)

            barcode_value = get_barcode_from_image(image_path)

            if barcode_value:
                new_name = f"{sanitize_filename(barcode_value)}.tif"
                new_image_path = os.path.join(script_directory, new_name)

                if os.path.exists(new_image_path):
                    new_name = f"{sanitize_filename(barcode_value)}_{os.path.splitext(filename)[0]}.tif"
                    new_image_path = os.path.join(script_directory, new_name)

                try:
                    os.rename(image_path, new_image_path)
                    print(f"Renamed {filename} to {new_name}")
                except Exception as e:
                    print(f"Error renaming {filename}: {e}")

if __name__ == '__main__':
    rename_files_in_directory(os.getcwd())
