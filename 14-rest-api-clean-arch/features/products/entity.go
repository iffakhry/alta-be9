package products

import "time"

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData(param string) (data []Core, err error)
}

type Data interface {
	SelectData(param string) (data []Core, err error)
	// InsertData(data Core) (resp Core, err error)
}
