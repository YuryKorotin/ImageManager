package image_manager

type GiphyItem struct {
  Url string `json:"url"`
}

type GiphyResponse struct {
  Data []GiphyItem `json:"data"`
}
