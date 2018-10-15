package config

import (
	"os"
)

// Get Value From Enviroment Vars
// MONGO_USERNAME_MLAB_DEV
// MONGO_PASSWORD_MLAB_DEV
// MONGO_HOST_MLAB_DEV
// MONGO_PORT_MLAB_DEV
// MONGO_DB_NAME_MLAB_DEV
func getEnvVar(envVar string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return ""
	}
	return value
}
