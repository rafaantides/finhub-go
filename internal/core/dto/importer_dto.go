package dto

type ImportDebtMessage struct {
	JobID        string `json:"job_id"`
	Filename     string `json:"filename"`
	IsFirstChunk bool   `json:"is_first_chunk"`
	IsLastChunk  bool   `json:"is_last_chunk"`
	Action       string `json:"action"`
	Data         struct {
		Invoice InvoiceRequest `json:"invoice"`
		Debt    DebtRequest    `json:"debt"`
	} `json:"data"`
}
