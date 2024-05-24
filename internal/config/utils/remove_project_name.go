package utils

import (
	"MyBalance/internal/projkeys"
	"strings"
)

func CheckStringAndRemovePrefix(key string) string {
	if strings.HasPrefix(key, projkeys.ProjectName+".") {
		// remove prefix
		if key2 := strings.Replace(key, projkeys.ProjectName+".", "", 1); key2 != "" {
			return key2
		} else {
			return key
		}
	} else if strings.HasPrefix(key, strings.ToLower(projkeys.ProjectName)+".") {
		// remove prefix
		if key2 := strings.Replace(key, strings.ToLower(projkeys.ProjectName)+".", "", 1); key2 != "" {
			return key2
		} else {
			return key
		}
	}

	return key
}
