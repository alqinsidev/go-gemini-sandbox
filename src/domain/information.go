package domain

type Information struct {
	ID       int64
	Question string
	Answer   string
}

type InformationBodyRequest struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type InformationResponse struct {
	ID       int64  `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
