package dto

type UrlEncurtadaForm struct {
	URL string `json:"url"`
}

type UrlEncurtadaCriadaResponse struct {
	ID     uint   `json:"id"`
	Codigo string `json:"codigo"`
}

type UrlEncurtadaResponse struct {
	URL string `json:"url"`
}
