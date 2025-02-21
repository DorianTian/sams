package model

type Customer struct {
	CustId      string `json:"id" gorm:"primaryKey"`
	CustName    string `json:"custName"`
	CustAddress string `json:"address"`
	CustCity    string `json:"city"`
	CustState   string `json:"state"`
	CustZip     string `json:"zip"`
	CustEmail   string `json:"email"`
	CustContact string `json:"contact"`
	CustCountry string `json:"country"`
}

func (Customer) TableName() string {
	return "sams.public.customers"
}
