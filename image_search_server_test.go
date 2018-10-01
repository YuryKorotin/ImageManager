package main

import (
	"testing"
	//"fmt"
)

func TestGetTranslations(t *testing.T) {
  var body = "{\"code\":102, \"lang\": \"ru-en\", \"text\": [\"translation\"] }"
  translation, _ := getTranslations([]byte(body))
  if (translation.Texts[0] != "translation") {
    t.Errorf("Translation was incorrect, got: %s, want: %s.", translation.Texts[0], "translation")
  }
}

func TestFormatUnsplashSearhResult(t *testing.T) {
  var searchJsonResult = "{}"
  var imagesCollectionExpectedSize = 5

  unsplashSearchResult := formatUnsplashSearchResult(searchJsonResult)
  realResultsLen := len(unsplashSearchResult.Results)

  if(realResultsLen < imagesCollectionExpectedSize) {
    t.Error("Images collection  size is incorrect, got %i, want %i.", realResultsLen, imagesCollectionExpectedSize)
  }
}
