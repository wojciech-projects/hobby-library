package usecase

import (
	"crawler/src/domain"
	"crawler/src/parser"
	"net/http"
)

type HttpMangaDownloader struct{}

func (d *HttpMangaDownloader) DownloadMangaByUuid(uuid string) (series domain.Series, err error) {
	url := "https://www.amazon.co.jp/-/en/kindle-dbs/product/" + uuid

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	return parser.Parse(resp.Body)
}
