package wget

import (
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// Wget main structure
type Wget struct {
	BasePath  string
	links     []string
	client    *http.Client
	collector *colly.Collector
}

// New return instance of *Wget
func New() *Wget {
	return &Wget{
		links:     []string{},
		client:    &http.Client{},
		collector: colly.NewCollector(),
	}
}

// GetPath cut path from url
func (wget *Wget) GetPath(url string) string {
	dir := strings.TrimPrefix(url, "https://")
	dir = strings.TrimPrefix(dir, "http://")
	fmt.Println(dir)
	return wget.BasePath + "/" + dir
}

// GetFileName cut file name from url
func (wget *Wget) GetFileName(url string) string {
	return wget.GetPath(url) + "/" + path.Base(url)
}

// CreateDir makes dir
func (wget *Wget) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// GetPage download and copy file
func (wget *Wget) GetPage(url string) error {
	c := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error %v when create request", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("error when send request: %v", err)
	}

	wget.CreateDir(wget.GetPath(url))

	file, err := os.Create(wget.GetFileName(url))
	if err != nil {
		return fmt.Errorf("error when create file: %v", err)
	}
	io.Copy(file, resp.Body)

	return nil
}

// VisitAndGet uses colly for gettint list of url and visiting it
func (wget *Wget) VisitAndGet(url string) {
	c := colly.NewCollector()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		wget.GetPage(r.URL.String())
	})

	c.Visit(url)
}

func (wget *Wget) Start() {
	if len(os.Args) < 2 {
		log.Fatal("wget: missing URL")
	}

	url := os.Args[1]

	if len(os.Args) == 3 {
		wget.BasePath = os.Args[2]
	}

	wget.VisitAndGet(url)
}
