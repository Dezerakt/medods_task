package repositories

type TokenRepository struct {
	*MainRepository
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{MainRepository: NewMainRepository()}
}
