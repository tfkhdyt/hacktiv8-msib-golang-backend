# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## assignment-2-0.4.0 - 2023-03-19
#### Bug Fixes
- **(order-dto)** remove omitempty from json tag - (99819ec) - tfkhdyt
#### Build system
- **(deps)** install swagger - (6cc39c4) - tfkhdyt
#### Documentation
- **(app)** setup gin-swagger - (b5f8ec6) - tfkhdyt
- **(order-handler)** add swagger API docs - (186602d) - tfkhdyt
- **(swagger)** generate API docs - (2874dfb) - tfkhdyt
#### Features
- **(order-service)** add items to returned response - (770caee) - tfkhdyt
#### Miscellaneous Chores
- **(assignment-2)** add .env.example - (50da244) - tfkhdyt
- **(assignment-2)** add link txt file for submission - (646658d) - tfkhdyt
- **(swagger)** bump swaggerinfo version to 0.4 - (2457a60) - tfkhdyt
#### Refactoring
- **(app)** change port to 8080, move SwaggerInfo initialization to init function - (420d772) - tfkhdyt
- **(errs)** rename messageErrData to MessageErrData - (88b736a) - tfkhdyt
#### Style
- **(order-service)** remove unneeded newline on CreateOrder - (78f3665) - tfkhdyt

- - -

## assignment-2-0.3.0 - 2023-03-18
#### Bug Fixes
- **(order-pg)** change not found error message to "Order with id %d is not found" - (1488245) - tfkhdyt
#### Features
- **(order-dto)** add DeleteOrderByIDResponse dto - (9eefe5c) - tfkhdyt
- **(order-handler)** add DeleteOrderByID method - (97a4573) - tfkhdyt
- **(order-pg)** add DeleteOrderByID implementation - (b0bc908) - tfkhdyt
- **(order-repo)** add DeleteOrderByID method definition - (00b3e0e) - tfkhdyt
- **(order-service)** add DeleteOrderByID method - (bd85814) - tfkhdyt
- **(routes)** add DELETE route for order - (a3f3c55) - tfkhdyt

- - -

## assignment-2-0.2.0 - 2023-03-18
#### Bug Fixes
- **(item-entity)** remove unique constraint from ItemCode property - (137892f) - tfkhdyt
- **(order-pg)** change return value of UpdateOrderByID to order - (1ad68ce) - tfkhdyt
- **(order-service)** change return type of UpdateOrderByID to *dto.GetOrderByIDResponse - (bd908b6) - tfkhdyt
#### Features
- **(dto)** add GetOrderByIDResponse dto - (5759ec7) - tfkhdyt
- **(errs)** add NewNotFound error - (7e4c483) - tfkhdyt
- **(http-handler)** add UpdateOrderByID method - (2c6c54d) - tfkhdyt
- **(order-handler)** add GetOrderByID method - (6fa39fe) - tfkhdyt
- **(order-pg)** add UpdateOrderByID method - (0d3c268) - tfkhdyt
- **(order-pg)** add GetOrderByID method - (3816408) - tfkhdyt
- **(order-service)** add UpdateOrderByID method - (114b4da) - tfkhdyt
- **(order-service)** add GetOrderByID method - (d95f7a5) - tfkhdyt
- **(route)** add route for get order by id - (f442959) - tfkhdyt
- **(routes)** add route to update order by id - (7769bc3) - tfkhdyt

- - -

## assignment-2-0.1.0 - 2023-03-18
#### Bug Fixes
- **(order-pg)** change return type from *[]entity.Order to []entity.Order - (7dab7d5) - tfkhdyt
- **(order-pg)** change internal server error to bad request error when creating order - (5e083f7) - tfkhdyt
- **(order-pg)** change item creation method to Association mode - (ea37d45) - tfkhdyt
- **(order-repo)** remove pointer from return type - (2c195fd) - tfkhdyt
- **(order-service)** change for range items to eachOrder.Items - (0194d35) - tfkhdyt
- **(order-service)** change return type of GetAllOrders - (0a7fc60) - tfkhdyt
#### Build system
- **(deps)** run go mod tidy - (17db969) - tfkhdyt
- **(deps)** install needed deps - (8c5478a) - tfkhdyt
#### Features
- **(app)** apply GetAllOrders handler to GET /orders route - (49a29fd) - tfkhdyt
- **(assignment2)** invoke StartApp function - (fab6b2a) - tfkhdyt
- **(assignment2)** create order handler - (a9b4821) - tfkhdyt
- **(config)** setup godotenv - (e31d076) - tfkhdyt
- **(config)** create postgres config - (2f113e0) - tfkhdyt
- **(database)** add success log when connected to db - (7729f8f) - tfkhdyt
- **(database)** create database instance - (7da6315) - tfkhdyt
- **(dto)** add GetAllOrdersResponse dto - (ecf7733) - tfkhdyt
- **(dto)** add GetAllItemsResponse dto - (be89e9d) - tfkhdyt
- **(dto)** add omitempty tag to items field in NewOrderRequest - (d4b628b) - tfkhdyt
- **(dto)** add item and order dto - (5ad6eea) - tfkhdyt
- **(entity)** create item and order entity - (72321cd) - tfkhdyt
- **(errs)** add NewBadRequest - (e5ff709) - tfkhdyt
- **(errs)** add NewUnprocessableEntity error - (07880a8) - tfkhdyt
- **(handler)** assign routers - (c87efac) - tfkhdyt
- **(handler)** create StartApp function - (b5edcdf) - tfkhdyt
- **(order-entity)** add foreign key constraint - (70bc14c) - tfkhdyt
- **(order-handler)** add GetAllOrders method - (289d1a7) - tfkhdyt
- **(order-repo)** add GetAllOrders method - (b270099) - tfkhdyt
- **(order-repo)** create postgres implementation of OrderRepository - (39d171f) - tfkhdyt
- **(order-repo)** create OrderRepository interface - (4c91d12) - tfkhdyt
- **(order-service)** add GetAllOrders method - (fbf35b5) - tfkhdyt
- **(pkg)** create custom error - (e62fe14) - tfkhdyt
- **(repo)** create ItemRepository interface - (39dccc2) - tfkhdyt
- **(service)** create order service - (a3191a9) - tfkhdyt
- **(startapp)** add order repo instance - (a0a9972) - tfkhdyt
#### Refactoring
- **(errs)** rename MessageErrData struct to messageErrData - (f1556cb) - tfkhdyt
- **(item-dto)** rename ItemDTO to ItemData - (7a20872) - tfkhdyt
- **(item-dto)** rename GetItemResponse to ItemDTO - (c06f88d) - tfkhdyt
- **(item-dto)** rename GetAllItemsResponse to GetItemResponse - (a4ee15a) - tfkhdyt
- **(item-dto)** Change Id to ID - (876aaeb) - tfkhdyt
- **(order-dto)** rename OrderDTO to OrderData - (b1317e9) - tfkhdyt
- **(order-dto)** rename GetOrderResponse to OrderDTO - (2be7e44) - tfkhdyt
- **(order-dto)** rename OrderResponse to GetOrderResponse - (02abc5e) - tfkhdyt
- **(order-dto)** separate GetAllOrdersResponse to OrderResponse - (1507f6e) - tfkhdyt
- **(order-repo)** rename itemPayloads to itemsPayload - (f485b0d) - tfkhdyt
- **(order-service)** apply rename refactor from dto - (993d626) - tfkhdyt
- **(repo)** delete item_repository - (e65a14d) - tfkhdyt
#### Style
- **(assignment2)** format with gofumpt - (5f5d4a9) - tfkhdyt

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).