package helpers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewsApiResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Currency struct {
	Unit         int    `xml:"Unit"`
	Isim         string `xml:"Isim"`
	CurrencyName string `xml:"CurrencyName"`
	ForexBuying  string `xml:"ForexBuying"`
	ForexSelling string `xml:"ForexSelling"`
}

type TcmbApiResponse struct {
	Tarih      string     `xml:"Tarih"`
	Date       string     `xml:"Date"`
	BultenNo   int        `xml:"Bulten_No"`
	Currencies []Currency `xml:"Currency"`
}

type AppConfig struct {
	DovizURL string
	NewsURL  string
}

var appConfig AppConfig

func init() {
	loadConfig()
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	appConfig.DovizURL = os.Getenv("DOVIZ_URL")
	appConfig.NewsURL = os.Getenv("NEWS_URL")
}

func GetDovizKurlari() (TcmbApiResponse, error) {
	resp, err := http.Get(appConfig.DovizURL)
	if err != nil {
		return TcmbApiResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TcmbApiResponse{}, err
	}

	var result TcmbApiResponse
	err = xml.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("XML Unmarshal Error:", err)
		return TcmbApiResponse{}, err
	}

	return result, nil
}

func GetNews() (NewsApiResponse, error) {
	resp, err := http.Get(appConfig.NewsURL)
	if err != nil {
		return NewsApiResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewsApiResponse{}, err
	}

	var result NewsApiResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("JSON Unmarshal Error:", err)
		return NewsApiResponse{}, err
	}

	var filteredArticles []Article
	for _, article := range result.Articles {
		if article.URLToImage != "" {
			filteredArticles = append(filteredArticles, article)
		}
	}
	result.Articles = filteredArticles
	return result, nil
}
