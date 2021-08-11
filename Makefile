API=koukaten2021_api
DB=koukaten2021_mysql
API_SERVICE=api
DB_SERVICE=mysql
CAMERA_SERVER_DIR=python/
CAMERA_SERVER_SCRIPT=server.py
CAMERA_SERVER_REQUIREMENTS=requirements.txt

# サービスの構築
up-build:
	docker compose build

# コンテナ立ち上げ
## バックグラウンド
up:
	docker compose up -d
	@make up-camera
	
## フォアグラウンド
up-f:
	docker compose up &
	@make up-camera-f

# カメラサーバ立ち上げ
## バックグラウンド
up-camera:
	python3 $(CAMERA_SERVER_DIR)$(CAMERA_SERVER_SCRIPT) & > /dev/null 2>&1

## フォアグラウンド
up-camera-f:
	python3 $(CAMERA_SERVER_DIR)$(CAMERA_SERVER_SCRIPT)

# 停止系
## 全てのサービスを停止する
stop:
	docker-compose stop
	@make stop-camera

## コンテナを kill (強制停止)
kill:
	docker-compose kill
	@make stop-camera

## カメラサーバのみ停止する
stop-camera:
	ps | grep server.py | head -1 | awk '{print $$1}' | xargs kill

## コンテナ・ネットワークの停止と削除
down:
	docker compose down
	@make stop-camera

# サービスの再起動
restart:
	docker-compose restart
	@make stop-camera
	@make up-camera