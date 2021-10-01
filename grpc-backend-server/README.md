# 購買バックエンド

## 要求仕様

- Golang 1.14 以降
- MySQL 5.7系

## ビルド

```
# proto-codegen
git submodule update --remote -i -f
./protoc.sh

# gomock generate
./mockgen.sh

# go build
go build

# go run
go run main.go
```

## テスト

```
# server_test実行
cd interfaces
go test -v -run TestServer
# テストケースを一つのみ実行する
go test -v -run TestServer/{*}/{**}
>例：go test -v -run TestServer/OrderType/GetOrderType

# repository_test実行
cd infra/repository
go test -v -run テスト用リポジトリ名
>例：go test -v -run TestOrderTypeRepository
```
