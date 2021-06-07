package http

import (
	"github.com/amartery/tp_db_forum/internal/app/service"
	"github.com/amartery/tp_db_forum/internal/pkg/utils"
	"github.com/valyala/fasthttp"
)

type ServiceHandler struct {
	usecaseService service.Usecase
}

func NewServiceHandler(serviceUsecase service.Usecase) *ServiceHandler {
	return &ServiceHandler{
		usecaseService: serviceUsecase,
	}
}

// TODO: s.router.POST("​/service​/clear", s.ServiceClear)
// TODO: s.router.GET("​/service​/status", s.ServiceStatus)

func (handler *ServiceHandler) ServiceClear(ctx *fasthttp.RequestCtx) {
	err := handler.usecaseService.ClearDB()
	if err != nil {
		utils.SendServerError(err.Error(), fasthttp.StatusInternalServerError, ctx)
		return
	}
	utils.SendResponse(fasthttp.StatusOK, "", ctx)
}

func (handler *ServiceHandler) ServiceStatus(ctx *fasthttp.RequestCtx) {
	status, err := handler.usecaseService.GetStatusDB()
	if err != nil {
		utils.SendServerError(err.Error(), fasthttp.StatusInternalServerError, ctx)
		return
	}
	// TODO: сделать easy json
	utils.SendResponse(fasthttp.StatusOK, status, ctx)
}
