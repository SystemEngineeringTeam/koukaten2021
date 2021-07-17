API=koukaten2021_api
DB=koukaten2021_mysql
API_SERVICE=api
DB_SERVICE=mysql
CAMERA_SERVER_DIR=python/
CAMERA_SERVER_SCRIPT=server.py
CAMERA_SERVER_REQUIREMENTS=requirements.txt

# サービスの構築及び立ち上げ
## バックグラウンド
up-build:
	docker compose up -d --build
	@make up-camera

up-api:
	docker compose up -d --build $(API_SERVICE)
	@make up-camera

up-database:
	docker compose up -d --build $(DB_SERVICE)
	@make up-camera

## フォアグラウンド
up-build-f:
	docker compose up --build &
	@make up-camera-f

up-api-f:
	docker compose up --build $(API_SERVICE) &
	@make up-camera-f

up-database-f:
	docker compose up -d --build $(DB_SERVICE) &
	@make up-camera-f

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

## ボリュームも消す
down-v:
	docker compose down -v
	@make stop-camera

## Composeファイルで定義していないサービス用のコンテナも削除
down-ro:
	docker compose down --remove-orphans
	@make stop-camera

# サービスの再起動
restart:
	docker-compose restart
	@make stop-camera
	@make up-camera

# その他
## データベースコンテナに入る
sh-database:
	docker exec -it $(DB) bash

## APIコンテナに入る
sh-api:
	docker exec -it $(API) sh

## 全てのサービスのログを出力する
logs:
	docker compose logs

## コンテナ一覧
ps:
	docker compose ps

## イメージの一覧を表示
images:
	docker compose images

## 全て消し
destroy:
	docker compose down --volumes --remove-orphans
	@make stop-camera

## イメージも消す
destroy-with-image:
	docker compose down --rmi all --volumes --remove-orphans
	@make stop-camera