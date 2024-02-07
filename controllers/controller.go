package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/e-hastono/bflpbootcamptit-golangassignment2/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	_, err := newOrder.CreateOrder()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    newOrder,
		"message": "Order successfully created",
		"status":  "success",
	})
}

func GetAllData(ctx *gin.Context) {
	orders, err := models.GetAllData()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"message": "All orders and items successfully fetched",
		"status":  "success",
	})
}

func UpdateOrderPatch(ctx *gin.Context) {
	rawOrderID := ctx.Param("orderid")
	orderID, err := strconv.Atoi(rawOrderID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var updatedOrder = models.Order{OrderID: uint(orderID)}

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	_, err = updatedOrder.UpdateOrderPatch()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    updatedOrder,
		"message": "Order successfully updated",
		"status":  "success",
	})
}

func UpdateOrderPut(ctx *gin.Context) {
	rawOrderID := ctx.Param("orderid")
	orderID, err := strconv.Atoi(rawOrderID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var updatedOrder = models.Order{OrderID: uint(orderID)}

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	_, err = updatedOrder.UpdateOrderPut()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    updatedOrder,
		"message": "Order successfully updated",
		"status":  "success",
	})
}

func DeleteOrder(ctx *gin.Context) {
	rawOrderID := ctx.Param("orderid")
	orderID, err := strconv.Atoi(rawOrderID)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	deletedOrderID, err := models.DeleteOrder(uint(orderID))

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    deletedOrderID,
		"message": fmt.Sprintf("Order data with id %d successfully deleted", deletedOrderID),
		"status":  "success",
	})
}
