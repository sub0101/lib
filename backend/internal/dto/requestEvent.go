package dto

type RequestIssueStatus struct {
	Type string `json:"requestType"`
}

type RequestEventDTO struct {
	BookID      uint   `json:"bookId" binding:"required"`
	RequestType string `json:"requestType" binding:"required"`
	ReaderID    uint   `json:"readerId"`
}

type RequestUserEvent struct {
	ReaderID uint
	BookID   uint
	Book     struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		ISBN  string `json:"isbn"`
	}
}
