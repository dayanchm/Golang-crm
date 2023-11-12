package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

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

func GetDovizKurlari() (TcmbApiResponse, error) {
	url := "https://www.tcmb.gov.tr/kurlar/today.xml"

	resp, err := http.Get(url)
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
