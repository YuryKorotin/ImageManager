package main

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"github.com/peterhellberg/giphy"
	"github.com/tkanos/gonfig"
	"path/filepath"
)

type ApiKeysConfiguration struct {
  YandexTranslationApiKey string
  UnsplashApiKey string
}

type TranslationResponse struct {
  Code int `json:"code"`
  Lang string `json:"lang"`
  Texts []string `json:"text"`
}

type UnsplashLink struct {
  Download string `json:"download"`
}

type UnsplashResult struct {
  Links []UnsplashLink `json:"links"`
}

type UnsplashResponse struct {
  Total int `json:"total"`
  Results []UnsplashResult `json:"results"`
}

func getTranslations(body []byte) (*TranslationResponse, error) {
    var translation = new(TranslationResponse)
    err := json.Unmarshal(body, &translation)
    if(err != nil){
        fmt.Println("Parsing error:", err)
    }
    return translation, err
}

func parseTranslationResponse(res *http.Response) string{
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	translationResponse, err := getTranslations([]byte(body))
	translation := "funny"

	if (err == nil) {
		translation = translationResponse.Texts[0]
	}

	return translation
}

func images(w http.ResponseWriter, r *http.Request) {
	translationResultResponse := requestTranslationForQuery(r.URL.Query().Get("filter"))
	translationResult := parseTranslationResponse(translationResultResponse)
	imagesResult := requestImageFromUnsplash(translationResult)
	w.Write([]byte(imagesResult))
}

func requestTranslationForQuery(query string) *http.Response {
	locale := "ru-en"
	var rootUrl = "https://translate.yandex.net/api/v1.5/tr.json/translate?"

	configuration := ApiKeysConfiguration{}
	configPath, _ := filepath.Abs("image_manager/api_keys.json")
	err := gonfig.GetConf(configPath, &configuration)
	fmt.Println(err)

	apiKey := configuration.YandexTranslationApiKey
	fmt.Println(apiKey)
 
	values := url.Values{}
	values.Set("lang", locale)
	values.Set("key", apiKey)
	values.Set("text", query)

	fmt.Println(rootUrl + values.Encode())
	resp, _ := http.Get(rootUrl + values.Encode())

	fmt.Println(resp.StatusCode)

	return resp
}

func requestImageFromQwant(query string) string {
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

func formatUnsplashSearchResult(searchResponseString string) UnsplashResponse{
  byteValue := []byte(searchResponseString)

  var response UnsplashResponse

  json.Unmarshal(byteValue, &response)

  return response
}

func requestImageFromUnsplash(query string) string {
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

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/images", images)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
