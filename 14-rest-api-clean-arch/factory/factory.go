package factory

import (
	"be9/restclean/config"
	_userBusiness "be9/restclean/features/users/business"
	_userData "be9/restclean/features/users/data"
	_userPresentation "be9/restclean/features/users/presentation"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
	// ProductPresenter *_productPresentation.ProductHandler
}

func InitFactory() Presenter {
	dbConn := config.InitDB()

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	// productData := _productData.NewProductRepository(dbConn)
	// productBusiness := _productBusiness.NewProductBusiness(productData)
	// productPresentation := _productPresentation.NewProductHandler(productBusiness)

	return Presenter{
		UserPresenter: userPresentation,
		// ProductPresenter: productPresentation,
	}
}
