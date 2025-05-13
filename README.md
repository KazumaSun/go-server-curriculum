# go-server-curriculum
Goのサーバーサイド用のカリキュラム

## カリキュラム概要
このカリキュラムでは、Goを使用してシンプルなサーバーサイドアプリケーションを構築する方法を学びます。以下の内容を含みます：
- RESTful APIの設計と実装
- 商品（Product）に対するCRUD操作
- Echoフレームワークを使用したルーティング
- ユースケース層とリポジトリ層の分離
- テストの実装

## テスト用のcurlコマンド例

### 商品一覧を取得
```bash
curl -X GET http://localhost:8080/products
```

### 商品をIDで取得
```bash
curl -X GET http://localhost:8080/products/1
```

### 新しい商品を作成
```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{
  "name": "Sample Product",
  "price": 1000
}'
```

### 商品を更新
```bash
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Updated Product",
  "price": 1500
}'
```

### 商品を削除
```bash
curl -X DELETE http://localhost:8080/products/1
```

## 実行方法
1. Dockerイメージをビルドします。
   ```bash
   docker build -t go-server-curriculum .
   ```

2. コンテナを起動します。
   ```bash
   docker run -p 8080:8080 go-server-curriculum
   ```

3. 上記のcurlコマンドを使用してAPIをテストします。

## 最終課題: お客ごとの合計金額を算出するAPI

### 概要
この課題では、カスタマー情報（名前、席番号）を管理し、各お客が注文した商品の合計金額を算出するAPIを実装します。

### エンドポイント
- **GET /customers/:id/total**  
  指定したカスタマーIDに基づき、注文の合計金額を返します。

### 実行例
```bash
curl -X GET http://localhost:8080/customers/1/total
```

### レスポンス例
```json
{
  "customer": {
    "id": 1,
    "name": "John Doe",
    "seat": "A1",
    "orders": [
      {
        "id": 1,
        "product_id": 1,
        "quantity": 2
      }
    ]
  },
  "total": 200
}
```