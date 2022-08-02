module github.com/yueqingkong/openApi

replace (
	github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
	golang.org/x/text => github.com/golang/text v0.3.2
)

require (
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.6
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/resty.v1 v1.12.0 // indirect
	xorm.io/builder v0.3.5
)

go 1.13
