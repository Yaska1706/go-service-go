package api

type API struct {
	PlaceholderService
}

func NewApi(placeHolderService PlaceholderService) API {
	return API{placeHolderService}
}
