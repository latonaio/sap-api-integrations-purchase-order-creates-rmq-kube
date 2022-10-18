package responses

type Header struct {
	D struct {
		PurchaseOrder string `json:"PurchaseOrder"`
	} `json:"d"`
}
