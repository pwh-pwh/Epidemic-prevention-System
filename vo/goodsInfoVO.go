package vo

type GoodsInfoVO struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Size  string `json:"size"`
	Total int32  `json:"total"`
}
