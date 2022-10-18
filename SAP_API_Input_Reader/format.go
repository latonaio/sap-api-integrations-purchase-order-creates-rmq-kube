package sap_api_input_reader

import (
	"sap-api-integrations-purchase-order-creates-rmq-kube/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.PurchaseOrder
	results := make([]requests.Item, 0, len(data.PurchaseOrderItem))

	header := sdc.ConvertToHeader()

	for i := range data.PurchaseOrderItem {
		results = append(results, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.PurchaseOrder
	return &requests.Header{
		PurchaseOrder:               data.PurchaseOrder,
		CompanyCode:                 data.CompanyCode,
		PurchaseOrderType:           data.PurchaseOrderType,
		PurchasingProcessingStatus:  data.PurchasingProcessingStatus,
		CreationDate:                data.CreationDate,
		Supplier:                    data.Supplier,
		Language:                    data.Language,
		PaymentTerms:                data.PaymentTerms,
		PurchasingOrganization:      data.PurchasingOrganization,
		PurchasingGroup:             data.PurchasingGroup,
		PurchaseOrderDate:           data.PurchaseOrderDate,
		DocumentCurrency:            data.DocumentCurrency,
		SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:         data.SupplierPhoneNumber,
		SupplyingPlant:              data.SupplyingPlant,
		IncotermsClassification:     data.IncotermsClassification,
		ManualSupplierAddressID:     data.ManualSupplierAddressID,
		AddressName:                 data.AddressName,
		AddressCityName:             data.AddressCityName,
		AddressFaxNumber:            data.AddressFaxNumber,
		AddressPostalCode:           data.AddressPostalCode,
		AddressStreetName:           data.AddressStreetName,
		AddressPhoneNumber:          data.AddressPhoneNumber,
		AddressRegion:               data.AddressRegion,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataPurchaseOrder := sdc.PurchaseOrder
	data := sdc.PurchaseOrder.PurchaseOrderItem[num]
	return &requests.Item{
		PurchaseOrder:                  dataPurchaseOrder.PurchaseOrder,
		PurchaseOrderItem:              data.PurchaseOrderItem,
		Plant:                          data.Plant,
		StorageLocation:                data.StorageLocation,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		SupplierMaterialNumber:         data.SupplierMaterialNumber,
		OrderQuantity:                  data.OrderQuantity,
		DocumentCurrency:               data.DocumentCurrency,
		TaxCode:                        data.TaxCode,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		IsFinallyInvoiced:              data.IsFinallyInvoiced,
		PurchaseOrderItemCategory:      data.PurchaseOrderItemCategory,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		InvoiceIsExpected:              data.InvoiceIsExpected,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		Customer:                       data.Customer,
		SupplierIsSubcontractor:        data.SupplierIsSubcontractor,
		IncotermsClassification:        data.IncotermsClassification,
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		RequisitionerName:              data.RequisitionerName,
		Material:                       data.Material,
		InternationalArticleNumber:     data.InternationalArticleNumber,
		PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
	}
}
