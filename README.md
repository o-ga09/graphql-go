# GraphQL を実装する

[Goで学ぶGraphQLサーバーサイド入門](https://zenn.dev/hsaki/books/golang-graphql)

## 実行

```
$ go run cmd/main.go
```

## テスト

```
go test ./...
```

## MVP

- [x] 環境準備
  - [x] dbを準備
  - [x] gqlgenによるコード生成 
- [x] 機能を作成
  - [x] ノートの作成機能
  - [x] ノートの閲覧機能
  - [x] ノートの更新機能
  - [x] ノートの削除機能
  - [x] ユーザーの作成機能
  - [x] ユーザーの閲覧機能
  - [x] ユーザーの更新機能
  - [x] ユーザーの削除機能
- [ ] テストを作成
  - [ ] E2Eテストを作成
  - [x] ユニットテストを作成

## 参考

- gqlgenによるコード生成

モデル定義は、schema.graphlsに定義する

```bash
# 初回生成
$ gqlgen init
# コード生成
$ gqlgen generate
```
