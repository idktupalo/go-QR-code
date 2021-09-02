package main

import (
	"io/ioutil"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	read, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	strRead := string(read)
	QRname := "go-QR.png"

	err = qrcode.WriteFile(strRead, qrcode.Medium, 512, QRname)
	if err != nil {
		log.Fatal(err)
	}
}
