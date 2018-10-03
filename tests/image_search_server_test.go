package image_manager

import (
  "testing"
  "io/ioutil"
  "path/filepath"
)
import server "../server"

func TestGetTranslations(t *testing.T) {
  var body = "{\"code\":102, \"lang\": \"ru-en\", \"text\": [\"translation\"] }"
  translation, _ := server.GetTranslations([]byte(body))
  if (translation.Texts[0] != "translation") {
    t.Errorf("Translation was incorrect, got: %s, want: %s.", translation.Texts[0], "translation")
  }
}

func TestFormatUnsplashSearhResult(t *testing.T) {
  fixturePath, _ := filepath.Abs("../fixtures/unsplash_response.json")
  fixtureBytes, _ := ioutil.ReadFile(fixturePath)

  var searchJsonResult = string(fixtureBytes)
  imagesCollectionExpectedSize := 5

  unsplashSearchResult := server.FormatUnsplashSearchResult(searchJsonResult)
  realResultsLen := len(unsplashSearchResult.Results)

  if(realResultsLen < imagesCollectionExpectedSize) {
    t.Errorf("Images collection size is incorrect, got: %d, want: %d.", realResultsLen, imagesCollectionExpectedSize)
  }
}
