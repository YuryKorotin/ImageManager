package image_manager

import (
	"fmt"
	"net/http"
	//"sort"
	//"time"
	"io/ioutil"
	"github.com/tkanos/gonfig"
	"path/filepath"
)

type ImageCollectionResult struct {
  GiphyImageUrl []string
  UnsplashImageUrl []string
}

func RequestImageFromQwant(query string) string {
	//var userAgentString = "Mozilla/5.0 (Windows NT 6.3; WOW64; rv:56.0) Gecko/20100101 Firefox/56.0"
	offset := 0
	locale := "en_en"
	count := 1
	var rootUrl = fmt.Sprintf(
		"https://api.qwant.com/api/search/news/?count=%d&q=%s&offset=%d&locale=%s",
		count,
		query,
		offset,
		locale)
	resp, _ := http.Get(rootUrl)

	var bodyString = "NO RESULTS"

	if resp.StatusCode == http.StatusOK {
		bodyBytes, erro := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
		fmt.Println(erro)
	}
	for k, v := range resp.Header {
		fmt.Printf("Headers, %s, %s \n", k, v)
	}

	fmt.Println(rootUrl)
	fmt.Println(resp.StatusCode)
	return bodyString
}

func RequestImageFromUnsplash(query string) string {
	configuration := ApiKeysConfiguration{}
	configPath, _ := filepath.Abs("image_manager/api_keys.json")
	err := gonfig.GetConf(configPath, &configuration)
	fmt.Println(err)

	clientId := configuration.UnsplashApiKey
	fmt.Println(clientId)
	rootUrl := fmt.Sprintf("https://api.unsplash.com/search/photos?client_id=%s&query=%s", clientId, query)

	resp, _ := http.Get(rootUrl)
	var bodyString = "NO RESULTS"

	if resp.StatusCode == http.StatusOK {
		bodyBytes, erro := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
		fmt.Println(erro)
	}
	fmt.Println(rootUrl)
	fmt.Println(resp.StatusCode)
	return bodyString
}

func ParallelGet(concurrencyLimit int) ImageCollectionResult {
	var result ImageCollectionResult

	return result
}
