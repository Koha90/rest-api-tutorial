## Сервис на Go.

# rest-api-tutorial

# user-service

# REST API

---
* GET /api/v1/users -- list of users -- 200, 404, 500
* GET /api/v1/users/:id -- user by ID -- 200, 404, 500
* POST /api/v1/users/:id -- create user -- 204, 4xx, Header Location: url
* PUT /api/v1/users/:id -- fully update user -- 204/200, 404, 500
* PATCH /api/v1/users/:id -- partially update user -- 204/200, 404, 400, 500
* DELETE /api/v1/users/:id -- delete user by id -- 204, 404, 400