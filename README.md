# sqs-message-sender

Amazon SQSにメッセージを送るサンプルプログラム

## 簡単な説明

Webアプリ上でDBに登録、更新、削除した時にAmazon SQSにメッセージを送ります。

## 必要要件

実行環境

- go 1.15.6 windows/amd64
- MariaDB 10.5.8

Amazon SQSを使用するためアカウント登録が必要です。

## 使い方

Webアプリ上で登録、更新、削除を行ってください。


## インストール

### データベース

MariaDBをインストールしてDBを作成してください。  
（テーブル`items`は初回起動時に自動で作成されます。）

`database/database.go` にDBの設定を入力してください。

- USER ： DBのユーザ名 例）"test_user"
- PASS ： 上記ユーザのパスワード 例）"test_pass"
- HOST ： DBのホスト名もしくはIP 例）"localhost"
- PORT ： DBのポート番号 例）"3306"
- DBNAME ： データベース名 例）"exampledb"

### Amazon SQS

AWSにログインし、SQSサービスのページに遷移します。  
  ⇒開発時にはリージョンは`アジアパシフィック（東京） ap-northeast-1`を選択しました。

キューを作成します。  
  ⇒開発時には`名前`のみ入力しそれ以外はデフォルトで作成しました。

キューの一覧にある作成したキューのリンクをクリックし、詳細情報からURLをコピーします。

`message/message.go` にSQSの設定を入力してください。

- AWS_REGION ： リージョン 例）"ap-northeast-1"
- QUEUE_URL  ： 上記で取得したURL 例）"https://sqs.ap-northeast-1.amazonaws.com/123456789012/TestQueue"
- DO_NOTHING ： 通常はfalse 例）false  
　⇒メッセージを送信したくない場合はtrueにしてください。

### AWS認証情報設定

AWSの`アクセスキー ID` と `シークレットアクセスキー` を取得します。
（AWSユーザ登録時に届くメールなど）

ここではWindwosのコマンドプロンプトから実行する時に一時的に環境変数に登録するコマンドを載せます。
詳細はAWSのユーザガイドを見てください。

https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-configure-envvars.html

コマンドプロンプトで以下のコマンドを実行してください。
この書き方の場合、コマンドプロンプトを起動するたびに実行する必要があります。
```
set AWS_ACCESS_KEY_ID=XXXXXXXXXXX
set AWS_SECRET_ACCESS_KEY=XXXXXXXXXXXXX
```

### go言語のライブラリ

以下のコマンドで必要なライブラリをインストールしてください。

```
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
go get github.com/gin-gonic/gin
go get github.com/aws/aws-sdk-go
```

### 起動と実行

`sqs-message-sender`ディレクトリに移動し、以下のコマンドで起動します。
```
go run main.go
```

localhostの場合、ブラウザで以下のURLにアクセスしてください。
```
http://localhost/
```

ポートを変更したい場合、`main.go`の`PORT`を変更してください。

## その他

* 実行したあと、Amazon SQS上でメッセージを確認したい場合、キューの詳細画面から「メッセージを送受信」をクリック⇒「メッセージをポーリング」で確認できます。
* 登録、更新、削除 を各々1回実行するたびにSQSにメッセージを発行します。



