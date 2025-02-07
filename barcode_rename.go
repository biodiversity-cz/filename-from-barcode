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
	"github.com/makiuchi-d/gozxing/qrcode"
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

	bitmap, err := gozxing.NewBinaryBitmapFromImage(i)
    if err != nil {
        return "", fmt.Errorf("error creating binary bitmap: %w", err)
    }

	reader := qrcode.NewQRCodeReader()
    result, err := reader.Decode(bitmap, nil)
    if err != nil {
        return "", fmt.Errorf("error scanning image: %w", err)
    }

    return result.GetText(), nil
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
