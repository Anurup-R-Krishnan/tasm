package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tasm-backend/models"
)

func GetAssets(c *gin.Context) {
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var assets []models.Asset
	if err := db.Order("created_at desc").Find(&assets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load assets"})
		return
	}

	if len(assets) == 0 {
		assets = []models.Asset{
			{
				Name:           "MacBook Pro 16\"",
				TagID:          "TAG-1082",
				Category:       "IT Equipment",
				Location:       "Bldg 3, Fl 4",
				Status:         "Checked Out",
				Custodian:      "Sarah Jenkins",
				CurrentStock:   1,
				ReorderLevel:   5,
				Value:          2499.99,
				PurchaseDate:   db.NowFunc().AddDate(-1, -6, 0),
				WarrantyStatus: "Active",
			},
			{
				Name:           "Dell Ultrasharp 27\"",
				TagID:          "TAG-2094",
				Category:       "IT Peripherals",
				Location:       "Stockroom A",
				Status:         "In Stock",
				Custodian:      "IT Admin",
				CurrentStock:   12,
				ReorderLevel:   10,
				Value:          399.99,
				PurchaseDate:   db.NowFunc().AddDate(-2, 0, 0),
				WarrantyStatus: "Expired",
			},
			{
				Name:           "Herman Miller Aeron",
				TagID:          "TAG-5032",
				Category:       "Office Furniture",
				Location:       "Bldg 1, Fl 2",
				Status:         "Reserved",
				Custodian:      "New Hire (Pending)",
				CurrentStock:   5,
				ReorderLevel:   2,
				Value:          1200.00,
				PurchaseDate:   db.NowFunc().AddDate(0, -1, 0),
				WarrantyStatus: "Active",
			},
		}
		for _, a := range assets {
			db.Create(&a)
		}
		db.Order("created_at desc").Find(&assets)
	}

	c.JSON(http.StatusOK, assets)
}

func CreateAsset(c *gin.Context) {
	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, ok := requireDB(c)
	if !ok {
		return
	}

	if err := db.Create(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create asset"})
		return
	}

	c.JSON(http.StatusCreated, asset)
}

func GetAssetByID(c *gin.Context) {
	id := c.Param("id")
	db, ok := requireDB(c)
	if !ok {
		return
	}

	var asset models.Asset
	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, asset)
}
