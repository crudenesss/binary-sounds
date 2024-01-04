package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var text string
	var binaryData string

	file0 := os.Args[1]
	file1 := os.Args[2]
	output := os.Args[3]

	fmt.Print("Enter text: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()

		binaryData = textToBinary(text)
		fmt.Printf("Binary representation: %s\n", binaryData)
	}

	// Create a file to hold the filenames
	createFile(binaryData, file0, file1)

	args := make([]string, 0)
	// Create the ffmpeg command
	args = append(
		args,
		"-f", "concat", "-safe", "0", "-i", "temp.txt", output, "-y",
	)
	cmd := exec.Command("ffmpeg", args...)

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("An error occured while creating output file.")
	} else {
		fmt.Println("The audiofile is successfully created!")
	}

	os.Remove("temp.txt")
}

func intToBinary(value byte) string {

	binaryString := ""
	for i := 7; i >= 0; i-- {
		bit := (value >> i) & 1
		binaryString += fmt.Sprintf("%d", bit)
	}

	return binaryString
}

func textToBinary(text string) string {

	asciiValues := []byte(text)

	var binaryData string
	for _, v := range asciiValues {
		bSymbol := intToBinary(v)
		binaryData += bSymbol
	}

	return binaryData
}

func createFile(seq string, file0 string, file1 string) {
	hFile, err := os.Create("temp.txt")
	if err != nil {
		fmt.Print("An error occured while creating temporary file.\n")
	}

	for _, c := range seq {
		switch c {
		case '0':
			content := fmt.Sprintf("file '%s'\n", file0)
			hFile.Write([]byte(content))
		case '1':
			content := fmt.Sprintf("file '%s'\n", file1)
			hFile.Write([]byte(content))
		}
	}

	hFile.Close()
	return
}
