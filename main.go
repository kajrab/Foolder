package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	printBanner()

	url := flag.String("url", "", "Target URL")
	wordlist := flag.String("wordlist", "wordlists/default.txt", "Path to wordlist")
	workers := flag.Int("workers", 50, "Number of concurrent workers")
	timeout := flag.Int("timeout", 3, "HTTP timeout in seconds")
	output := flag.String("output", "", "Save results to file")
	flag.Parse()

	if *url == "" {
		log.Fatal("URL is required: --url http://target.com")
	}

	if !strings.HasPrefix(*url, "http://") && !strings.HasPrefix(*url, "https://") {
		*url = "http://" + *url
	}

	file, err := os.Open(*wordlist)
	if err != nil {
		log.Fatal("Could not open wordlist:", err)
	}
	defer file.Close()

	var outFile *os.File
	if *output != "" {
		outFile, err = os.Create(*output)
		if err != nil {
			log.Fatal("Could not create output file:", err)
		}
		defer outFile.Close()
	}

	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	jobs := make(chan string, 500)
	var wg sync.WaitGroup

	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go Scanner(*url, jobs, &wg, client, outFile)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- scanner.Text()
	}
	close(jobs)

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDone.")
}
