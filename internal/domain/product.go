package domain

type ProductResponse struct {
	Message string  `json:"message"`
	Product Product `json:"product"`
}

type Product struct {
	ID    uint    `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
	Stock int     `json:"stock,omitempty"`
}
