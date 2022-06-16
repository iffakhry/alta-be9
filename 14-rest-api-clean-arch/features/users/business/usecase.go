package business

import "be9/restclean/features/users"

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) GetAllData(limit, offset int) (resp []users.Core, err error) {
	var param string
	resp, err = uc.userData.SelectData(param)
	return resp, err
}
