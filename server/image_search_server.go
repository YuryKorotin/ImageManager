package image_manager

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

import . "../models"

type ApiKeysConfiguration struct {
  YandexTranslationApiKey string
  UnsplashApiKey string
}

func GetTranslations(body []byte) (*TranslationResponse, error) {
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

	translationResponse, err := GetTranslations([]byte(body))
	translation := "funny"

	if (err == nil) {
		translation = translationResponse.Texts[0]
	}

	return translation
}

func images(w http.ResponseWriter, r *http.Request) {
	translationResultResponse := requestTranslationForQuery(r.URL.Query().Get("filter"))
	translationResult := parseTranslationResponse(translationResultResponse)
	imagesResult := RequestImageFromUnsplash(translationResult)
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

func FormatUnsplashSearchResult(searchResponseString string) UnsplashResponse{
  byteValue := []byte(searchResponseString)

  var response UnsplashResponse

  json.Unmarshal(byteValue, &response)

  return response
}
func FormatGiphySearchResult(searchResponseString string) GiphyResponse {
  byteValue := []byte(searchResponseString)

  var response GiphyResponse

  json.Unmarshal(byteValue, &response)

  return response
}



func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/images", images)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
