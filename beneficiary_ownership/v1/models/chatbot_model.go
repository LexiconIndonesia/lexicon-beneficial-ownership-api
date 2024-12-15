package bo_v1_models

type ChatbotRequest struct {
	ThreadID    string `json:"thread_id"`
	UserMessage string `json:"user_message"`
}

type ChatbotResponse struct {
	Response   string   `json:"response"`
	References []string `json:"references"`
}

type ChatbotReferenceRequest struct {
	CaseNumbers []string `json:"references"`
}
