# GraphQL を実装する

https://zenn.dev/hsaki/books/golang-graphql

## 実行

```
go run server.go
```

## テスト

```
go test ./...
```

## MVP

- [x] 環境準備
  - [x] dbを準備
  - [x] gqlgenによるコード生成 
- [ ] 機能を作成
  - [ ] ノートの作成機能
  - [ ] ノートの閲覧機能
  - [ ] ノートの更新機能
  - [ ] ノートの削除機能
  - [ ] ユーザーの作成機能
  - [ ] ユーザーの閲覧機能
  - [ ] ユーザーの更新機能
  - [ ] ユーザーの削除機能
  - [ ] ユーザーの認証機能
  - [ ] ユーザーの認可機能
- [ ] テストを作成
  - [ ] E2Eテストを作成
  - [ ] ユニットテストを作成

## 参考

- gqlgenによるコード生成

モデル定義は、schema.graphlsに定義する

```bash
# 初回生成
$ gqlgen init
# コード生成
$ gqlgen generate
```
