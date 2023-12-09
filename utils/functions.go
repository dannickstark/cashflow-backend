package utils

import (
	"fmt"
	"os"
)

// A function to save JSON data into a file
func SaveFile(data string, filename string) error {
	// Write the JSON data to the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	l, err := f.WriteString(data)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
