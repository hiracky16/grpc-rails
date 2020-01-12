# gRPC-Rails

## 概要

gRPC を使って外部にある処理を Rails から呼び出すサンプル

## 使い方

```
# gRPC サーバーの立ち上げ方
$ go run grpc-server.go

# お試し実行
$ go run client.go

# Rails アプリケーションの立ち上げ方
$ cd server
$ bundle exec rails s
```

その後 http://localhost:3000/echo にて結果を確認
