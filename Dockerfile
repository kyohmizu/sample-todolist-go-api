FROM golang:latest

# 作業ディレクトリを指定
WORKDIR /app

# go.mod、go.sum ファイルをコピー
COPY go.mod go.sum ./

# 依存関係のあるパッケージをダウンロード
RUN go mod download

# APIのソースコードをコピー
COPY main.go .

# ソースコードをビルド
RUN go build .

# APIの使用するポートを公開
EXPOSE 9999

# デフォルトの実行コマンド
CMD ["./go-api"]

