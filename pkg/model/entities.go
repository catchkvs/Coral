package model

// Store your fact data like order, trips etc.
type FactEntity struct {
	Id string
	DimensionId string
	Status string
	Attributes map[string] interface{}
}

// Store your dimension data like restaurant, driver which needs to receive the update.
type DimensionEnty struct {
	Id string
	Name string
	Type string
	Attributes map[string]interface{}

}

