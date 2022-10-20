package intities

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Cod          string  `json:"cod"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
}
