package vo

type PageVO struct {
	Records any   `json:"records"`
	Total   int64 `json:"total"`
}
