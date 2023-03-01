package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func downloadPage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Cannot download page with url = ", url)
		return nil
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Cannot read page data with url = ", url)
	}

	return data
}

func writeToFile(data []byte, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func getDirectory(name string) {
	err := os.Mkdir(name, 0750)
	if err != nil {
		log.Fatal("Cannot create directory")
	}
	err = os.Chdir(name)
	if err != nil {
		log.Fatal("Cannot change work dir")
	}
}

func wget(url string, outPath string, depth int, timeout time.Duration) {
	getDirectory(outPath)
	wgetRec(url, url, depth, 0, 1, timeout)
}

func wgetRec(rootUrl string, curUrl string, depth int, curDepth int, pageInd int, timeout time.Duration) {
	if curDepth == depth {
		return
	}

	data := downloadPage(curUrl)

	if data != nil {
		links := getLinks(data)

		linksNorm := make([]string, len(links))
		copy(linksNorm, links)
		normalizeLinks(linksNorm, rootUrl, curUrl)

		for i := 0; i < len(links); i++ {
			oldLink := []byte(links[i])
			newLink := []byte(strconv.Itoa(curDepth+1) + "_" + strconv.Itoa(i) + ".html")
			copy(data, bytes.ReplaceAll(data, oldLink, newLink))
		}

		err := writeToFile(data, strconv.Itoa(curDepth)+"_"+strconv.Itoa(pageInd)+".html")
		if err != nil {
			log.Println("Cannot write file: ", err)
		}

		for i := 0; i < len(linksNorm); i++ {
			time.Sleep(timeout)
			wgetRec(rootUrl, linksNorm[i], depth, curDepth+1, i, timeout)
		}
	}
}

func getLinks(body []byte) []string {
	var links []string
	bodyReader := bytes.NewReader(body)
	z := html.NewTokenizer(bodyReader)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
	}
}

func normalizeLinks(links []string, rootUrl string, parentUrl string) {
	for i, link := range links {
		if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
			if strings.HasPrefix(link, "/") {
				locator, err := url.Parse(rootUrl)
				if err != nil {
					log.Println("Cannot parse url: ", locator)
				}
				links[i] = locator.Scheme + "://" + urlAddSuffixIfNeeded(locator.Host) + link[1:]
			} else {
				links[i] = urlAddSuffixIfNeeded(parentUrl) + link
			}
		}
	}
}

func urlAddSuffixIfNeeded(url string) string {
	if !strings.HasSuffix(url, "/") {
		return url + "/"
	} else {
		return url
	}
}

func getDepth(scanner *bufio.Scanner) int {
	for {
		fmt.Println("Select depth, it must be a number>0")
		scanner.Scan()
		depth, err := strconv.Atoi(scanner.Text())
		if err == nil && depth > 0 {
			return depth
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter directory name (files will be stored in current directory+name)")
	scanner.Scan()
	path := scanner.Text()

	fmt.Println("Enter full url, example: https://gobyexample.com/")
	scanner.Scan()
	locator := scanner.Text()

	depth := getDepth(scanner)

	wget(locator, path, depth, 3)

	fmt.Println("Download completed")
}
