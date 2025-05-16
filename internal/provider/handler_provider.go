package provider

import "DRX_Test/internal/delivery/http/handler"

type Handlers struct {
	Root    *handler.AppHandler
	Product *handler.ProductHandler
}

func ProvideHandlers(usecase *Usecases) *Handlers {
	return &Handlers{
		Root:    &handler.AppHandler{},
		Product: handler.NewProductHandler(usecase.Product),
	}
}
