package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/page"

)

func createFileName(url string) string {
	url = strings.Replace(url, "/", "_", -1)

	return url + ".pdf"
}

func main() {
	url := "https://vim.rtorr.com/"
	outputFile := createFileName(url)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var pdfData []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfData,_, err = page.PrintToPDF().
				WithPrintBackground(true). // include background colors/images
				WithPaperWidth(8.27).      // A4 width
				WithPaperHeight(11.69).    // A4 height
				Do(ctx)
			return err
		}),
	)

	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}
    if err := os.WriteFile(outputFile, pdfData, 0644); err != nil {
		fmt.Println("❌ Error saving PDF:", err)
		return
	}
	

// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("Error fetching the URL:", err)
	// }

	// defer resp.Body.Close()
	// fileName := createFileName(url)
	// file,err := os.Create(fileName)
	// if(err!=nil){
	// 	fmt.Println("Error creating the file:", err)
	// 	return
	// }
	// defer file.Close()

	// _, err = io.Copy(file, resp.Body)
	// if err != nil {
	// 	fmt.Println("Error saving the file:", err)
	// }

	fmt.Println("File saved successfully as", outputFile)

}
