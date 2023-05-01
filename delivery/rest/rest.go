package rest

import (
	"errors"
	"net/http"

	d "github.com/fazarrahman/cognotiv/delivery"
	"github.com/fazarrahman/cognotiv/lib"
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
	g.POST("/order", d.MustAuthorize(), r.PostOrder)
	g.GET("/order", d.MustAuthorize(), r.GetOrders)
}

func (r *Rest) PostOrder(c *gin.Context) {
	var req model.Orders
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("error parsing body"))
		return
	}
	username := lib.GetUsernameFromToken(c)
	err := r.svc.CreateOrder(c, &req, *username)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (r *Rest) GetOrders(c *gin.Context) {
	username := lib.GetUsernameFromToken(c)
	orderList, err := r.svc.GetOrderListByUserID(c, *username)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, orderList)
}
