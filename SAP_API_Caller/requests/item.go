package requests

type Item struct {
	PurchaseOrder                  string  `json:"PurchaseOrder"`
	PurchaseOrderItem              string  `json:"PurchaseOrderItem"`
	Plant                          *string `json:"Plant"`
	StorageLocation                *string `json:"StorageLocation"`
	MaterialGroup                  *string `json:"MaterialGroup"`
	PurchasingInfoRecord           *string `json:"PurchasingInfoRecord"`
	SupplierMaterialNumber         *string `json:"SupplierMaterialNumber"`
	OrderQuantity                  *string `json:"OrderQuantity"`
	DocumentCurrency               *string `json:"DocumentCurrency"`
	TaxCode                        *string `json:"TaxCode"`
	UnlimitedOverdeliveryIsAllowed *bool   `json:"UnlimitedOverdeliveryIsAllowed"`
	IsCompletelyDelivered          *bool   `json:"IsCompletelyDelivered"`
	IsFinallyInvoiced              *bool   `json:"IsFinallyInvoiced"`
	PurchaseOrderItemCategory      *string `json:"PurchaseOrderItemCategory"`
	AccountAssignmentCategory      *string `json:"AccountAssignmentCategory"`
	GoodsReceiptIsExpected         *bool   `json:"GoodsReceiptIsExpected"`
	GoodsReceiptIsNonValuated      *bool   `json:"GoodsReceiptIsNonValuated"`
	InvoiceIsExpected              *bool   `json:"InvoiceIsExpected"`
	InvoiceIsGoodsReceiptBased     *bool   `json:"InvoiceIsGoodsReceiptBased"`
	Customer                       *string `json:"Customer"`
	SupplierIsSubcontractor        *bool   `json:"SupplierIsSubcontractor"`
	IncotermsClassification        *string `json:"IncotermsClassification"`
	PurchaseRequisition            *string `json:"PurchaseRequisition"`
	PurchaseRequisitionItem        *string `json:"PurchaseRequisitionItem"`
	RequisitionerName              *string `json:"RequisitionerName"`
	Material                       *string `json:"Material"`
	InternationalArticleNumber     *string `json:"InternationalArticleNumber"`
	PurchasingDocumentDeletionCode *string `json:"PurchasingDocumentDeletionCode"`
	//	ToItemPricingElement *struct {
	//		ToItemPricingElementResults []*ItemPricingElement `json:"results"`
	//	} `json:"to_PricingElement"`
}
