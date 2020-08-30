package httpClient

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var sid string

func SetSID(newSID string) {
	sid = newSID
}

func PostRequest(url string, rawBody string) (res *http.Response, err error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(rawBody))
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "toster_sid", Value: sid})

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Referer", "https://qna.habr.com/user/sniggering_deus/answers?page=57")

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		return nil, errors.New("resp.StatusCode: " + strconv.Itoa(res.StatusCode))
	}

	return
}

func GetRequest(url string) (res *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "toster_sid", Value: sid})

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		return nil, errors.New(url + "\nresp.StatusCode: " + strconv.Itoa(res.StatusCode))
	}

	return
}
