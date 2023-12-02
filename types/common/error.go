package common

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Context interface{} `json:"context"`
}
