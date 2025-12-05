package controllers

import (
	"net/http"
	// "strconv"

	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/models"
	"github.com/gin-gonic/gin"
)

// GET /bookings
// GetBookings godoc
// @Summary      Get all bookings
// @Description  Menampilkan semua data booking
// @Tags         Bookings
// @Produce      json
// @Success      200  {array}  models.Booking
// @Router       /api/bookings [get]
func GetBookings(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, court_id, user_id, customer_name, booking_date, start_time, end_time, total_price 
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
		if err := rows.Scan(&b.ID, &b.CourtID, &b.UserID, &b.CustomerName, &b.BookingDate, &b.StartTime, &b.EndTime, &b.TotalPrice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		bookings = append(bookings, b)
	}

	c.JSON(http.StatusOK, bookings)
}

// GET /bookings/:id
// GetBookingByID godoc
// @Summary      Get booking by ID
// @Description  Menampilkan detail booking berdasarkan ID
// @Tags         Bookings
// @Produce      json
// @Param        id   path      int  true  "Booking ID"
// @Success      200  {object}  models.Booking
// @Failure      404  {object}  map[string]string
// @Router       /api/bookings/{id} [get]
func GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Booking

	query := `
		SELECT id, court_id, user_id, customer_name, booking_date, start_time, end_time, total_price
		FROM bookings WHERE id = $1
	`
	err := config.DB.QueryRow(query, id).Scan(
		&b.ID, &b.CourtID, &b.UserID, &b.CustomerName, &b.BookingDate, &b.StartTime, &b.EndTime, &b.TotalPrice,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// POST /bookings
// CreateBooking godoc
// @Summary      Create new booking
// @Description  Membuat data booking baru
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Param        booking  body  models.Booking  true  "Booking Data"
// @Success      201  {object}  models.Booking
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/bookings [post]
func CreateBooking(c *gin.Context) {
	var newBooking models.Booking
	if err := c.ShouldBindJSON(&newBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO bookings (court_id, user_id, customer_name, booking_date, start_time, end_time, total_price)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	err := config.DB.QueryRow(query,
		newBooking.CourtID,
		newBooking.UserID,
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
// UpdateBooking godoc
// @Summary      Update booking
// @Description  Memperbarui data booking berdasarkan ID
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Param        id      path      int          true  "Booking ID"
// @Param        booking body      models.Booking true  "Booking Data"
// @Success      200     {object}  map[string]string
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/bookings/{id} [put]
func UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var updated models.Booking

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE bookings 
		SET court_id=$1, user_id=$2, customer_name=$3, booking_date=$4, start_time=$5, end_time=$6, total_price=$7
		WHERE id=$8
	`
	res, err := config.DB.Exec(query,
		updated.CourtID,
		updated.UserID,
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
// DeleteBooking godoc
// @Summary      Delete booking
// @Description  Menghapus data booking berdasarkan ID
// @Tags         Bookings
// @Produce      json
// @Param        id   path      int  true  "Booking ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/bookings/{id} [delete]
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
