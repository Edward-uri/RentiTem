package infraestructure

import (
	"log"

	"main/src/features/items/application"
	"main/src/features/items/domain/entities"
	"main/src/features/items/infraestructure/controller"
	"main/src/features/items/infraestructure/persistence"
	"main/src/features/items/infraestructure/services"
	"main/src/features/users/domain"

	"gorm.io/gorm"
)

type Dependencies struct {
	Controller *controller.ItemController
}

func NewDependencies(db *gorm.DB, userRepo domain.UserRepository, uploadDir string) Dependencies {
	repo := persistence.NewGormItemRepository(db)
	catRepo := persistence.NewGormCategoryRepository(db)

	defaultCategories := []entities.Category{
		{Name: "Tecnologia"},
		{Name: "Herramientas"},
		{Name: "Hogar"},
		{Name: "Deportes"},
		{Name: "Vehiculos"},
		{Name: "Moda"},
		{Name: "Musica"},
		{Name: "Libros"},
		{Name: "Juguetes"},
		{Name: "Servicios"},
		{Name: "Otro"},
	}

	if err := catRepo.EnsureDefaults(defaultCategories); err != nil {
		log.Printf("failed seeding categories: %v", err)
	}
	storage := services.NewLocalStorage(uploadDir)

	createUC := application.NewCreateItemUseCase(repo, catRepo)
	listUC := application.NewListItemsUseCase(repo)
	listCatUC := application.NewListCategoriesUseCase(catRepo)
	createCatUC := application.NewCreateCategoryUseCase(catRepo)
	updateCatUC := application.NewUpdateCategoryUseCase(catRepo)
	deleteCatUC := application.NewDeleteCategoryUseCase(catRepo)
	getUC := application.NewGetItemUseCase(repo, userRepo)
	updateUC := application.NewUpdateItemUseCase(repo)
	deleteUC := application.NewDeleteItemUseCase(repo)

	ctrl := controller.NewItemController(createUC, listUC, listCatUC, createCatUC, updateCatUC, deleteCatUC, getUC, updateUC, deleteUC, storage)

	return Dependencies{Controller: ctrl}
}
