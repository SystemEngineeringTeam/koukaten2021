# シス研人数出る蔵
シス研人数出る蔵は，カメラで撮った部屋の写真から人体を検出し，部屋に滞在している人数を教えてくれるシステムです．

## Requirements
|言語/FW|Version|
|---|---|
|docker|20.10.2|
|docker-compose|1.27.4|
|python|3.9.6|
|node|v16.5.0|
|yarn|1.22.4|
  
## Usage
初回実行時には以下のコマンドを実行します: 
```
git clone https://github.com/SystemEngineeringTeam/koukaten2021.git
cd koukaten2021
```
### Backend
バックエンドの初回実行時には以下のコマンドを実行します: 
```
make build
```
以下のコマンドでバックエンドのシステムが起動します:
```
make up
```
以下のコマンドでバックエンドのシステムを終了します: 
```
make stop
```
コンテナをリセットしたい時は以下のコマンドを実行してください: 
```
make down
```
その他のコマンドを確認するには`make help`を実行してください． 
### Frontend
基本的に，フロントエンドに関する作業は`nuxt`ディレクトリで行います: 
```
cd nuxt
```
フロントエンドの初回実行時には以下のコマンドを実行します: 
```
yarn
```
以下のコマンドでフロントエンドのシステムが起動します: 
```
yarn dev
```
プロセスを終了するまで，`http://localhost:3000`にアクセスすることによってウェブページを確認することができます．
## Licence
<a href="https://github.com/SystemEngineeringTeam/koukaten2021/blob/master/LICENSE">GNU General Public License v3.0</a>