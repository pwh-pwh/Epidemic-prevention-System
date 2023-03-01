package dto

type GoodsDto struct {
	Id       int64  `json:"id"`
	GoodName string `json:"goodName"`
	GoodNum  int32  `json:"goodNum"`
	GoodSize string `json:"goodSize"`
}
