package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"tasm-backend/database"
	"tasm-backend/models"
)

func GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not connected"})
		return
	}

	database.DB.Order("id desc").Find(&reservations)

	if len(reservations) == 0 {
		now := time.Now()
		reservations = []models.Reservation{
			{Title: "Nila 102", Type: "Meeting Room", Status: "Active", StartTime: now.Add(-time.Hour), EndTime: now.Add(time.Hour), Location: "Nila Building", BookedBy: "Alice Johnson"},
			{Title: "Epson Projector Pro", Type: "AV Equipment", Status: "Upcoming", StartTime: now.Add(2 * time.Hour), EndTime: now.Add(4 * time.Hour), Location: "IT Store", BookedBy: "Bob Smith"},
			{Title: "Corporate Van", Type: "Vehicle", Status: "Upcoming", StartTime: now.AddDate(0, 0, 1), EndTime: now.AddDate(0, 0, 2), Location: "Parking Basement", BookedBy: "Charlie Davis"},
		}
		for _, r := range reservations {
			database.DB.Create(&r)
		}
		database.DB.Order("id desc").Find(&reservations)
	}

	c.JSON(http.StatusOK, reservations)
}
