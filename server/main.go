package main

import (
	"go-server-curriculum/handler"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// DB 初期化
	// db, err := infrastructure.NewMySQLDB()
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	// // リポジトリ初期化
	// productRepo := repository.NewProductRepository(db)
	// orderRepo := repository.NewOrderRepository(db)

	// // ユースケース初期化
	// productUsecase := usecase.NewProductUsecase(productRepo)
	// orderUsecase := usecase.NewOrderUsecase(orderRepo)

	// // ハンドラー初期化
	 healthHandler := handler.NewHealthHandler()
	// productHandler := handler.NewProductHandler(productUsecase)
	// orderHandler := handler.NewOrderHandler(orderUsecase)

	// Echoルーター設定
	e := echo.New()
	e.GET("/", healthHandler.HealthCheck)
	// e.GET("/products", productHandler.GetProducts)
	// e.GET("/products/:id", productHandler.GetProduct)
	// e.GET("/orders", orderHandler.GetOrders)

	// サーバー起動
	log.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
