package main

import "testing"

func TestSearchImage(t *testing.T) {
	var body = "{\"code\":102, \"lang\": \"ru-en\", \"text\": [\"translation\"] }"
	translation, _ := getTranslations([]byte(body))
	if (translation.Texts[0] != "translation") {
       		t.Errorf("Translation was incorrect, got: %s, want: %s.", translation.Texts[0], "translation")
    	}
}
