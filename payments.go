package neonomics

import "context"

func (c *client) DomesticPayment(ctx context.Context, req *DomesticPaymentRequest) (*DomesticPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) DomesticScheduledPayment(ctx context.Context, req *DomesticScheduledPaymentRequest) (*DomesticScheduledPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) SEPAPayment(ctx context.Context, req *SEPAPaymentRequest) (*SEPAPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) SEPAScheduledPayment(ctx context.Context, req *SEPAScheduledPaymentRequest) (*SEPAScheduledPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetPaymentByID(ctx context.Context, ID string) (*GetPaymentByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) AuthorizePayment(ctx context.Context, paymentProduct, paymentId string) (*AuthorizePaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) CompletePayment(ctx context.Context, paymentProduct, paymentId string) (*CompletePaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}
