package utils

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

const (
	batchSize   = 1000
	workerCount = 100
)

type CSVProcessor struct {
	wg       sync.WaitGroup
	jobsChan chan map[string]string
}

func NewCSVProcessor() *CSVProcessor {
	return &CSVProcessor{
		jobsChan: make(chan map[string]string, batchSize),
	}
}

func ProssesCSV(file *multipart.FileHeader, enterFunc func(string, string) *customerror.CustomError) *customerror.CustomError {
	processor := NewCSVProcessor()
	return processor.ProcessCSVFile(file, enterFunc)
}

func (p *CSVProcessor) ProcessCSVFile(file *multipart.FileHeader, enterFunc func(string, string) *customerror.CustomError) *customerror.CustomError {
	urlIndex := -1
	titleIndex := -1

	f, err := file.Open()
	if err != nil {
		return customerror.NewCustomError(err, err.Error(), 400)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	//read header and find index of url and title
	header, err := reader.Read()
	if err != nil {
		return customerror.NewCustomError(err, err.Error(), 400)
	}
	for i, v := range header {
		if v == "title" {
			titleIndex = i
		} else if v == "url" {
			urlIndex = i
		}
	}
	if titleIndex == -1 || urlIndex == -1 {
		return customerror.NewCustomError(errors.New("title or url not found"), "title or url not found", 400)

	}
	// Start workers
	for i := 0; i < workerCount; i++ {
		go p.worker(i, enterFunc)
	}

	// Read and process CSV
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return customerror.NewCustomError(err, err.Error(), 400)
		}

		data := map[string]string{"title": "", "url": ""}
		for i, v := range record {
			if i == titleIndex {
				data["title"] = v
			} else if i == urlIndex {
				data["url"] = v
			}
		}

		//check if title or url is empty
		if data["title"] == "" || data["url"] == "" {
			log.Printf("Title or url is empty")
			continue
		}

		p.wg.Add(1)
		p.jobsChan <- data
	}

	close(p.jobsChan)
	p.wg.Wait()

	return nil
}

func (p *CSVProcessor) worker(id int, enterFunc func(string, string) *customerror.CustomError) *customerror.CustomError {
	counter := 0
	for data := range p.jobsChan {
		//scraping
		title := data["title"]
		content, err := p.scraping(data["url"])
		if err != nil {
			log.Println("Error scraping url:", data["url"])
			continue
		}

		// Enter data
		cuserr := enterFunc(title, content)
		if cuserr != nil {
			return cuserr
		}
		counter++
		log.Printf("Worker %d processed %d records success with title %s", id, counter, title)
	}

	return nil
}

func (c *CSVProcessor) scraping(url string) (string, error) {
	// stringsbuilder
	var content strings.Builder

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return "", errors.New("status code error")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the content
	doc.Find("div.detail__body-text").Contents().Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "#text" {
			content.WriteString(s.Text())
		} else if goquery.NodeName(s) == "p" {
			content.WriteString(s.Text())
		}
	})

	return content.String(), nil
}
