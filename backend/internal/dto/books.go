package dto

import "libraryManagement/internal/models"

type ResponseBookInfo struct {
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Authors   string `json:"authors"`
	Publisher string `json:"publisher"`
	Version   string `json:"version"`
}

type RequestUpdateBook struct {
	Title           string `json:"title" validate:"alpha_space"`
	Isbn            string `json:"isbn" validate:"isbn" `
	Authors         string `json:"authors" validate:"alpha_space"`
	Publisher       string `json:"publisher" validate:"alpha_space"`
	Version         string `json:"version"`
	TotalCopies     uint   `json:"totalCopies"`
	AvailableCopies uint   `json:"availableCopies"`
}

type SearchBookPayload struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	ISBN      string `json:"isbn"`
}

type IssuedBooks struct {
	models.IssueRegistery
	Book struct {
		Name string
	}
}
