package main

import (
	"go-server-curriculum/handler"
	"go-server-curriculum/infrastructure"
	"go-server-curriculum/repository"
	"go-server-curriculum/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// DB 初期化
	db, err := infrastructure.NewMySQLDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// // リポジトリ初期化
	productRepo := repository.NewProductRepository(db) // 追加
	orderRepo := repository.NewOrderRepository(db)
	customerRepo := repository.NewCustomerRepository(db)

	// // ユースケース初期化
	productUsecase := usecase.NewProductUsecase(productRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo, productRepo) // 修正

	// // ハンドラー初期化
	healthHandler := handler.NewHealthHandler()
	productHandler := handler.NewProductHandler(productUsecase)
	orderHandler := handler.NewOrderHandler(orderUsecase)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	// Echoルーター設定
	e := echo.New()
	
	// CORS設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	
	e.GET("/", healthHandler.HealthCheck)
	
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProduct)
	e.POST("/products", productHandler.CreateProduct)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)

	e.GET("/orders", orderHandler.GetOrders)
	e.GET("/orders/:id", orderHandler.GetOrder)
	e.POST("/orders", orderHandler.CreateOrder)
	e.PUT("/orders/:id", orderHandler.UpdateOrder)
	e.DELETE("/orders/:id", orderHandler.DeleteOrder)

	e.GET("/customers", customerHandler.GetCustomers)
	e.GET("/customers/:id", customerHandler.GetCustomer)
	e.POST("/customers", customerHandler.CreateCustomer)
	e.PUT("/customers/:id", customerHandler.UpdateCustomer)
	e.DELETE("/customers/:id", customerHandler.DeleteCustomer)

	e.GET("/customers/:id/total", customerHandler.GetCustomerTotal)

	// サーバー起動
	log.Println("Server running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
