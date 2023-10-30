# yuyuancha-tool

芋圓茶常用小工具。

## 開發環境

- Docker compose: 3.5
- 後端：
  - Golang: 1.20
  - MySQL: 8.0
- 前端：
  - Node: latest
  - Vue vite

## 開啟方式

1. 複製根目錄、前端、後端之 `.env.example`。
    ```=shell
   cp .env.example .env
   cp ./backend/.env.example ./backend/.env 
   cp ./frontend/.env.example ./frontend/.env 
   ```
2. 開啟 `Docker` 服務。
   ```shell
   docker-compose up -d
   ```
3. 開啟後端服務。
   ```shell
   docker-compose exec backend go run main.go
   ```
4. 開啟網頁 `http://localhost:2001`。
