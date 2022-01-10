module github.com/yueqingkong/openApi

replace (
	github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
	golang.org/x/text => github.com/golang/text v0.3.2
)

require (
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.6
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/resty.v1 v1.12.0 // indirect
	xorm.io/builder v0.3.5
)

go 1.13
