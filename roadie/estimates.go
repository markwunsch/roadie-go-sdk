package roadie

import (
	"context"
	"time"
)

// EstimateService has methods for interacting with estimates in roadie
type EstimateService service

type CreateEstimateRequest struct {
	// Items is an array of one or more item (required)
	Items []Item
	// PickupLocation is a location object (required)
	PickupLocation Location
	// DeliveryLocation is a location object (required)
	DeliveryLocation Location
	// PickupAfter is the time when the shipment is ready for pickup
	PickupAfter time.Time
	// DeliverBetween is the window within which the shipment must be complete
	DeliverBetween TimeWindow
}

type CreateEstimateResponse struct {
	// Price is the estimated price
	Price float32
	// Size is the size category
	Size string
	// Estimated distance is the estimated distance between pickup and delivery
	EstimatedDistance float32
	// Errors contains any errors returned from the API
	Errors ErrorResponse
}

// Create creates an estimate
func (est *EstimateService) Create(ctx context.Context, estimate CreateEstimateRequest) (*CreateEstimateResponse, error) {
	req, err := est.client.createRequest("POST", "estimates", estimate)
	if err != nil {
		return nil, err
	}

	estimateResponse := new(CreateEstimateResponse)
	err = est.client.do(ctx, req, estimateResponse)
	if err != nil {
		return nil, err
	}
	return estimateResponse, nil
}
