package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Ward represents a Ward from the database
type Ward struct {
	ObjectID       int     `json:"objectid"`
	Country        string  `json:"country"`
	ISO3           string  `json:"iso3"`
	State          string  `json:"state"`
	StateCode      string  `json:"statecode"`
	LGA            string  `json:"lga"`
	LGAAltNames    string  `json:"lga_alt_names"`
	Ward           string  `json:"ward"`
	WardAltNames   string  `json:"ward_alt_names"`
	MultipartCount int     `json:"multipart_count"`
	Source         string  `json:"source"`
	Date           string  `json:"date"`
	AreaSqkm       int     `json:"area_sqkm"`
}

var db *sql.DB

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get database configuration
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "grid3_wards")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	// Get server configuration
	serverPort := getEnv("SERVER_PORT", "8080")
	ginMode := getEnv("GIN_MODE", "debug")

	// Set Gin mode
	gin.SetMode(ginMode)

	// Connect to PostgreSQL
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		// Wards endpoints
		v1.GET("/wards", getWards)
		v1.GET("/wards/state/:state", getWardsByState)
		v1.GET("/wards/lga/:lga", getWardsByLGA)
	}

	log.Printf("Starting server on port %s", serverPort)
	r.Run(":" + serverPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Handlers

func getWards(c *gin.Context) {
	rows, err := db.Query("SELECT \"OBJECTID\", country, iso3, state, statecode, lga, lga_alt_names, ward, ward_alt_names, multipart_count, source, date, area_sqkm FROM wards")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var wards []Ward
	for rows.Next() {
		var w Ward
		err := rows.Scan(&w.ObjectID, &w.Country, &w.ISO3, &w.State, &w.StateCode, &w.LGA, &w.LGAAltNames, &w.Ward, &w.WardAltNames, &w.MultipartCount, &w.Source, &w.Date, &w.AreaSqkm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		wards = append(wards, w)
	}

	c.JSON(http.StatusOK, wards)
}

func getWardsByState(c *gin.Context) {
	state := c.Param("state")
	rows, err := db.Query("SELECT \"OBJECTID\", country, iso3, state, statecode, lga, lga_alt_names, ward, ward_alt_names, multipart_count, source, date, area_sqkm FROM wards WHERE state ILIKE $1", "%"+state+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var wards []Ward
	for rows.Next() {
		var w Ward
		err := rows.Scan(&w.ObjectID, &w.Country, &w.ISO3, &w.State, &w.StateCode, &w.LGA, &w.LGAAltNames, &w.Ward, &w.WardAltNames, &w.MultipartCount, &w.Source, &w.Date, &w.AreaSqkm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		wards = append(wards, w)
	}

	c.JSON(http.StatusOK, wards)
}

func getWardsByLGA(c *gin.Context) {
	lga := c.Param("lga")
	rows, err := db.Query("SELECT \"OBJECTID\", country, iso3, state, statecode, lga, lga_alt_names, ward, ward_alt_names, multipart_count, source, date, area_sqkm FROM wards WHERE lga ILIKE $1", "%"+lga+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var wards []Ward
	for rows.Next() {
		var w Ward
		err := rows.Scan(&w.ObjectID, &w.Country, &w.ISO3, &w.State, &w.StateCode, &w.LGA, &w.LGAAltNames, &w.Ward, &w.WardAltNames, &w.MultipartCount, &w.Source, &w.Date, &w.AreaSqkm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		wards = append(wards, w)
	}

	c.JSON(http.StatusOK, wards)
}
