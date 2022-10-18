# sap-api-integrations-purchase-order-creates-rmq-kube
sap-api-integrations-purchase-order-creates-rmq-kube は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 購買発注データを登録するマイクロサービスです。    
sap-api-integrations-purchase-order-creates-rmq-kube には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-purchase-order-creates-rmq-kube は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_PURCHASEORDER_PROCESS_SRV_0001/overview   

## 動作環境  
sap-api-integrations-purchase-order-creates-rmq-kube は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-purchase-order-creates-rmq-kube は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchase-order-creates-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PURCHASEORDER_PROCESS_SRV_0001/overview    
* APIサービス名(=baseURL): API_PURCHASEORDER_PROCESS_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-purchase-order-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_PurchaseOrder（購買発注）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"HeaderItem" が指定されています。    
  
```
	"api_schema": "SAPPurchaseOrderCreates",
	"accepter": ["HeaderItem"],
	"purchase_order": "",
	"deleted": false
```
  
* 全データを登録する際のsample.jsonの記載例(2)  

全データを登録する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPPurchaseOrdercreates-rmq-kube",
	"accepter": ["All"],
	"purchase_order": "",
	"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostPurchaseOrder(
	headerItem *requests.HeaderItem,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "HeaderItem":
			func() {
				c.HeaderItem(headerItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 購買発注 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"PurchaseOrder" ～ "to_PurchaseOrderItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-purchase-order-creates-rmq-kube/SAP_API_Caller/caller.go#L59",
	"function": "sap-api-integrations-purchase-order-creates-rmq-kube/SAP_API_Caller.(*SAPAPICaller).HeaderItem",
	"level": "INFO",
	"message": {
		"PurchaseOrder": "4500000020",
		"CompanyCode": "0001",
		"PurchaseOrderType": "NB",
		"PurchasingProcessingStatus": "",
		"CreationDate": "2022-09-18",
		"Supplier": "100000",
		"Language": "ja",
		"PaymentTerms": "0001",
		"PurchasingOrganization": "0001",
		"PurchasingGroup": "001",
		"PurchaseOrderDate": "2022-09-22",
		"DocumentCurrency": "",
		"SupplierRespSalesPersonName": "",
		"SupplierPhoneNumber": "",
		"SupplyingPlant": "",
		"IncotermsClassification": "",
		"ManualSupplierAddressID": "",
		"AddressName": "",
		"AddressCityName": "",
		"AddressFaxNumber": "",
		"AddressPostalCode": "",
		"AddressStreetName": "",
		"AddressPhoneNumber": "",
		"AddressRegion": "",
		"to_PurchaseOrderItem": {
			"results": [
				{
					"PurchaseOrder": "",
					"PurchaseOrderItem": "10",
					"Plant": "0001",
					"StorageLocation": "0001",
					"MaterialGroup": "",
					"PurchasingInfoRecord": "",
					"SupplierMaterialNumber": "",
					"OrderQuantity": "1",
					"DocumentCurrency": "",
					"TaxCode": "V1",
					"UnlimitedOverdeliveryIsAllowed": null,
					"IsCompletelyDelivered": null,
					"IsFinallyInvoiced": null,
					"PurchaseOrderItemCategory": "",
					"AccountAssignmentCategory": "",
					"GoodsReceiptIsExpected": true,
					"GoodsReceiptIsNonValuated": false,
					"InvoiceIsExpected": true,
					"InvoiceIsGoodsReceiptBased": true,
					"Customer": "",
					"SupplierIsSubcontractor": null,
					"IncotermsClassification": "",
					"PurchaseRequisition": "",
					"PurchaseRequisitionItem": "",
					"RequisitionerName": "",
					"Material": "21",
					"InternationalArticleNumber": "",
					"PurchasingDocumentDeletionCode": "",
					"to_PurchaseOrderPricingElement": {
						"results": [
							{
								"PurchaseOrder": "",
								"PurchaseOrderItem": "10",
								"PricingDocument": "",
								"PricingDocumentItem": "10",
								"PricingProcedureStep": "1",
								"PricingProcedureCounter": "1",
								"ConditionType": "PBXX",
								"ConditionRateValue": "800",
								"ConditionCurrency": "",
								"PriceDetnExchangeRate": "",
								"TransactionCurrency": "",
								"ConditionAmount": null,
								"ConditionQuantityUnit": "",
								"ConditionQuantity": null,
								"ConditionApplication": "",
								"PricingDateTime": "",
								"ConditionCalculationType": "",
								"ConditionBaseValue": null,
								"ConditionToBaseQtyNmrtr": null,
								"ConditionToBaseQtyDnmntr": null,
								"ConditionCategory": "",
								"PricingScaleType": "",
								"ConditionOrigin": "",
								"IsGroupCondition": "",
								"ConditionSequentialNumber": "",
								"ConditionInactiveReason": "",
								"PricingScaleBasis": "",
								"ConditionScaleBasisValue": null,
								"ConditionScaleBasisCurrency": "",
								"ConditionScaleBasisUnit": "",
								"ConditionIsManuallyChanged": null,
								"ConditionRecord": ""
							}
						]
					}
				}
			]
		}
	},
	"time": "2022-09-28T19:42:59+09:00"
}

```