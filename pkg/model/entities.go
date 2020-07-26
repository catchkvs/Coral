package model

// Store your fact data like order, trips etc.
type FactEntity struct {
	Id string
	Type string
	Name string
	Dimension DimensionRef
	Status string
	Data interface{}
}

type DeviceSubscription struct {
	DeviceId string
	Topic string
}

type DimensionRef struct {
	Id string
}

// Store your dimension data like restaurant, driver which needs to receive the update.
type DimensionEnty struct {
	Id string
	Name string
	Type string
	Attributes map[string]interface{}

}

