package usecase

import (
	"ingredients-list/model"
	"ingredients-list/repository"
)

type IDishUsecase interface {
	GetAllDishes(userId uint) ([]model.DishResponse, error)
	GetDishById(userId, dishId uint) (model.DishResponse, error)
	CreateDish(dish model.Dish) (model.DishResponse, error)
	DeleteDish(userId, dishId uint) error
}

type dishUsecase struct {
	dr repository.IDishRepository
}

func NewDishUsecase(dr repository.IDishRepository) IDishUsecase {
	return &dishUsecase{dr}
}

func (du *dishUsecase) GetAllDishes(userId uint) ([]model.DishResponse, error) {
	dishes := []model.Dish{}
	if err := du.dr.GetAllDishes(&dishes, userId); err != nil {
		return nil, err
	}

	resDishes := []model.DishResponse{}
	for _, v := range dishes {
		dish := model.DishResponse{
			ID:        v.ID,
			DishName:  v.DishName,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resDishes = append(resDishes, dish)
	}

	return resDishes, nil
}

func (du *dishUsecase) GetDishById(userId, dishId uint) (model.DishResponse, error) {
	dish := model.Dish{}
	if err := du.dr.GetDishById(&dish, userId, dishId); err != nil {
		return model.DishResponse{}, err
	}

	resDish := model.DishResponse{
		ID:        dish.ID,
		DishName:  dish.DishName,
		CreatedAt: dish.CreatedAt,
		UpdatedAt: dish.UpdatedAt,
	}

	return resDish, nil
}

func (du *dishUsecase) CreateDish(dish model.Dish) (model.DishResponse, error) {
	if err := du.dr.CreateDish(&dish); err != nil {
		return model.DishResponse{}, err
	}

	resDish := model.DishResponse{
		ID:        dish.ID,
		DishName:  dish.DishName,
		CreatedAt: dish.CreatedAt,
		UpdatedAt: dish.UpdatedAt,
	}

	return resDish, nil
}

func (du *dishUsecase) DeleteDish(userId, dishId uint) error {
	if err := du.dr.DeleteDish(userId, dishId); err != nil {
		return err
	}

	return nil
}
