package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	List(pageSize int, page int) ([]Order, error)
	// GetTotal() (int, error)
}
