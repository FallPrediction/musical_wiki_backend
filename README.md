## 簡介
這是一個音樂劇的 wiki 專案後端的部分

[實體網站連結](https://musical_wiki.fall-prediction.net/)

還有很多功能還沒做，目前的 API 有：
- 演員列表
- 演員的出演作品
- 演員的大頭貼和圖片集（圖片儲存於 S3）

## 架構
![架構](https://musical_wiki.fall-prediction.net/assets/%E6%9E%B6%E6%A7%8B%E5%9C%96.b08244c7.jpg)

這是我剛開始學習 GO 的專案，使用 [Gin](https://github.com/gin-gonic/gin) 框架。因為習慣 PHP 的Laravel，所以這個專案的某些檔案目錄會參照 Laravel，如：
- handler(controller)：驗證輸入和 Response
- service：業務邏輯
- repository：跟 DB 溝通
- models：因為使用了 Gorm，該資料夾定義各 model 的 field，透過 orm 操作資料庫
- helper：全局用的小功能。並且使用 testing 寫測試
- request：各 API input 的驗證規則
- initialize：初始化部件，大部分都用 singleton，如 Redis, DB Connection 等
- utils：把常用的功能包裝成 struct，目前有上傳和 cache
- deployments
    - local：本地開發用的環境檔
    - codeDeploy：AWS Code Deploy 用的 scripts

### 前端
[Repository](https://github.com/FallPrediction/musical_wiki_frontend)
非前端工程師所以選了[Quasar](https://github.com/quasarframework/quasar)，是使用 Vue（專案使用 Vue3）的 UI 框架

### Infra
[Terraform Repository](https://github.com/FallPrediction/musical_wiki_terraform), [Ansible Repository](https://github.com/FallPrediction/musical_wiki_ansible)

此專案使用 Terraform + Ansible 架在 AWS 上

為了~~省錢~~簡單，使用託管的資料庫 [Supabase](https://supabase.com/)，也因為這樣所以服務全放在一台在 public subnet 的 server

使用 Gihub action 觸發 CodeDeploy 部署

## 本地開發
Docker Compose 啟動 GO, PostgreSQL, Redis 和 Nginx 服務

如何設定 GO 開發環境可以參考我的部落格文章[用 VSCode Debug GO 程式吧](https://fallprediction.github.io/blog/posts/vscode-debug-go/)

如果要部署到EC2，可以參考我的部落格文章[CodeDeploy 部署 GO APP 到 EC2](https://fallprediction.github.io/blog/posts/code-deploy/)

### 設定.env
#### 後端
設定 PG 的資料庫密碼、S3 bucket 等
Redis 的密碼需要 SHA256 以後修改 users.acl
開發時 GIN_MODE 設為`debug`，另外有設定 CORS，要配合前端使用時，要設定前端的 URL
#### 前端
前端環境可以在後端的 .env 指定打包好的目錄
前端的 .env 設定後端的 URL
如果不想打包，可以在前端目錄下開啟開發用 server（在這之前需要安裝 Quasar）
```
npm install
npm i -g @quasar/cli
quasar dev
```
