package controllers

import (
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/models"
)

// GET /bookings
func GetBookings(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, court_id, customer_name, booking_date, start_time, end_time, total_price 
		FROM bookings ORDER BY booking_date DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.CourtID, &b.CustomerName, &b.BookingDate, &b.StartTime, &b.EndTime, &b.TotalPrice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		bookings = append(bookings, b)
	}

	c.JSON(http.StatusOK, bookings)
}

// GET /bookings/:id
func GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Booking

	query := `
		SELECT id, court_id, customer_name, booking_date, start_time, end_time, total_price
		FROM bookings WHERE id = $1
	`
	err := config.DB.QueryRow(query, id).Scan(
		&b.ID, &b.CourtID, &b.CustomerName, &b.BookingDate, &b.StartTime, &b.EndTime, &b.TotalPrice,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// POST /bookings
func CreateBooking(c *gin.Context) {
	var newBooking models.Booking
	if err := c.ShouldBindJSON(&newBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO bookings (court_id, customer_name, booking_date, start_time, end_time, total_price)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := config.DB.QueryRow(query,
		newBooking.CourtID,
		newBooking.CustomerName,
		newBooking.BookingDate,
		newBooking.StartTime,
		newBooking.EndTime,
		newBooking.TotalPrice,
	).Scan(&newBooking.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBooking)
}

// PUT /bookings/:id
func UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var updated models.Booking

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE bookings 
		SET court_id=$1, customer_name=$2, booking_date=$3, start_time=$4, end_time=$5, total_price=$6
		WHERE id=$7
	`
	res, err := config.DB.Exec(query,
		updated.CourtID,
		updated.CustomerName,
		updated.BookingDate,
		updated.StartTime,
		updated.EndTime,
		updated.TotalPrice,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully"})
}

// DELETE /bookings/:id
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM bookings WHERE id = $1`
	res, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
