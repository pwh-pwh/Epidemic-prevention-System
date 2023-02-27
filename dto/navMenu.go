package dto

type NavMenu struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Title     string     `json:"title"`
	Icon      string     `json:"icon"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	OrderNum  int        `json:"orderNum"`
	Children  []*NavMenu `json:"children"`
}
