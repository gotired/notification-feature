package model

type Order string

const (
	ASC  Order = "asc"
	Desc Order = "desc"
)

type SearchOptions struct {
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Keyword  string `json:"keyword"`
	Order    Order  `json:"order"`
	OrderKey string `json:"order_key"`
}

func (o *SearchOptions) Default() {
	if o.Limit == 0 {
		o.Limit = 10
	}
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Order == "" {
		o.Order = ASC
	}
	if o.OrderKey == "" {
		o.OrderKey = "updated_at"
	}
}
