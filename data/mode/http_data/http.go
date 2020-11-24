package http_data

import "fmt"

type DataHttp struct {
	URL string
}

// NewDataHttp
func NewDataHttp(url string) *DataHttp {
	return &DataHttp{
		URL: url,
	}
}

func (h *DataHttp) GetData() (ret interface{}, err error) {
	fmt.Println("wwwwwwwwwwwwwwwwww")
	return
}
