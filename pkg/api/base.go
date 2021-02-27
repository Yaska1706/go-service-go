package api

type Services struct {
	PlaceholderService
}

func NewApi(placeHolderService PlaceholderService) Services {
	return Services{placeHolderService}
}
