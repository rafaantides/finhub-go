package dto

type NotifierMessage struct {
	JobID       string
	Filename    string
	IsLastChunk bool
}
