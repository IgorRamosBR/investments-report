package reader

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type XlsxReader interface {
	ReadFile() ([][]string, error)
}

type XlsxReaderImpl struct {
	Path string
}

func NewXlsxReader(path string) XlsxReader {
	return XlsxReaderImpl{
		Path: path,
	}
}

func (r XlsxReaderImpl) ReadFile() ([][]string, error) {
	f, err := excelize.OpenFile("./kinvo__resumo-produtos (1).xlsx")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Resumo de produtos")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	return rows, nil
}
