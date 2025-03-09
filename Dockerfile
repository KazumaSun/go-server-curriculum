FROM golang:1.23-alpine

WORKDIR /app

COPY . .

# `main.go` が存在するか確認
RUN ls -l /app/server/cmd

# 依存関係を解決
RUN go mod tidy

# アプリケーションを実行
CMD ["go", "run", "./server/cmd/main.go"]