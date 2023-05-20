package application

import (
	"github.com/PuerkitoBio/goquery"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	"net/http"
)

type UrlUsecase struct{}

func (UrlUsecase) UrlGet(header http.Header, url string) (*miikov1.UrlGetResponse, error) {
	
	if err := authorizationService.CheckLoggedIn(header); err != nil {
		return &miikov1.UrlGetResponse{}, err
	}

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{Transport: http.DefaultTransport}
	res, err := client.Do(req)
	if err != nil {
		return &miikov1.UrlGetResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return &miikov1.UrlGetResponse{}, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return &miikov1.UrlGetResponse{}, err
	}

	return &miikov1.UrlGetResponse{Title: doc.Find("title").Text()}, nil
}
