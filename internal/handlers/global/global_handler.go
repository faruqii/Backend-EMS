package handlers

import service "github.com/Magetan-Boyz/Backend/internal/services/global"

type GlobalHandler struct {
	globalService service.GlobalService
}

func NewGlobalHandler(globalService service.GlobalService) *GlobalHandler {
	return &GlobalHandler{
		globalService: globalService,
	}
}
