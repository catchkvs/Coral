package repo

import (
	"github.com/catchkvs/Coral/pkg/config"
	"github.com/catchkvs/Coral/pkg/model"
)


func SaveFactEntity(factEntity *model.FactEntity) {
	AddWithId(factEntity.Name, factEntity.Id, factEntity)

}

func SaveDimensionEntity(dimensionEntity *model.DimensionEnty) {
	AddWithId(config.GetCollectionName(dimensionEntity.Name), dimensionEntity.Id, dimensionEntity)
}

func GetFactEntity(name, id string) *model.FactEntity {
	docSnapshot := Get(config.GetCollectionName(name), id)
	var factEntity model.FactEntity
	docSnapshot.DataTo(&factEntity)
	return &factEntity
}

func GetDimensionEntity(name , id string) *model.DimensionEnty{
	docSnapshot := Get(config.GetCollectionName(name), id)
	var dimensionentity model.DimensionEnty
	docSnapshot.DataTo(&dimensionentity)
	return &dimensionentity
}