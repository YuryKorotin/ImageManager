package image_manager

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


