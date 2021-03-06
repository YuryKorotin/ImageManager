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

type GatheredCollectionResult struct {
  GiphyImages []string
  UnsplashImages []string
  PexelImages []string
  QwantImages []string
}

type ImageCollectionResult struct {
  ImageUrls []string
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

func requestsSequency(query string) {
  actions := map[string]func(query string){
    "unsplash": func(query string) { 
       responseString := RequestImageFromUnsplash(query)
       FormatUnsplashSearchResult(responseString)
     },
    "qwant": func(query string) { RequestImageFromQwant(query) },
  }
}

func ParallelGet(query string, concurrencyLimit int) ImageCollectionResult {
  var result ImageCollectionResult

  semaphoreChan := make(chan struct{}, concurrencyLimit)
  resultsChan :=  make(chan *result)

  defer func() {
     close(semaphoreChan)
     close(resultsChan)
  }()

  requests := requestsSequency(query)

  for key, value := range m {
    value(query)
    go func(i int, url string) {
      semaphoreChan <- struct{}{}
      res := value(query)
      result := &ImageCollectionResult()

      resultsChan <- result
      <-semaphoreChan
    }(key, value)
  }

  var finalResult GatheredCollectionResult

  requestIndex := 0

  for {
    result := <-resultsChan

    if (len(requests) == requestIndex) {
	break
    }
    requestIndex += 1
  }

  return result
}
