package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/fatih/color"
)

func Scanner(url string, jobs <-chan string, wg *sync.WaitGroup, client *http.Client, outFile *os.File) {
	defer wg.Done()
	for word := range jobs {
		target := url + "/" + word

		resp, err := client.Get(target)
		if err != nil {
			continue
		}
		resp.Body.Close()

		var result string
		switch resp.StatusCode {
		case 200:
			result = fmt.Sprintf("[200] Access Granted:  %s", target)
			color.Green(result)
		case 301:
			result = fmt.Sprintf("[301] Redirected:      %s", target)
			color.Yellow(result)
		case 403:
			result = fmt.Sprintf("[403] Access Denied:   %s", target)
			color.Red(result)
		default:
			continue
		}

		if outFile != nil {
			fmt.Fprintln(outFile, result)
		}
	}
}
