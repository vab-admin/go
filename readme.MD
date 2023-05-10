. ├── app │   └── admin │       ├── handler │       │   ├── admin_group.go │       │   ├── admin_rule.go │       │   ├── admin_rule_action.go │       │   ├── admin_user.go │       │   └── public.go │       ├── injector │       │   ├── app.go │       │   ├── injector.go │       │   ├── wire.go │       │   └── wire_gen.go │       ├── repository │       │   ├── admin_group.go │       │   ├── admin_group_rule.go │       │   ├── admin_rule.go │       │   ├── admin_rule_action.go │       │   ├── admin_user.go │       │   └── admin_user_group.go │       ├── router │       │   ├── api.go │       │   ├── middleware │       │   └── route.go │       ├── schema │       │   ├── admin_group.go │       │   ├── admin_menu.go │       │   ├── admin_menu_action.go │       │   ├── admin_user.go │       │   └── common.go │       └── service │           ├── admin_group.go │           ├── admin_group_rule.go │           ├── admin_rule.go │           ├── admin_rule_action.go │           ├── admin_user.go │           └── admin_user_group.go ├── cmd │   └── admin │       └── main.go ├── config │   └── app.yaml ├── go.mod ├── go.sum ├── pkg │   ├── auth │   │   └── jwtauth │   │       └── jwtauth.go │   ├── config │   │   ├── config.go │   │   ├── dev.go │   │   └── prod.go │   ├── contextx │   │   └── contextx.go │   ├── db │   │   ├── mysql.go │   │   └── redis.go │   ├── errors │   │   ├── error.go │   │   └── response.go │   ├── httpx │   │   └── httpx.go │   ├── json_serializer │   │   └── json_serializer.go │   ├── model │   │   ├── admin_group.go │   │   ├── admin_group_rule.go │   │   ├── admin_rule.go │   │   ├── admin_rule_action.go │   │   ├── admin_user.go │   │   ├── admin_user_group.go │   │   └── supplier.go │   ├── pagination │   │   ├── pagination.go │   │   └── schema.go │   ├── util │   │   ├── hash │   │   │   ├── md5.go │   │   │   ├── password.go │   │   │   └── sha1.go │   │   └── util.go │   └── validate │       ├── exist.go │       ├── is.go │       ├── regex.go │       ├── unique.go │       └── validate.go └── readme.MD.
├── app/
│   └── admin/
│       ├── handler/
│       │   ├── admin_group.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── public.go
│       ├── injector/
│       │   ├── app.go*
│       │   ├── injector.go*
│       │   ├── wire.go*
│       │   └── wire_gen.go
│       ├── repository/
│       │   ├── admin_group.go
│       │   ├── admin_group_rule.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── admin_user_group.go
│       ├── router/
│       │   ├── api.go
│       │   ├── middleware/
│       │   └── route.go
│       ├── schema/
│       │   ├── admin_group.go
│       │   ├── admin_menu.go
│       │   ├── admin_menu_action.go
│       │   ├── admin_user.go
│       │   └── common.go
│       └── service/
│           ├── admin_group.go
│           ├── admin_group_rule.go
│           ├── admin_rule.go
│           ├── admin_rule_action.go
│           ├── admin_user.go
│           └── admin_user_group.go
├── cmd/
│   └── admin/
│       └── main.go
├── config/
│   └── app.yaml
├── go.mod
├── go.sum
├── pkg/
│   ├── auth/
│   │   └── jwtauth/
│   │       └── jwtauth.go
│   ├── config/
│   │   ├── config.go
│   │   ├── dev.go
│   │   └── prod.go
│   ├── contextx/
│   │   └── contextx.go
│   ├── db/
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── errors/
│   │   ├── error.go*
│   │   └── response.go*
│   ├── httpx/
│   │   └── httpx.go*
│   ├── json_serializer/
│   │   └── json_serializer.go
│   ├── model/
│   │   ├── admin_group.go
│   │   ├── admin_group_rule.go
│   │   ├── admin_rule.go
│   │   ├── admin_rule_action.go
│   │   ├── admin_user.go
│   │   ├── admin_user_group.go
│   │   └── supplier.go
│   ├── pagination/
│   │   ├── pagination.go
│   │   └── schema.go
│   ├── util/
│   │   ├── hash/
│   │   │   ├── md5.go
│   │   │   ├── password.go
│   │   │   └── sha1.go
│   │   └── util.go
│   └── validate/
│       ├── exist.go
│       ├── is.go
│       ├── regex.go
│       ├── unique.go
│       └── validate.go
└── readme.MD

