package handler

type Handlers struct {
	Authentication *AuthenticationHandler
	Storage        *StorageHandler
	User           *UserHandler
}
