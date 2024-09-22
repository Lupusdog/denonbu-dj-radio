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
│           └── auth_middleware.go # 認証ミドルウェア
├── pkg/                         # 外部にも公開可能な再利用可能なライブラリ
│   └── logger/                  # ログ関連の共通ライブラリ
│       └── logger.go
├── configs/                     # 設定ファイル
│   └── config.yaml              # アプリケーション設定
└── go.mod                       # Goモジュールファイル
```
