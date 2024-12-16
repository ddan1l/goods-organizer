package models

type ProductModel struct {
	Title          string `json:"title"`
	Article        string `json:"article"`
	Code           string `json:"code"`
	PurchasesPrice int32  `json:"purchases_price"`
	MinPrice       int32  `json:"min_price"`
	SellingPrice   int32  `json:"selling_price"`
	CountryId      int32  `json:"country_id"`
	VatId          int32  `json:"vat_id"`
	BarcodeId      int32  `json:"barcode_id"`
}
