package model

type PageRequest struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Order     string `json:"count"`
	Direction string `json:"direction"`
}

type PageResponse struct {
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Count int           `json:"count"`
	Items []interface{} `json:"items"`
}
