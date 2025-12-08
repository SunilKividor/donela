package handler

type Handlers struct {
	Authentication *AuthenticationHandler
	SongHandler    *SongHandler
	Storage        *StorageHandler
	User           *UserHandler
}
