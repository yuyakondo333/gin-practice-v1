# Go Backend API

Go言語とGinフレームワークを使用したバックエンドAPIです。

## 技術スタック

- **言語**: Go 1.25.0
- **フレームワーク**: Gin
- **データベース**: PostgreSQL
- **ORM**: GORM
- **コンテナ**: Docker & Docker Compose

## 前提条件

- Go 1.25.0以上
- Docker & Docker Compose
- Air（ホットリロード用）
- Make（オプション）

## セットアップ手順

### 1. リポジトリのクローン

```bash
git clone <repository-url>
cd <project-directory>
```

### 3. Dockerコンテナの起動

```bash
docker-compose up -d
```

### 4. データベースの初期化

#### マイグレーションの実行
```bash
make migrate
```

#### シードデータの投入
```bash
make seed
```

#### データベースの完全リセット（マイグレーション + シードデータ）
```bash
make reset
```

### 5. アプリケーションの起動

#### 通常の起動
```bash
go run main.go
```

#### Airを使用したホットリロード起動（推奨）
```bash
air
```

#### ビルドして実行
```bash
go build -o tmp/main
./tmp/main
```

## 利用可能なコマンド

### Makeコマンド

- `make help` - 利用可能なコマンドを表示
- `make migrate-up` - データベースマイグレーションを実行
- `make migrate-down` - マイグレーションを適用解除
- `make seed` - シードデータを投入
- `make reset` - データベースを完全にリセット

### 手動でのデータベース操作

```bash
# マイグレーションの手動実行（UP）
for file in db/migrations/*_up.sql; do
  docker exec -i <container-name> psql -U postgres -d myapp < $file
done

# マイグレーションの手動実行（DOWN）
for file in db/migrations/*_down.sql; do
  docker exec -i <container-name> psql -U postgres -d myapp < $file
done

# シードデータの手動投入
for file in db/seeds/*.sql; do
  docker exec -i <container-name> psql -U postgres -d myapp < $file
done
```

## API エンドポイント

### ヘルスチェック

- `GET /api/health` - アプリケーションのヘルスチェック
- `GET /api/dbhealth` - データベース接続のヘルスチェック

## プロジェクト構造

```
<project-directory>/
├── config/          # 設定ファイル
├── db/              # データベース関連
│   ├── migrations/  # マイグレーションファイル
│   └── seeds/       # シードデータ
├── dto/             # データ転送オブジェクト
├── handler/         # HTTPハンドラー
├── infra/           # インフラストラクチャ層
├── models/          # データモデル
├── repositories/    # リポジトリ層
├── services/        # ビジネスロジック層
├── main.go          # エントリーポイント
├── docker-compose.yml
├── Makefile
└── README.md
```

## データベーススキーマ

### テーブル構成

- **testmigrates** - テスト用テーブル

## 開発

### 新しいマイグレーションの追加

1. `db/migrations/`ディレクトリに新しいSQLファイルを作成
2. ファイル名は`XXX_description_up.sql`と`XXX_description_down.sql`の形式（XXXは連番）
3. `make migrate-up`でマイグレーションを実行
4. `make migrate-down`でマイグレーションを適用解除

### 新しいシードデータの追加

1. `db/seeds/`ディレクトリに新しいSQLファイルを作成
2. ファイル名は`XXX_description.sql`の形式（XXXは連番）
3. `make seed`でシードデータを投入

## Airの設定

### Airのインストール
```bash
go install github.com/cosmtrek/air@latest
```

## トラブルシューティング

### データベース接続エラー

1. Dockerコンテナが起動しているか確認
```bash
docker-compose ps
```

2. データベースのログを確認
```bash
docker-compose logs db
```

### ポート競合エラー

デフォルトポート（8080）が使用中の場合は、`.env`ファイルで`SERVER_PORT`を変更してください。

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。
