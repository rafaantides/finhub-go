package upload

import (
	"encoding/csv"
	"finhub-go/internal/core/ports/inbound"
	"finhub-go/internal/core/ports/outbound/messagebus"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/xuri/excelize/v2"
)

type uploadService struct {
	mb messagebus.MessageBus
}

func NewUploadService(mb messagebus.MessageBus) inbound.UploadService {
	return &uploadService{mb: mb}
}

func (c *uploadService) ImportFile(resource, model, action string, file multipart.File, fileHeader *multipart.FileHeader) error {
	fileType, err := detectFileType(file)
	if err != nil {
		return err
	}

	filename := fileHeader.Filename

	var rows [][]string

	switch fileType {
	case "csv":
		rows, err = c.readCSV(file)
	case "xlsx":
		rows, err = c.readXLSX(file)
	default:
		return fmt.Errorf("unsupported file type")
	}

	if err != nil {
		return err
	}

	return c.readRows(resource, model, action, filename, rows)
}

func (c *uploadService) readCSV(file multipart.File) ([][]string, error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func (c *uploadService) readXLSX(file multipart.File) ([][]string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}

	sheetName := f.GetSheetName(0)
	return f.GetRows(sheetName)
}

func (c *uploadService) readRows(resource, model, action, filename string, rows [][]string) error {
	if len(rows) < 2 {
		return fmt.Errorf("invalid file: no data found")
	}

	columnIndex := indexColumns(rows[0])

	switch resource {
	case "debt":
		c.processDebts(model, action, filename, rows, columnIndex)
	default:
		return fmt.Errorf("unknown resource: %s", resource)
	}

	return nil

}

func detectFileType(file multipart.File) (string, error) {
	buffer := make([]byte, 4)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}
	file.Seek(0, io.SeekStart)

	if buffer[0] == 0x50 && buffer[1] == 0x4B {
		return "xlsx", nil
	}
	return "csv", nil
}

func indexColumns(headers []string) map[string]int {
	index := make(map[string]int)
	for i, header := range headers {
		index[strings.ToLower(header)] = i
	}
	return index
}
