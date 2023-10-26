# Product and Order Services

This repository contains two Go (Golang) services - ProductService and OrderService - designed to provide information
about products and manage orders. These services are built using Go and utilize MySQL/PostgreSQL as the database
backend.

Services Overview ProductService Provides information about products, including availability, price, and category.
Offers endpoints for retrieving the product catalog. OrderService Provides information about orders, including order
value, dispatch date, and order status. Allows users to place new orders and update order statuses.

## Setup Instructions

### 1.Clone the Repository:

git clone https://github.com/hp-pandey/Product-Service_Repo.git

### 2.Database Setup:

Create a MongoDB database and update the connection details in productservice/database.go and orderservice/database.go.

### 3.Install Dependencies:

go mod tidy

### 4.Run Services:

cd Product-Service_Repo/

go run main.go

### 5.API Endpoints:

#### a.ProductService:

GET "/products" : Get product catalog.

#### b.OrderService:

POST "/order" : Place a new order.

PUT /orders/{id}/status: Update order status.

### 6.Postman Collection:

Link to Postman Collection
=> https://elements.getpostman.com/redirect?entityId=17766028-686cbd4e-4656-4c5e-8ca0-12de103e5461&entityType=collection