26 directories, 62 files
[01;34m.[00m
├── [01;34mapp[00m
│   └── [01;34madmin[00m
│       ├── [01;34mhandler[00m
│       │   ├── admin_group.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── public.go
│       ├── [01;34minjector[00m
│       │   ├── [01;32mapp.go[00m
│       │   ├── [01;32minjector.go[00m
│       │   ├── [01;32mwire.go[00m
│       │   └── wire_gen.go
│       ├── [01;34mrepository[00m
│       │   ├── admin_group.go
│       │   ├── admin_group_rule.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── admin_user_group.go
│       ├── [01;34mrouter[00m
│       │   ├── api.go
│       │   ├── [01;34mmiddleware[00m
│       │   └── route.go
│       ├── [01;34mschema[00m
│       │   ├── admin_group.go
│       │   ├── admin_menu.go
│       │   ├── admin_menu_action.go
│       │   ├── admin_user.go
│       │   └── common.go
│       └── [01;34mservice[00m
│           ├── admin_group.go
│           ├── admin_group_rule.go
│           ├── admin_rule.go
│           ├── admin_rule_action.go
│           ├── admin_user.go
│           └── admin_user_group.go
├── [01;34mcmd[00m
│   └── [01;34madmin[00m
│       └── main.go
├── [01;34mconfig[00m
│   └── app.yaml
├── go.mod
├── go.sum
├── [01;34mpkg[00m
│   ├── [01;34mauth[00m
│   │   └── [01;34mjwtauth[00m
│   │       └── jwtauth.go
│   ├── [01;34mconfig[00m
│   │   ├── config.go
│   │   ├── dev.go
│   │   └── prod.go
│   ├── [01;34mcontextx[00m
│   │   └── contextx.go
│   ├── [01;34mdb[00m
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── [01;34merrors[00m
│   │   ├── [01;32merror.go[00m
│   │   └── [01;32mresponse.go[00m
│   ├── [01;34mhttpx[00m
│   │   └── [01;32mhttpx.go[00m
│   ├── [01;34mjson_serializer[00m
│   │   └── json_serializer.go
│   ├── [01;34mmodel[00m
│   │   ├── admin_group.go
│   │   ├── admin_group_rule.go
│   │   ├── admin_rule.go
│   │   ├── admin_rule_action.go
│   │   ├── admin_user.go
│   │   ├── admin_user_group.go
│   │   └── supplier.go
│   ├── [01;34mpagination[00m
│   │   ├── pagination.go
│   │   └── schema.go
│   ├── [01;34mutil[00m
│   │   ├── [01;34mhash[00m
│   │   │   ├── md5.go
│   │   │   ├── password.go
│   │   │   └── sha1.go
│   │   └── util.go
│   └── [01;34mvalidate[00m
│       ├── exist.go
│       ├── is.go
│       ├── regex.go
│       ├── unique.go
│       └── validate.go
└── readme.MD

26 directories, 62 files
.
├── app
│   └── admin
│       ├── handler
│       │   ├── admin_group.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── public.go
│       ├── injector
│       │   ├── app.go
│       │   ├── injector.go
│       │   ├── wire.go
│       │   └── wire_gen.go
│       ├── repository
│       │   ├── admin_group.go
│       │   ├── admin_group_rule.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── admin_user_group.go
│       ├── router
│       │   ├── api.go
│       │   ├── middleware
│       │   └── route.go
│       ├── schema
│       │   ├── admin_group.go
│       │   ├── admin_menu.go
│       │   ├── admin_menu_action.go
│       │   ├── admin_user.go
│       │   └── common.go
│       └── service
│           ├── admin_group.go
│           ├── admin_group_rule.go
│           ├── admin_rule.go
│           ├── admin_rule_action.go
│           ├── admin_user.go
│           └── admin_user_group.go
├── cmd
│   └── admin
│       └── main.go
├── config
│   └── app.yaml
├── go.mod
├── go.sum
├── pkg
│   ├── auth
│   │   └── jwtauth
│   │       └── jwtauth.go
│   ├── config
│   │   ├── config.go
│   │   ├── dev.go
│   │   └── prod.go
│   ├── contextx
│   │   └── contextx.go
│   ├── db
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── errors
│   │   ├── error.go
│   │   └── response.go
│   ├── httpx
│   │   └── httpx.go
│   ├── json_serializer
│   │   └── json_serializer.go
│   ├── model
│   │   ├── admin_group.go
│   │   ├── admin_group_rule.go
│   │   ├── admin_rule.go
│   │   ├── admin_rule_action.go
│   │   ├── admin_user.go
│   │   ├── admin_user_group.go
│   │   └── supplier.go
│   ├── pagination
│   │   ├── pagination.go
│   │   └── schema.go
│   ├── util
│   │   ├── hash
│   │   │   ├── md5.go
│   │   │   ├── password.go
│   │   │   └── sha1.go
│   │   └── util.go
│   └── validate
│       ├── exist.go
│       ├── is.go
│       ├── regex.go
│       ├── unique.go
│       └── validate.go
└── readme.MD

26 directories, 62 files
.
├── app
│   └── admin
│       ├── handler
│       │   ├── admin_group.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── public.go
│       ├── injector
│       │   ├── app.go
│       │   ├── injector.go
│       │   ├── wire.go
│       │   └── wire_gen.go
│       ├── repository
│       │   ├── admin_group.go
│       │   ├── admin_group_rule.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── admin_user_group.go
│       ├── router
│       │   ├── api.go
│       │   ├── middleware
│       │   └── route.go
│       ├── schema
│       │   ├── admin_group.go
│       │   ├── admin_menu.go
│       │   ├── admin_menu_action.go
│       │   ├── admin_user.go
│       │   └── common.go
│       └── service
│           ├── admin_group.go
│           ├── admin_group_rule.go
│           ├── admin_rule.go
│           ├── admin_rule_action.go
│           ├── admin_user.go
│           └── admin_user_group.go
├── cmd
│   └── admin
│       └── main.go
├── config
│   └── app.yaml
├── go.mod
├── go.sum
├── pkg
│   ├── auth
│   │   └── jwtauth
│   │       └── jwtauth.go
│   ├── config
│   │   ├── config.go
│   │   ├── dev.go
│   │   └── prod.go
│   ├── contextx
│   │   └── contextx.go
│   ├── db
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── errors
│   │   ├── error.go
│   │   └── response.go
│   ├── httpx
│   │   └── httpx.go
│   ├── json_serializer
│   │   └── json_serializer.go
│   ├── model
│   │   ├── admin_group.go
│   │   ├── admin_group_rule.go
│   │   ├── admin_rule.go
│   │   ├── admin_rule_action.go
│   │   ├── admin_user.go
│   │   ├── admin_user_group.go
│   │   └── supplier.go
│   ├── pagination
│   │   ├── pagination.go
│   │   └── schema.go
│   ├── util
│   │   ├── hash
│   │   │   ├── md5.go
│   │   │   ├── password.go
│   │   │   └── sha1.go
│   │   └── util.go
│   └── validate
│       ├── exist.go
│       ├── is.go
│       ├── regex.go
│       ├── unique.go
│       └── validate.go
└── readme.MD

26 directories, 62 files
.
|-- app
|   `-- admin
|       |-- handler
|       |   |-- admin_group.go
|       |   |-- admin_rule.go
|       |   |-- admin_rule_action.go
|       |   |-- admin_user.go
|       |   `-- public.go
|       |-- injector
|       |   |-- app.go
|       |   |-- injector.go
|       |   |-- wire.go
|       |   `-- wire_gen.go
|       |-- repository
|       |   |-- admin_group.go
|       |   |-- admin_group_rule.go
|       |   |-- admin_rule.go
|       |   |-- admin_rule_action.go
|       |   |-- admin_user.go
|       |   `-- admin_user_group.go
|       |-- router
|       |   |-- api.go
|       |   |-- middleware
|       |   `-- route.go
|       |-- schema
|       |   |-- admin_group.go
|       |   |-- admin_menu.go
|       |   |-- admin_menu_action.go
|       |   |-- admin_user.go
|       |   `-- common.go
|       `-- service
|           |-- admin_group.go
|           |-- admin_group_rule.go
|           |-- admin_rule.go
|           |-- admin_rule_action.go
|           |-- admin_user.go
|           `-- admin_user_group.go
|-- cmd
|   `-- admin
|       `-- main.go
|-- config
|   `-- app.yaml
|-- go.mod
|-- go.sum
|-- pkg
|   |-- auth
|   |   `-- jwtauth
|   |       `-- jwtauth.go
|   |-- config
|   |   |-- config.go
|   |   |-- dev.go
|   |   `-- prod.go
|   |-- contextx
|   |   `-- contextx.go
|   |-- db
|   |   |-- mysql.go
|   |   `-- redis.go
|   |-- errors
|   |   |-- error.go
|   |   `-- response.go
|   |-- httpx
|   |   `-- httpx.go
|   |-- json_serializer
|   |   `-- json_serializer.go
|   |-- model
|   |   |-- admin_group.go
|   |   |-- admin_group_rule.go
|   |   |-- admin_rule.go
|   |   |-- admin_rule_action.go
|   |   |-- admin_user.go
|   |   |-- admin_user_group.go
|   |   `-- supplier.go
|   |-- pagination
|   |   |-- pagination.go
|   |   `-- schema.go
|   |-- util
|   |   |-- hash
|   |   |   |-- md5.go
|   |   |   |-- password.go
|   |   |   `-- sha1.go
|   |   `-- util.go
|   `-- validate
|       |-- exist.go
|       |-- is.go
|       |-- regex.go
|       |-- unique.go
|       `-- validate.go
`-- readme.MD

26 directories, 62 files
.
├── app
│   └── admin
│       ├── handler
│       │   ├── admin_group.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── public.go
│       ├── injector
│       │   ├── app.go
│       │   ├── injector.go
│       │   ├── wire.go
│       │   └── wire_gen.go
│       ├── repository
│       │   ├── admin_group.go
│       │   ├── admin_group_rule.go
│       │   ├── admin_rule.go
│       │   ├── admin_rule_action.go
│       │   ├── admin_user.go
│       │   └── admin_user_group.go
│       ├── router
│       │   ├── api.go
│       │   ├── middleware
│       │   └── route.go
│       ├── schema
│       │   ├── admin_group.go
│       │   ├── admin_menu.go
│       │   ├── admin_menu_action.go
│       │   ├── admin_user.go
│       │   └── common.go
│       └── service
│           ├── admin_group.go
│           ├── admin_group_rule.go
│           ├── admin_rule.go
│           ├── admin_rule_action.go
│           ├── admin_user.go
│           └── admin_user_group.go
├── cmd
│   └── admin
│       └── main.go
├── config
│   └── app.yaml
├── go.mod
├── go.sum
├── pkg
│   ├── auth
│   │   └── jwtauth
│   │       └── jwtauth.go
│   ├── config
│   │   ├── config.go
│   │   ├── dev.go
│   │   └── prod.go
│   ├── contextx
│   │   └── contextx.go
│   ├── db
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── errors
│   │   ├── error.go
│   │   └── response.go
│   ├── httpx
│   │   └── httpx.go
│   ├── json_serializer
│   │   └── json_serializer.go
│   ├── model
│   │   ├── admin_group.go
│   │   ├── admin_group_rule.go
│   │   ├── admin_rule.go
│   │   ├── admin_rule_action.go
│   │   ├── admin_user.go
│   │   ├── admin_user_group.go
│   │   └── supplier.go
│   ├── pagination
│   │   ├── pagination.go
│   │   └── schema.go
│   ├── util
│   │   ├── hash
│   │   │   ├── md5.go
│   │   │   ├── password.go
│   │   │   └── sha1.go
│   │   └── util.go
│   └── validate
│       ├── exist.go
│       ├── is.go
│       ├── regex.go
│       ├── unique.go
│       └── validate.go
└── readme.MD

26 directories, 62 files
