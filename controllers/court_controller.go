package controllers

import (
	"net/http"

	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/models"
	"github.com/gin-gonic/gin"
)

// GetCourts godoc
// @Summary      Get all courts
// @Description  Menampilkan semua lapangan futsal
// @Tags         Courts
// @Produce      json
// @Success      200  {array}  models.Court
// @Router       /api/courts [get]
func GetCourts(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, location, price_per_hour, is_available FROM courts ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var courts []models.Court
	for rows.Next() {
		var court models.Court
		if err := rows.Scan(&court.ID, &court.Name, &court.Location, &court.PricePerHour, &court.IsAvailable); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		courts = append(courts, court)
	}

	c.JSON(http.StatusOK, courts)
}

// GetCourtByID godoc
// @Summary      Get court by ID
// @Description  Menampilkan detail lapangan futsal berdasarkan ID
// @Tags         Courts
// @Produce      json
// @Param        id   path      int  true  "Court ID"
// @Success      200  {object}  models.Court
// @Failure      404  {object}  map[string]string
// @Router       /api/courts/{id} [get]
func GetCourtByID(c *gin.Context) {
	id := c.Param("id")
	var court models.Court

	err := config.DB.QueryRow("SELECT id, name, location, price_per_hour, is_available FROM courts WHERE id = $1", id).
		Scan(&court.ID, &court.Name, &court.Location, &court.PricePerHour, &court.IsAvailable)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Court not found"})
		return
	}

	c.JSON(http.StatusOK, court)
}

// CreateCourt godoc
// @Summary      Tambah lapangan baru
// @Description  Menambahkan data lapangan futsal baru
// @Tags         Courts
// @Accept       json
// @Produce      json
// @Param        court  body  models.Court  true  "Court Data"
// @Success      201  {object}  models.Court
// @Router       /api/courts [post]
func CreateCourt(c *gin.Context) {
	var court models.Court
	if err := c.ShouldBindJSON(&court); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.QueryRow(
		"INSERT INTO courts (name, location, price_per_hour, is_available) VALUES ($1, $2, $3, $4) RETURNING id",
		court.Name, court.Location, court.PricePerHour, court.IsAvailable,
	).Scan(&court.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, court)
}

// UpdateCourt godoc
// @Summary      Update court
// @Description  Memperbarui data lapangan futsal berdasarkan ID
// @Tags         Courts
// @Accept       json
// @Produce      json
// @Param        id     path      int          true  "Court ID"
// @Param        court  body      models.Court true  "Court Data"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/courts/{id} [put]
func UpdateCourt(c *gin.Context) {
	id := c.Param("id")
	var court models.Court

	if err := c.ShouldBindJSON(&court); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"UPDATE courts SET name=$1, location=$2, price_per_hour=$3, is_available=$4 WHERE id=$5",
		court.Name, court.Location, court.PricePerHour, court.IsAvailable, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Court updated successfully"})
}

// DeleteCourt godoc
// @Summary      Delete court
// @Description  Menghapus data lapangan futsal berdasarkan ID
// @Tags         Courts
// @Produce      json
// @Param        id   path      int  true  "Court ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/courts/{id} [delete]
func DeleteCourt(c *gin.Context) {
	id := c.Param("id")
	_, err := config.DB.Exec("DELETE FROM courts WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Court deleted successfully"})
}
