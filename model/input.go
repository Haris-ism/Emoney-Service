package model

type Sign struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type TopUp struct {
	TopUp int `json:"top_up"`
}
type Confirm struct {
	InquiryID int `json:"inquiry_id"`
}
type Requests struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    []Child `json:"data"`
}
type Requests1 struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Child  `json:"data"`
}
type Child struct {
	ID          int
	Category    string
	Product     string
	Description string
	Price       int
	Fee         int
}