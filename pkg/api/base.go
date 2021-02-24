package api

type API struct {
	UserService
}

func NewApi(usrService UserService) API {
	return API{usrService}
}
