package config

import (
	"github.com/magiconair/properties"
	"log"
	"os"
)

const envVar = "APP_PROFILE"

var PROFILE = "dev"
var PROPERTIES *properties.Properties

func init() {
	if profile := os.Getenv(envVar); profile != "" {
		PROFILE = profile
	}
	log.Println("Loading the properties for profile: " + PROFILE)
	PROPERTIES = properties.MustLoadFile("resources/application-"+PROFILE+".properties", properties.UTF8)
}


func GetProfile() string {
	return PROFILE
}

func GetProperty(key string) string {
	return PROPERTIES.MustGet(key)
}

func GetProjectId() string {
	return GetProperty("coral.googlecloud.projectid")
}

func GetDatawarehousePublisherEnabled() bool {
	return PROPERTIES.GetBool("coral.dwpublisher.enabled", false)
}

func GetSessionTimeout() int64 {
	return int64(PROPERTIES.MustGetInt("coral.session.timeout"))
}

func GetCollectionName(name string) string {
	return PROPERTIES.MustGet("coral.firestore.collection."+ name);
}