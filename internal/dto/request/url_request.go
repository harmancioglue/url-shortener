package request

type ShortenURLRequest struct {
	URL string `json:"url" validate:"required,url"`
}
