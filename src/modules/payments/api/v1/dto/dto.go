package dto

type CreatePaymentReq struct {
	OrderId    string  `json:"orderId"`
	BillAmount float64 `json:"billAmount"`
	TipAmount  float64 `json:"tipAmount"`
	PaymentId  string  `json:"paymentId"`
}

type CreatePaymentRes struct {
}
