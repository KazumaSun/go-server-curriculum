package domain

type Order struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
    ProductID  uint   `json:"product_id"`
    CustomerID uint   `json:"customer_id"`
    Quantity   int    `json:"quantity"`
}
