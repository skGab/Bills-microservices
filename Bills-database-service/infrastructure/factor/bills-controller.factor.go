package factor

import (
	"github.com/skGab/Bills-management-service/app/controllers"
	"github.com/skGab/Bills-management-service/app/usecases"
	"github.com/skGab/Bills-management-service/infrastructure/databases"
	"gorm.io/gorm"
)

func BillsController(db *gorm.DB) *controllers.BillsController {
	// DATABASE INSTANCE
	billsDatabase := &databases.BillsDatabase{
		DB: db,
	}

	// USECASES INSTANCE
	billsUsecases := &usecases.BillsUsecases{
		Repository: billsDatabase,
	}

	// CONTROLLER INSTANCE
	billsController := &controllers.BillsController{
		Usecases: billsUsecases,
	}

	return billsController
}