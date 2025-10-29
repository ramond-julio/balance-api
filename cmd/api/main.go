package main

import (
	"balance-api/config"
	"balance-api/internal/handler"
	"balance-api/internal/repository/mysql"
	"balance-api/internal/service"
	"balance-api/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.NewMySQLConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	balanceRepo := mysql.NewBalanceRepository(db)

	// Initialize service
	balanceService := service.NewBalanceService(balanceRepo)

	// Initialize handler
	balanceHandler := handler.NewBalanceHandler(balanceService)

	// Setup routes
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/balance/{user_id}", balanceHandler.GetBalance).Methods("GET")
	router.HandleFunc("/api/v1/withdraw", balanceHandler.Withdraw).Methods("POST")

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}