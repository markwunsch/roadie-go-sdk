# Roadie Go SDK
[![GoDoc](https://godoc.org/github.com/markwunsch/roadie-go-sdk/roadie?status.svg)](https://godoc.org/github.com/markwunsch/roadie-go-sdk/roadie) [![Go Report Card](https://goreportcard.com/badge/github.com/markwunsch/roadie-go-sdk)](https://goreportcard.com/report/github.com/markwunsch/roadie-go-sdk)

roadie-go-sdk is a client library for accessing the [Roadie API](https://docs.roadie.com/#api-overview)

## Installation

```console
go get github.com/markwunsch/roadie-go-sdk/roadie
```

## Usage

The `roadie` package provides a `Client` for accessing the Roadie API. Authentication is handled by including your access token in the `WithAccessToken` function, as shown below. 

```go
client, err := roadie.NewClient(roadie.WithAccessToken(context.Background(), "YOUR-ACCESS-TOKEN"))

// Get shipment by ID
shipment, err := client.Shipments.Get(context.Background(), 413042)

// Create an estimate and get the response from the roadie API
response, err := client.Estimates.Create(context.Background(), roadie.CreateEstimateRequest{
    Items: []roadie.Item{
        roadie.Item{
            Length:      4.3,
            Width:       1.0,
            Height:      4.2,
            Weight:      1.1,
            Quantity:    1,
            Value:       25.00,
            Description: "Item description",
            ReferenceId: "UUID",
        },
    },
    PickupLocation: roadie.Location{
        Address: roadie.Address{
            Street1: "7778 McGinnis Ferry Rd #270",
            City: "Suwanee",
            State: "GA",
            Zip: "30024",
        },
    },
    DeliveryLocation: roadie.Location{
        Address: roadie.Address{
            Street1: "100 North Ave",
            City: "Atlanta",
            State: "GA",
            Zip: "30332",
        },
    },
    PickupAfter: time.Now(),
    DeliverBetween: roadie.TimeWindow{
        Start: time.Now(),
        End: time.Now().Add(time.Hour * 4),
    },
})
```

## Supported Endpoints

Method | HTTP request                                     | Description
------------- |--------------------------------------------------| -------------
[*Create Estimate*](https://docs.roadie.com/#estimates) | **POST** /estimate                               | Create an estimate
[*Create Shipment*](https://docs.roadie.com/#create-a-shipment) | **POST** /shipments                              | Create a shipment
[*Retrieve Shipment*](https://docs.roadie.com/#retrieve-a-shipment) | **GET** /shipments/{shipment-id}                 | Retrieve a shipment
[*Retrieve List of Shipments*](https://docs.roadie.com/#retrieve-a-list-of-shipments) | **GET** /shipments/reference_ids={reference-ids} | Retrieve a list of shipments
[*Update Shipment*](https://docs.roadie.com/#update-a-shipment) | **PATCH** /shipments/{shipment-id}               | Update a shipment
[*Cancel Shipment*](https://docs.roadie.com/#cancel-a-shipment) | **DELETE** /shipments/{shipment-id}              | Cancel a shipment
[*Tip the driver*](https://docs.roadie.com/#leave-a-tip-for-the-driver) | **POST** /shipments/{shipment-id}/tips           | Leave a tip for the driver
[*Rate the driver*](https://docs.roadie.com/#leave-a-rating-for-the-driver) | **POST** /shipments/{shipment-id}/ratings        | Rate the driver







