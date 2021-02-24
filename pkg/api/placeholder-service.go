package api

//TODO: this is just a placeholder to get you started.

// PlaceholderService contains the methods of the dummy service
type PlaceholderService interface{}

// Placeholder repository is what lets our service do db operations without knowing anything about the implementation
type PlaceholderRepository interface{}

type placeholderService struct {
	storage PlaceholderRepository
}

func NewUserService(placeholderRepo PlaceholderRepository) PlaceholderService {
	return &placeholderService{placeholderRepo}
}
