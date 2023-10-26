package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hp-pandey/Product-Service/mongo"
	order "github.com/hp-pandey/Product-Service/orderService"
	product "github.com/hp-pandey/Product-Service/productservice"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func main() {
	mongo.InitMongoDBConnection()

	product.InitProductCollection(mongo.Client, "ProductService", "product")
	order.InitOrderCollection(mongo.Client, "ProductService", "order")

	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		productInfo, err := product.GetProducts()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, productInfo)
	})

	r.POST("/order", func(c *gin.Context) {
		var newOrder order.Order
		if err := c.BindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
		if newOrder.Quantity > 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Maixmum limitation reached"})
			return
		}
		newOrder.ID = primitive.NewObjectID()
		pId, err := primitive.ObjectIDFromHex(newOrder.ProductId)
		product1, err := product.GetProductById(pId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product Not available"})
			return
		}

		// Logic to validate order, apply discounts and availability of Product.
		if product1.Availability-newOrder.Quantity < 0 {
			c.JSON(http.StatusBadRequest, "Product not Available")
			return
		}
		if newOrder.IsPremium {
			if newOrder.Quantity >= 3 {
				discount := (10 / 100) * newOrder.OrderValue
				newOrder.OrderValue = newOrder.OrderValue - discount
			}
		}
		err = order.CreateOrder(&newOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}
		product1.Availability = product1.Availability - newOrder.Quantity
		err = product.UpdateProduct(product1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update Product"})
			return
		}
		c.JSON(http.StatusCreated, newOrder)
	})

	r.PUT("/order/:id", func(c *gin.Context) {
		orderID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid OrderId"})
			return
		}
		err = order.UpdateOrderStatus(orderID, c.Query("status"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
	})

	// Run Server
	r.Run(":8080")
}
