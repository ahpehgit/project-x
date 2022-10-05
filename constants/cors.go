package constants

func GetOrigins() []string {
	return []string{"*"}
}

func GetAllowMethods() []string {
	return []string{"GET", "HEAD", "OPTIONS", "POST", "PUT"}
}

func GetAllowHeaders() []string {
	return []string{"Origin", "Content-Type"}
}

func GetExposedHeaders() []string {
	return []string{"Content-Length"}
}
