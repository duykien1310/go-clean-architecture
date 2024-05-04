package presenter

type BasicResponse struct {
	Status  string      `json:"status" example:"200"`
	Message string      `json:"message" example:"SUCCESS"`
	Results interface{} `json:"results"`
}

type PaginationResponse struct {
	Status     string      `json:"status" example:"200"`
	Message    string      `json:"message" example:"SUCCESS"`
	Results    interface{} `json:"results"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Count         int `json:"count" example:"24"`
	NumPages      int `json:"numPages" example:"3"`
	DisplayRecord int `json:"displayRecord" example:"10"`
	Page          int `json:"page" example:"1"`
}

type Error400Response struct {
	Status  string      `json:"status" example:"400"`
	Message string      `json:"message" example:"Error Message"`
	Results interface{} `json:"results"`
}
type Error404Response struct {
	Status  string      `json:"status" example:"500"`
	Message string      `json:"message" example:"Error Message"`
	Results interface{} `json:"results"`
}
type Error500Response struct {
	Status  string      `json:"status" example:"500"`
	Message string      `json:"message" example:"Error Message"`
	Results interface{} `json:"results"`
}
