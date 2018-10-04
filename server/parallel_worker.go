package image_manager

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

type ImageCollectionResult struct {
  GiphyImageUrl []string
  UnsplashImageUrl []string
}


func ParallelGet(concurrencyLimit int) ImageCollectionResult {

}
