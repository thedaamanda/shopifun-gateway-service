package payment

import (
	"encoding/json"
	"gateway-service/helper"
	model "gateway-service/models"
)

func CreatePayment(req interface{}) (*model.PaymentResponse, error) {
	var (
		paymentChannel = make(chan helper.Response)
		paymentUri     = "https://31c3-36-72-214-46.ngrok-free.app/payments/payments"
	)

	payload := helper.NewNetClientRequest(paymentUri, helper.DefaultNetClient)
	payload.Post(req, paymentChannel)

	paymentRes := <-paymentChannel
	if paymentRes.Err != nil {
		return nil, paymentRes.Err
	}

	var response *model.PaymentResponse
	if err := json.Unmarshal(paymentRes.Res, &response); err != nil {
		return nil, err
	}

	return response, nil
}
