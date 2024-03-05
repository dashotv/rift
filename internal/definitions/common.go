package definitions

type Request struct {
	ID    string `json:"id" query:"id"`
	Limit int    `json:"limit" query:"limit"`
	Skip  int    `json:"skip" query:"skip"`
}

type Response struct {
	Total   int64       `json:"total,omitempty"`
	Results interface{} `json:"results,omitempty"`
}
