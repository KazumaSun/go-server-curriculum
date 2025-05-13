package domain

type Customer struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Name     string `json:"name"`
    Seat     string `json:"seat"`
    Orders   []Order `json:"orders" gorm:"foreignKey:CustomerID"`
}