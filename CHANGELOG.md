# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## 0.3.0 - 2023-03-18
### Package updates
- assignment-2 bumped to assignment-2-0.2.0
### Global changes

- - -

## 0.2.0 - 2023-03-18
### Package updates
- assignment-2 bumped to assignment-2-0.1.0
### Global changes
#### Miscellaneous Chores
- **(cog.toml)** add assignment-2 package - (fb991df) - tfkhdyt

- - -

## 0.1.0 - 2023-03-18
#### Bug Fixes
- **(order-pg)** change return type from *[]entity.Order to []entity.Order - (7dab7d5) - tfkhdyt
- **(order-pg)** change internal server error to bad request error when creating order - (5e083f7) - tfkhdyt
- **(order-pg)** change item creation method to Association mode - (ea37d45) - tfkhdyt
- **(order-repo)** remove pointer from return type - (2c195fd) - tfkhdyt
- **(order-service)** change for range items to eachOrder.Items - (0194d35) - tfkhdyt
- **(order-service)** change return type of GetAllOrders - (0a7fc60) - tfkhdyt
#### Build system
- **(a1)** build project - (f69888a) - tfkhdyt
- **(deps)** run go mod tidy - (17db969) - tfkhdyt
- **(deps)** install needed deps - (8c5478a) - tfkhdyt
- create makefile and build project for 3 OS' - (4cb834f) - tfkhdyt
#### Features
- **(a1/biodata)** finish all feature - (414e5ed) - tfkhdyt
- **(a1/data)** create students data - (7580b09) - tfkhdyt
- **(a1/domain)** create student entity - (7c40323) - tfkhdyt
- **(a1/lib)** add comment above functions to explain what that function does - (738cbe2) - tfkhdyt
- **(a1/lib)** create function to find student by name - (74dead6) - tfkhdyt
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
- **(project-pertama)** add main.go - (a622307) - tfkhdyt
- **(repo)** create ItemRepository interface - (39dccc2) - tfkhdyt
- **(service)** create order service - (a3191a9) - tfkhdyt
- **(sesi-1)** learn about const and operator - (99cef0a) - tfkhdyt
- **(sesi-1)** learn about data type - (0c0398b) - tfkhdyt
- **(sesi-2)** learn about array, slice, and map - (6d734a9) - tfkhdyt
- **(sesi-2)** learn about looping - (764a5ac) - tfkhdyt
- **(sesi-3)** learn about struct - (fcdac20) - tfkhdyt
- **(sesi-3)** learn about pointer - (9905598) - tfkhdyt
- **(sesi-3)** learn about closure - (df07262) - tfkhdyt
- **(sesi-3)** learn about function - (3602ff2) - tfkhdyt
- **(sesi-4)** learn about reflect - (f6b079f) - tfkhdyt
- **(sesi-5)** learn about buffered channel - (76f9ff1) - tfkhdyt
- **(sesi-7)** learn about gorm - (06474ff) - tfkhdyt
- **(sesi-8)** learn about DDD - (7224641) - tfkhdyt
- **(sesi-8)** learn about url parsing - (772b633) - tfkhdyt
- **(sesi-8)** learn about json decoding and parsing - (b593704) - tfkhdyt
- **(startapp)** add order repo instance - (a0a9972) - tfkhdyt
- **(tugas-dokumen)** add summary database - (e7b3ebe) - tfkhdyt
- **(tugas-dokumen)** add summary golang tingkat pemula - (83a6fa1) - tfkhdyt
- add sesi-7, learn about gin and native sql - (807287b) - tfkhdyt
- add sesi 6 - (0ce894e) - tfkhdyt
- add sesi-5 - (8afcfce) - tfkhdyt
- add sesi-2 - (df12dca) - tfkhdyt
#### Miscellaneous Chores
- **(a1/lib)** rename lib file name to camelCase - (7daae80) - tfkhdyt
- **(assignments)** create assignment 1 - (974fbcf) - tfkhdyt
- **(gitignore)** add bin/ - (5d05f28) - tfkhdyt
- **(project-pertama)** setup editorconfig - (934a567) - tfkhdyt
- **(tugas-dokumen)** add name to "summary golang tingkat pemula" file name - (c717a91) - tfkhdyt
- initialize cog.toml - (7131eb9) - tfkhdyt
- initialize sesi-3 - (25d7d40) - tfkhdyt
#### Refactoring
- **(a1)** rename domain package to entity - (a17ebf0) - tfkhdyt
- **(a1)** move print output process to separate function - (cbfc30b) - tfkhdyt
- **(a1)** move parse name process to a function - (69fde71) - tfkhdyt
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
- **(project-pertama)** comment out all statement - (fc8741b) - tfkhdyt

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).