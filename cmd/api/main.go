package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	
	"kasir-api/internal/database"
	"kasir-api/internal/services"
	"kasir-api/internal/handlers"
	"kasir-api/internal/repositories"

)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	viper.AutomaticEnv()

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	fmt.Println("DB_CONN =", config.DBConn)

	//handling confignya
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	//report
	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	//routing kasir
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST
	http.HandleFunc("/api/report", reportHandler.GetTodayReport)

	//tes localhost dengan port dari .env
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//panggil
		json.NewEncoder(w).Encode(map[string]string{
			"Status":  "OK",
			"Message": "API RUNNING",
		})
	})
	fmt.Println("Server Localhost Running di: " + config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}