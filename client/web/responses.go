package web

type Responses struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
