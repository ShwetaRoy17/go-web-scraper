package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	
)

func createFileName(url string) string {
	url = strings.Replace(url, "/", "_", -1)

	return url + ".html"
}

func main() {
	url := "https://vim.rtorr.com/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
	}

	defer resp.Body.Close()
	fileName := createFileName(url)
	file,err := os.Create(fileName)
	if(err!=nil){
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving the file:", err)
	}

	fmt.Println("File saved successfully as", fileName)

}
