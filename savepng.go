package main

import (
	"bufio"
	"image"
	"image/png"
	"os"
)

func savePng(rgba *image.RGBA, fileName string) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	bf := bufio.NewWriter(outFile)
	err = png.Encode(bf, rgba)
	if err != nil {
		return err
	}

	err = bf.Flush()
	if err != nil {
		return err
	}

	return nil
}
