BACK=koukaten2021_api
DB=koukaten_mysql
BACK_SERVICE = backend
DB_SERVICE = mysql

## サービスの構築または再構築
up-build:
	docker compose up -d --build

up-back:
	docker compose up -d --build $(BACK_SERVICE)

up-database:
	docker compose up -d --build $(DB_SERVICE)

#コンテナ立ち上げ（バックグランド）
run:
	docker compose up -d

start:
	docker compose start

sh-database:
	docker exec -it $(DB) bash

# コンテナ・ネットワーク・イメージ・ボリュームの停止と削除
down:
	docker compose down

#ボリュームも消す
down-v:
	docker compose down -v

down-ro:
	docker compose down --remove-orphans

#全てのサービスのログを出力する
logs:
	docker compose logs

#コンテナ一覧
ps:
	docker compose ps

#イメージの一覧を表示
images:
	docker compose images

#全てのサービスを停止する
stop:
	docker compose stop

# コンテナを kill (強制停止)
kill:
	docker compose kill

restart-dr:
	@make down
	@make run

#サービスの再起動
restart:
	docker compose restart

exec-back:
	docker compose exec $(BACK_SERVICE) bash

#全て消し
destroy:
	docker compose down --rmi all --volumes --remove-orphans

destroy-volumes:
	docker compose down --volumes --remove-orphans
