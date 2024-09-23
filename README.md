# denonbu-dj-radio-api
This is a API code of "Denonbu DJ Radio" which is fan-made Denonbu Content

全体構成(初期イメージ)
```
myapp/
├── cmd/                         # アプリケーションのエントリポイント
│   └── myapp/
│       └── main.go              # APIサーバーのエントリポイント
├── internal/                    # 内部パッケージ（外部に公開されない）
│   ├── domain/                  # ドメイン層（ビジネスロジック）
│   │   ├── user/                # ドメインモデル「User」
│   │   │   ├── user.go          # ドメインエンティティ（User）
│   │   │   ├── repository.go    # リポジトリインターフェース
│   │   └── order/               # ドメインモデル「Order」
│   │       ├── order.go         # ドメインエンティティ（Order）
│   │       └── repository.go    # リポジトリインターフェース
│   ├── application/             # アプリケーション層（ユースケース）
│   │   ├── user/                # Userに関連するユースケース
│   │   │   └── service.go       # ユースケースサービス
│   │   ├── order/               # Orderに関連するユースケース
│   │       └── service.go       # ユースケースサービス
│   ├── infrastructure/          # インフラ層（データベースや外部サービス）
│   │   ├── persistence/         # 永続化層（DBリポジトリ実装）
│   │   │   ├── user/            # Userリポジトリ実装
│   │   │   │   └── postgres_user_repository.go  # PostgreSQLの実装
│   │   │   └── order/           # Orderリポジトリ実装
│   │   │   └── postgres_order_repository.go     # PostgreSQLの実装
│   │   └── external_services/   # 外部サービスとの通信
│   │       └── payment_service.go # 支払いサービスの実装
│   └── interfaces/              # インターフェース層（HTTP, gRPCなど）
│       ├── http/                # HTTPハンドラーとルーティング
│       │   ├── user_handler.go  # User用のHTTPハンドラー
│       │   ├── order_handler.go # Order用のHTTPハンドラー
│       │   └── router.go        # ルーティング設定
│       └── middleware/          # ミドルウェア
│           ├── auth_middleware.go      # 認証のミドルウェア
│           ├── logging_middleware.go   # ログ出力のミドルウェア
│           └── recovery_middleware.go  # パニックリカバリーのミドルウェア
├── pkg/                         # 外部にも公開可能な再利用可能なライブラリ
│   └── logger/                  # ログ関連の共通ライブラリ
│       └── logger.go
├── configs/                     # 設定ファイル
│   └── config.yaml              # アプリケーション設定
└── go.mod                       # Goモジュールファイル
```

## ステータスコード、メッセージ一覧

| 独自成功コード | メッセージ           | 説明                                                              |
|----------------|----------------------|-------------------------------------------------------------------|
| **1000**       | OK              | リクエストが正常に処理されました。標準的な成功レスポンスです。        |
| **1001**       | Created              | 新しいリソースが正常に作成されました。                             |
| **1002**       | Accepted             | リクエストが受理されましたが、処理はまだ完了していません。              |
| **1003**       | No Content           | リクエストは正常に処理されましたが、レスポンスボディは空です。          |

---

| 独自エラーコード | メッセージ                     | 説明                                                              |
|------------------|--------------------------------|-------------------------------------------------------------------|
| **2000**         | Invalid Request                | リクエスト自体が不正です。                                         |
| **2001**         | Missing Required Parameter     | 必須パラメータがリクエストに含まれていない場合のエラーです。          |
| **2002**         | Invalid Parameter Value        | パラメータの値が無効な場合です。                                   |
| **2003**         | Unauthorized Access            | 認証が必要なリソースに対して、認証されていないユーザーがアクセスしようとした場合のエラーです。 |
| **2004**         | Forbidden                      | 認証されているが、リソースに対するアクセス権限がない場合のエラーです。   |
| **2005**         | Resource Not Found             | リクエストされたリソースが存在しない場合のエラーです。                |
| **2006**         | Conflict                       | リクエストが既存のリソースと競合している場合のエラーです。            |
| **2007**         | Unprocessable Entity           | リクエストは受理されたが、データが処理できない場合のエラーです。        |
| **2008**         | Too Many Requests              | クライアントが短期間に過剰なリクエストを送信した場合のエラーです。      |
| **2009**         | Internal Server Error          | サーバー内部で予期しないエラーが発生した場合のエラーです。             |
| **2010**         | Service Unavailable            | サーバーが過負荷またはメンテナンス中でリクエストを処理できない場合のエラーです。 |
| **2011**         | Gateway Timeout                | 外部サービスからの応答が遅延またはタイムアウトした場合のエラーです。     |

---

## 成功時のレスポンス例

**正常なリクエスト**

```json
{
    "status": "success",
    "status_code": 1000,
    "status_msg": "OK",
    "data": {
        "id": 123,
        "name": "John Doe"
    }
}

{
    "status": "success",
    "status_code": 1001,
    "status_msg": "Created",
    "data": {
        "id": 123,
        "name": "New Resource",
        "created_at": "2024-09-23T10:20:00Z"
    }
}

{
    "status": "success",
    "status_code": 1002,
    "status_msg": "Accepted",
    "data": {
        "detail": "Your request has been accepted and is being processed."
    }
}

{
    "status": "success",
    "status_code": 1003,
    "status_msg": "No Content",
    "data": {}
}
```

 **不正なリクエスト**
```json
{
    "status": "error",
    "status_code": 2000,
    "status_msg": "Invalid Parameter Value",
}

//  エラーは以降、同様の形式
```
