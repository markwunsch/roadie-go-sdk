package roadie

import "time"

type Item struct {
	// Length is the length in inches of the item
	Length float32
	// Width is the width in inches of the item
	Width float32
	// Height is the height in inches of the item
	Height float32
	// Weight is the weight in pounds of the item
	Weight float32
	// Quantity is the number of items included in the shipment
	Quantity int
	// Value is the monetary value of the item
	Value float32
	// Description is the description of the item. Max length 200 characters
	Description string
	// ReferenceId is the user supplied ID for the item. Max length 100 characters
	ReferenceId string
}

type Location struct {
	// Address is the address information (required)
	Address Address
	// Contact is the contact information (required for shipment)
	Contact Contact
	// Notes is any additional information the drivers needs regarding the location. Max length 500 characters
	Notes string
}

type Address struct {
	// Name is the name of the location if applicable. This helpful for the driver if the location is a business. Max length 200 characters
	Name string
	// StoreNumber is the store identifier, if the address is a retail location. Max length 20 characters
	StoreNumber string
	// Street1 is the first line of the address. Max length 200 characters (required)
	Street1 string
	// Street2 is the first line of the address. Max length 200 characters
	Street2 string
	// City is the city. Max length 200 characters (required)
	City string
	// State is the two letter state code (required)
	State string
	// Zip is the postal code. Max length 10 characters (required)
	Zip string
	// Latitude us tge exact pickup/delivery latitude
	Latitude float32
	// Longitude is the exact pickup/delivery longitude
	Longitude float32
}

type Contact struct {
	// Name is the name of the contact. Max length 200 characters (required)
	Name string
	// Phone is the phone number of the contact. Max length 10 characters (required)
	Phone string
}

type DeliveryOptions struct {
	// SignatureRequired indicates whether driver must receive a signature from the recipient (required)
	SignatureRequired bool
	// NotificationsEnabled indicates whether the recipient should receive SMS updates for their delivery
	NotificationsEnables bool
	// Over21Required indicates if the driver must be over 21 to be eligible to deliver the item (typically for alcohol deliveries)
	Over21Required bool
	// ExtraCompensation is the additional compensation for the driver prior to assignment
	ExtraCompensation float32
	// TrailerRequired indicates if the driver must use a trailer to deliver the item(s)
	TrailerRequired bool
}

type TimeWindow struct {
	// Start is the start of the window in RFC 3339 format (required)
	Start time.Time
	// End is the end of the window in RFC 3339 format (required)
	End time.Time
}

type Driver struct {
	// Name is the driver's name
	Name string
	// Phone is the driver's phone number
	Phone string
}

type ShipmentEvent struct {
	// Name is the name of the event type
	Name string
	// OccurredAt is the time when the event occurred
	OccurredAt time.Time
	// Location is the location of the event
	Location Location
}
