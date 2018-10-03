package image_manager

type TranslationResponse struct {
  Code int `json:"code"`
  Lang string `json:"lang"`
  Texts []string `json:"text"`
}
