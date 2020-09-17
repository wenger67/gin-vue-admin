package response

import "panta/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
