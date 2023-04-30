package rest

import (
	"errors"
	"net/http"

	"github.com/fazarrahman/cognotiv/model"
	"github.com/fazarrahman/cognotiv/service"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	svc *service.Svc
}

// New ...
func New(svc *service.Svc) *Rest {
	return &Rest{svc: svc}
}

func (r *Rest) Register(g *gin.RouterGroup) {
	g.POST("/order", r.Order)
}

func (r *Rest) Order(c *gin.Context) {
	var req model.Orders
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("error parsing body"))
		return
	}
	err := r.svc.CreateOrder(c, &req)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, "success")
}
