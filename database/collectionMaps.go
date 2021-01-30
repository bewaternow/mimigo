package database

func InitCollections() {
	SupportUser = MongoDB.Database("Support").Collection("users")
	SupportPersonalAccessToken = MongoDB.Database("Support").Collection("personal_access_tokens")
}
