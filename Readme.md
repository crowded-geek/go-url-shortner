[![GCI Badge](https://img.shields.io/badge/Google%20Code--in-JBoss%20Community-red?labelColor=2096F3)](https://gitter.im/JBossOutreach/gci)

# Go Url Shortner
## Prerequisites
- Have Go installed
- Have a working internet connection

## Run
- `go get https://github.com/speps/go-hashids`
- `go get https://github.com/gorilla/mux`
- `cd src`
- `go build main.go`
- `./main`

## Use:
### Create short url syntax
- localhost:8282/create/?longUrl=<url_you_wish_to_short>
- This outputs shortend url, use it to redirect to the original url

Google Code-in and the Google Code-in logo are trademarks of Google Inc.
