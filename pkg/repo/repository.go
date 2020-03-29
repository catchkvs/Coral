package repo

import (
	"github.com/catchkvs/Coral/pkg/config"
	"github.com/catchkvs/Coral/pkg/model"
	"log"
)

func FindRecentFactEntities(restaurantId string) []model.FactEntity {


}

func SaveFactEntity(factEntity *model.FactEntity) {
	AddWithId(factEntity.Name, factEntity.Id, factEntity)

}

func SaveDimensionEntity(dimensionEntity *model.DimensionEnty) {
	AddWithId(config.GetCollectionName(dimensionEntity.Name), dimensionEntity.Id, dimensionEntity)
}

func getFactEntity(id string) {

}

func getDimensionEntity(id string) {

}