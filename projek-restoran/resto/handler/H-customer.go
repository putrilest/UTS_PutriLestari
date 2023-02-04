package handler

import (
	"net/http"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func RoutesCust(customer resto.UsecaseCust, r *gin.RouterGroup) {
	u := Handler{
		customer,
	}

	v2 := r.Group("customers")
	v2.GET("", u.GetAll)
	v2.POST("", u.CreateCustomer)
	v2.PUT(":id_cust", u.PutCustomer)
	v2.DELETE(":id_cust", u.DeleteCustomer)
}

// Voucher bind with voucher usecase interface
type Handler struct {
	UC resto.UsecaseCust
}

func (customer Handler) GetAll(ctx *gin.Context) {
	result, err := customer.UC.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (customer Handler) CreateCustomer(ctx *gin.Context) {
	err := customer.UC.Create(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (customer Handler) PutCustomer(ctx *gin.Context) {
	err := customer.UC.Update(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (customer Handler) DeleteCustomer(ctx *gin.Context) {
	err := customer.UC.Delete(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
