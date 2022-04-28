module github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway

go 1.17

replace github.com/XWS-2022-Tim12/Dislinkt/back/common => ../common
replace github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_service => ../common/proto/user_service
replace github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/startup/config => ../api_gateway/startup/config
replace github.com/XWS-2022-Tim12/Dislinkt/back/api_gateway/startup => ../api_gateway/startup

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0
	github.com/XWS-2022-Tim12/Dislinkt/back/common v1.0.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
