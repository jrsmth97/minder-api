package controller

import (
	"minder/src/server/param"
	"minder/src/server/service"
	"minder/src/server/view"

	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	svc *service.PurchaseService
}

func NewPurchaseHandler(svc *service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{
		svc: svc,
	}
}

func (p *PurchaseHandler) GetPurchases(c *gin.Context) {
	resp := p.svc.GetPurchases()

	WriteJsonResponse(c, resp)
}

func (m *PurchaseHandler) CreatePurchase(c *gin.Context) {
	var req param.CreatePurchase
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(c, view.ErrBadRequest(err.Error()))
		return
	}

	err = param.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponse(c, resp)
		return
	}

	m.svc.Preparation(c)
	resp := m.svc.CreatePurchase(&req)

	WriteJsonResponse(c, resp)
}

func (m *PurchaseHandler) GetPurchase(c *gin.Context) {
	purchaseId, isExist := c.Params.Get("purchaseId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("purchase id not found"))
		return
	}

	resp := m.svc.GetPurchaseById(purchaseId)
	WriteJsonResponse(c, resp)
}

func (m *PurchaseHandler) CancelPurchase(c *gin.Context) {
	purchaseId, isExist := c.Params.Get("purchaseId")
	if !isExist {
		WriteJsonResponse(c, view.ErrBadRequest("purchase not found"))
		return
	}

	resp := m.svc.CancelPurchase(purchaseId)

	WriteJsonResponse(c, resp)
}

func (m *PurchaseHandler) SyncPurchase(c *gin.Context) {
	resp := m.svc.SyncPurchase()
	WriteJsonResponse(c, resp)
}
