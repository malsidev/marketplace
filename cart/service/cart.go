package service

// import (
// 	"order/models"
// 	"order/repository"
// )

// type CartService struct {
// 	cartRepository *repository.CartRepository
// }

// func (s CartService) CheckCart(userid int) (*models.Carts, error) {
// 	cart, err := s.cartRepository.CheckCart(userid)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if cart != nil {
// 		return cart, nil
// 	}

// 	cart = &models.Carts{
// 		UserID: userid,
// 	}
// 	return cart, nil

// }
