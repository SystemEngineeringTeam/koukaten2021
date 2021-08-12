CAMERA_DIR=python/
CAMERA_SERVER_SCRIPT=server.py
CAMERA_SERVER_REQUIREMENTS=requirements.txt

build: ## サービスの構築
	docker compose build
	pip install -r $(CAMERA_DIR)$(CAMERA_SERVER_REQUIREMENTS)

up: ## サービス立ち上げ(バックグラウンド)
	docker compose up -d
	@make up-camera
	
up-f: ## サービス立ち上げ(フォアグラウンド)
	docker compose up &
	@make up-camera-f

stop: ## サービスを停止
	docker-compose stop
	@make stop-camera

kill: ## サービスを強制停止
	docker-compose kill
	@make stop-camera

down: ## サービスの停止とコンテナの削除
	docker compose down
	@make stop-camera

restart: ## サービスの再起動
	docker-compose restart
	@make stop-camera
	@make up-camera

help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# カメラ関連
## カメラサーバ立ち上げ(バックグラウンド)
up-camera:
	python3 $(CAMERA_DIR)$(CAMERA_SERVER_SCRIPT) & > /dev/null 2>&1

## カメラサーバ立ち上げ(フォアグラウンド)
up-camera-f:
	python3 $(CAMERA_DIR)$(CAMERA_SERVER_SCRIPT)

## カメラサーバ停止
stop-camera: 
	ps | grep server.py | head -1 | awk '{print $$1}' | xargs kill