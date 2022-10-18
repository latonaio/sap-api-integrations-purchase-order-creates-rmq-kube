package requests

type HeaderItem struct {
	Header
	ToItem `json:"to_PurchaseOrderItem"`
}

type ToItem struct {
	Results []Item `json:"results"`
}
