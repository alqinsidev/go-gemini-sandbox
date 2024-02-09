package domain

type ChatRequestPayload struct {
	Question string `json:"question" binding:"required"`
}

type ChatResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
