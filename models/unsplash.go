package image_manager

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


