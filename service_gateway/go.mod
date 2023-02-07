module api-gateway

go 1.16

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-contrib/sessions v0.0.4
	github.com/gin-gonic/gin v1.7.7
	github.com/micro/go-micro/v2 v2.9.1
	service_common v0.0.0

)

replace service_common => ../service_common
