## 下載包
go install github.com/zeromicro/go-zero/tools/goctl@latest

## 檢查版本  
goctl--version

## 創建API服務
goctl api new 服務名稱

## 生成服務代碼
goctl api go -api 服務名稱.api -dir . -style go_zero
- go_zero代表都會保持下滑線
## 啟動服務
go run 服務名稱.go -f 配置文件地址
1. 啟動user服務
    - go run user.go -f etc/user-api.yaml

## 須知
1. 修改代碼地方在user/internal/logic
2. 初始化資料庫在user/internal/svc