package main

import (
  "testing"
  "io/ioutil"
  "path/filepath"
)

func TestGetTranslations(t *testing.T) {
  var body = "{\"code\":102, \"lang\": \"ru-en\", \"text\": [\"translation\"] }"
  translation, _ := getTranslations([]byte(body))
  if (translation.Texts[0] != "translation") {
    t.Errorf("Translation was incorrect, got: %s, want: %s.", translation.Texts[0], "translation")
  }
}

func TestFormatUnsplashSearhResult(t *testing.T) {
  fixturePath, _ := filepath.Abs("fixtures/unsplash_response.json")
  fixtureBytes, _ := ioutil.ReadFile(fixturePath)

  var searchJsonResult = string(fixtureBytes)
  imagesCollectionExpectedSize := 5

  unsplashSearchResult := formatUnsplashSearchResult(searchJsonResult)
  realResultsLen := len(unsplashSearchResult.Results)

  if(realResultsLen < imagesCollectionExpectedSize) {
    t.Errorf("Images collection size is incorrect, got: %d, want: %d.", realResultsLen, imagesCollectionExpectedSize)
  }
}
