package usecase

import (
	"CleanArch/internal/entity"
)

type ListOrderInputDTO struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}
type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute(input ListOrderInputDTO) ([]ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List(input.PageSize, input.Page)
	if err != nil {
		return []ListOrderOutputDTO{}, err
	}

	var dtos []ListOrderOutputDTO

	for _, order := range orders {
		dto := ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}
