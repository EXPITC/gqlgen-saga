// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type MarkTodo struct {
	ID   uint `json:"id"`
	Done bool `json:"done"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID uint   `json:"userId"`
}

type PaginationRequest struct {
	Batch     int `json:"batch"`
	BatchSize int `json:"batchSize"`
}
