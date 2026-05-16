package handlers

import (
	"net/http"
	"strings"

	"tasm-backend/models"

	"github.com/gin-gonic/gin"
)

func GetReservations(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var reservations []models.Reservation
	query := db.Order("start_time desc")
	if status := strings.TrimSpace(c.Query("status")); status != "" {
		query = query.Where("status = ?", status)
	}
	if resourceType := strings.TrimSpace(c.Query("type")); resourceType != "" {
		query = query.Where("type = ?", resourceType)
	}
	if bookedBy := strings.TrimSpace(c.Query("bookedBy")); bookedBy != "" {
		query = query.Where("booked_by = ?", bookedBy)
	}
	if err := query.Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load reservations"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func GetReservationByID(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var reservation models.Reservation
	if err := db.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

type reservationRequest struct {
	Title     string `json:"title"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Location  string `json:"location"`
	BookedBy  string `json:"bookedBy"`
}

func CreateReservation(c *gin.Context) {
	var payload reservationRequest
	if !bindJSON(c, &payload) {
		return
	}

	payload.Title = trimSpace(payload.Title)
	payload.Type = trimSpace(payload.Type)
	payload.Status = trimSpace(payload.Status)
	payload.Location = trimSpace(payload.Location)
	payload.BookedBy = trimSpace(payload.BookedBy)

	if !requireNonEmpty(c, "title", payload.Title) ||
		!requireNonEmpty(c, "type", payload.Type) ||
		!requireNonEmpty(c, "location", payload.Location) ||
		!requireNonEmpty(c, "bookedBy", payload.BookedBy) {
		return
	}

	startTime, err := parseTime(payload.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startTime must be a valid datetime"})
		return
	}
	endTime, err := parseTime(payload.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endTime must be a valid datetime"})
		return
	}

	if !endTime.After(startTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endTime must be after startTime"})
		return
	}
	if payload.Status == "" {
		payload.Status = "Active"
	}
	if !validateStatus(c, "status", payload.Status, []string{"Active", "Upcoming", "Completed", "Cancelled"}) {
		return
	}

	reservation := models.Reservation{
		Title:     payload.Title,
		Type:      payload.Type,
		Status:    payload.Status,
		StartTime: startTime,
		EndTime:   endTime,
		Location:  payload.Location,
		BookedBy:  payload.BookedBy,
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	var existing models.Reservation
	if err := db.Where(
		"type = ? AND location = ? AND status <> ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?))",
		reservation.Type,
		reservation.Location,
		"Cancelled",
		reservation.StartTime,
		reservation.StartTime,
		reservation.EndTime,
		reservation.EndTime,
		reservation.StartTime,
		reservation.EndTime,
	).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Reservation overlaps with an existing booking"})
		return
	}

	if err := db.Create(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	c.JSON(http.StatusCreated, reservation)
}

type reservationUpdateRequest struct {
	Title     *string `json:"title"`
	Type      *string `json:"type"`
	Status    *string `json:"status"`
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
	Location  *string `json:"location"`
	BookedBy  *string `json:"bookedBy"`
}

func UpdateReservation(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var item models.Reservation
	if err := db.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	var payload reservationUpdateRequest
	if !bindJSON(c, &payload) {
		return
	}

	if payload.Title != nil {
		item.Title = trimSpace(*payload.Title)
	}
	if payload.Type != nil {
		item.Type = trimSpace(*payload.Type)
	}
	if payload.Status != nil {
		item.Status = trimSpace(*payload.Status)
	}
	if payload.StartTime != nil {
		parsed, err := parseTime(*payload.StartTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "startTime must be a valid datetime"})
			return
		}
		item.StartTime = parsed
	}
	if payload.EndTime != nil {
		parsed, err := parseTime(*payload.EndTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "endTime must be a valid datetime"})
			return
		}
		item.EndTime = parsed
	}
	if payload.Location != nil {
		item.Location = trimSpace(*payload.Location)
	}
	if payload.BookedBy != nil {
		item.BookedBy = trimSpace(*payload.BookedBy)
	}

	if !requireNonEmpty(c, "title", item.Title) ||
		!requireNonEmpty(c, "type", item.Type) ||
		!requireNonEmpty(c, "location", item.Location) ||
		!requireNonEmpty(c, "bookedBy", item.BookedBy) {
		return
	}
	if item.StartTime.IsZero() || item.EndTime.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startTime and endTime are required"})
		return
	}
	if !item.EndTime.After(item.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endTime must be after startTime"})
		return
	}
	if item.Status == "" {
		item.Status = "Active"
	}
	if !validateStatus(c, "status", item.Status, []string{"Active", "Upcoming", "Completed", "Cancelled"}) {
		return
	}

	var existing models.Reservation
	if err := db.Where(
		"id <> ? AND type = ? AND location = ? AND status <> ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?))",
		item.ID,
		item.Type,
		item.Location,
		"Cancelled",
		item.StartTime,
		item.StartTime,
		item.EndTime,
		item.EndTime,
		item.StartTime,
		item.EndTime,
	).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Reservation overlaps with an existing booking"})
		return
	}

	if err := db.Save(&item).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(200, item)
}

func DeleteReservation(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	db, ok := requireDB(c)
	if !ok {
		return
	}
	deleteEntity(c, db, &models.Reservation{}, id)
}
