package response

type ShortenURLResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url,omitempty"`
}

type GetURLResponse struct {
	OriginalURL string `json:"original_url"`
	ClickCount  int64  `json:"click_count"`
}
