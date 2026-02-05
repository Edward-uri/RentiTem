package controller

import (
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"main/src/features/items/application"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct {
	createUC    *application.CreateItemUseCase
	listUC      *application.ListItemsUseCase
	listCatUC   *application.ListCategoriesUseCase
	createCatUC *application.CreateCategoryUseCase
	updateCatUC *application.UpdateCategoryUseCase
	deleteCatUC *application.DeleteCategoryUseCase
	getUC       *application.GetItemUseCase
	updateUC    *application.UpdateItemUseCase
	deleteUC    *application.DeleteItemUseCase
	storage     FileStorage
}

type FileStorage interface {
	Save(file *multipart.FileHeader) (string, error)
}

type ItemWithOwnerResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	PriceType   string  `json:"price_type"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image"`
	OwnerName   string  `json:"owner_name"`
	OwnerPhone  string  `json:"owner_phone"`
	IsAvailable bool    `json:"is_available"`
}

type UpdateItemRequest struct {
	Title       *string  `json:"title"`
	Price       *float64 `json:"price"`
	IsAvailable *bool    `json:"is_available"`
}

type categoryRequest struct {
	Name string `json:"name"`
}

func NewItemController(createUC *application.CreateItemUseCase, listUC *application.ListItemsUseCase, listCatUC *application.ListCategoriesUseCase, createCatUC *application.CreateCategoryUseCase, updateCatUC *application.UpdateCategoryUseCase, deleteCatUC *application.DeleteCategoryUseCase, getUC *application.GetItemUseCase, updateUC *application.UpdateItemUseCase, deleteUC *application.DeleteItemUseCase, storage FileStorage) *ItemController {
	return &ItemController{createUC: createUC, listUC: listUC, listCatUC: listCatUC, createCatUC: createCatUC, updateCatUC: updateCatUC, deleteCatUC: deleteCatUC, getUC: getUC, updateUC: updateUC, deleteUC: deleteUC, storage: storage}
}

func (c *ItemController) Create(ctx *gin.Context) {
	ownerIDVal, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ownerID := ownerIDVal.(uint)

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	priceStr := ctx.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	imagePath, err := c.storage.Save(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	input := application.CreateItemInput{
		Title:       ctx.PostForm("title"),
		Description: ctx.PostForm("description"),
		Price:       price,
		PriceType:   ctx.PostForm("price_type"),
		Category:    ctx.PostForm("category"),
		ImageURL:    imagePath,
		OwnerID:     ownerID,
	}

	item, err := c.createUC.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": item.ID})
}

func (c *ItemController) List(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	items, err := c.listUC.Execute(application.ListItemsInput{
		Category: ctx.Query("category"),
		Search:   ctx.Query("search"),
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// Categories returns the predefined categories list.
func (c *ItemController) Categories(ctx *gin.Context) {
	cats, err := c.listCatUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cats)
}

// CreateCategory creates a new category (protected).
func (c *ItemController) CreateCategory(ctx *gin.Context) {
	var req categoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat, err := c.createCatUC.Execute(req.Name)
	if err != nil {
		status := http.StatusBadRequest
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			status = http.StatusConflict
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, cat)
}

// UpdateCategory updates an existing category name (and slug).
func (c *ItemController) UpdateCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req categoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat, err := c.updateCatUC.Execute(uint(id64), req.Name)
	if err != nil {
		status := http.StatusBadRequest
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		} else if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			status = http.StatusConflict
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cat)
}

// DeleteCategory removes a category.
func (c *ItemController) DeleteCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.deleteCatUC.Execute(uint(id64)); err != nil {
		status := http.StatusBadRequest
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (c *ItemController) Detail(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := c.getUC.Execute(uint(id64))
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "item not found" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ItemWithOwnerResponse{
		ID:          res.Item.ID,
		Title:       res.Item.Title,
		Description: res.Item.Description,
		Price:       res.Item.Price,
		PriceType:   res.Item.PriceType,
		Category:    res.Item.Category,
		ImageURL:    res.Item.ImageURL,
		OwnerName:   res.OwnerName,
		OwnerPhone:  res.OwnerPhone,
		IsAvailable: res.Item.IsAvailable,
	})
}

func (c *ItemController) Update(ctx *gin.Context) {
	ownerIDVal, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ownerID := ownerIDVal.(uint)

	idParam := ctx.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req UpdateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := c.updateUC.Execute(uint(id64), ownerID, application.UpdateItemInput{
		Title:       req.Title,
		Price:       req.Price,
		IsAvailable: req.IsAvailable,
	})
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "item not found" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": item.ID})
}

func (c *ItemController) Delete(ctx *gin.Context) {
	ownerIDVal, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ownerID := ownerIDVal.(uint)

	idParam := ctx.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.deleteUC.Execute(uint(id64), ownerID); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		} else if err.Error() == "item not found" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
