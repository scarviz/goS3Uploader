goS3Uploader
---
S3にファイルをアップロードする。  

|対象|バージョン|
|:--|:--|
|go|1.7.4|

## 使用ライブラリ
aws-sdk-goライブラリを使用している。下記を実行してライブラリを取得する。  

```
go get -u github.com/aws/aws-sdk-go
```

## goS3Uploaderの使用方法
下記(カレントディレクトリ以下)に認証鍵ファイル(credentials)を配置する。  

```
./.aws/credentials
```

または、実行時にcredirパラメータを指定し、.awsディレクトリ以下に配置する。    

認証鍵ファイルは、下記のようなJSONにする。  

```json
{"aws_access_key_id":"AKID1234567890", "aws_secret_access_key":"MY-SECRET-KEY"}
```

パラメータを指定して  

```
./goS3Uploader (各パラメータ)
```

を実行する。  

|パラメータ名|型|値|
|:--|:--|:--|
|dirpath|string|アップロード対象のファイルが格納されているディレクトリパス|
|searchword|string|アップロード対象のファイルを検索するためのワード|
|bucket|string|アップロード先のS3のBucket名|
|partsize|int|分割ファイルサイズ(MB)。ファイルサイズが大きい場合(100MBより大きい)に指定する|
|credir|string|認証鍵ファイル(credentials)が格納されている .aws ディレクトリの親ディレクトリ|
