package usecase

type CatUsecase struct {
}

func NewCatUsecase() CatUsecaseInterface {
	return &CatUsecase{}
}

func (u *CatUsecase) GetCatById(id int) interface{} {
	return nil
}
