package main

import (
	"fmt"
	"image"
	_ "golang.org/x/image/tiff"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/makiuchi-d/gozxing"
    "github.com/makiuchi-d/gozxing/oned"

	"github.com/disintegration/imaging"
)

func getBarcodeFromImage(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("error opening image: %w", err)
	}
	defer file.Close()

	i, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("error decoding image: %w", err)
	}

	i = imaging.AdjustContrast(i, 50)
	i = imaging.Resize(i, 0, 2000,imaging.Lanczos )

	bitmap, err := gozxing.NewBinaryBitmapFromImage(i)
    if err != nil {
        return "", fmt.Errorf("error creating binary bitmap: %w", err)
    }

    hints := map[gozxing.DecodeHintType]interface{}{
            gozxing.DecodeHintType_POSSIBLE_FORMATS: []gozxing.BarcodeFormat{
                gozxing.BarcodeFormat_EAN_13,
                gozxing.BarcodeFormat_UPC_A,
                gozxing.BarcodeFormat_EAN_8,
                gozxing.BarcodeFormat_UPC_E,
                gozxing.BarcodeFormat_CODE_128,
                gozxing.BarcodeFormat_CODE_39,
                gozxing.BarcodeFormat_ITF,
                gozxing.BarcodeFormat_CODABAR,
            },
        }

    readers := []gozxing.Reader{
    		oned.NewCode128Reader(),
    		oned.NewCode93Reader(),
    		oned.NewCode39Reader(),
    		oned.NewCodaBarReader(),
    		oned.NewITFReader(),
    		oned.NewMultiFormatUPCEANReader(hints),
    	}
    for _, reader := range readers {
        result, err := reader.Decode(bitmap, nil)
        if err == nil {
            return result.GetText(), nil
        }
    }

    return "", fmt.Errorf("no barcode found")

}

func sanitizeFilename(filename string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return strings.ToUpper(re.ReplaceAllString(filename, "-"))
}

func renameFilesInDirectory(directory string) {
	scriptDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	err = filepath.Walk(scriptDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".tif") {
			barcodeValue, err := getBarcodeFromImage(path)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			newName := sanitizeFilename(barcodeValue) + ".tif"
			newPath := filepath.Join(scriptDirectory, newName)

			if _, err := os.Stat(newPath); err == nil {
				newName = sanitizeFilename(barcodeValue) + "_" + strings.TrimSuffix(info.Name(), ".tif") + ".tif"
				newPath = filepath.Join(scriptDirectory, newName)
			}

			err = os.Rename(path, newPath)
			if err != nil {
				fmt.Println("Error renaming", info.Name(), ":", err)
			} else {
				fmt.Println("Renamed", info.Name(), "to", newName)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}
}

func main() {
	renameFilesInDirectory(".")
}
