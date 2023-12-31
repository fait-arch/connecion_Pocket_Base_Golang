/* CreateProduct.go */
package CreateProduct

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Definimos la estructura JSON de Collections "products"
type Product struct {
	ProductName  string `json:"product_Name"`
	ProductPrice string `json:"product_Price"`
}

func CreateProduct(product *Product) (*http.Response, error) {
	url := "http://localhost:8090/api/collections/products/records"

	jsonData, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
