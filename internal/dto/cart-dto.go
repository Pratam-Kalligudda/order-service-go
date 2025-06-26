package dto

type AddUpdateProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
