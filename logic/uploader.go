package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	region = "ap-northeast-1"
)

var s3Cli *s3.S3

/*
Credentials : 認証鍵情報
*/
type Credentials struct {
	AccessKey string `json:"aws_access_key_id"`
	SecretKey string `json:"aws_secret_access_key"`
}

/*
S3Upload : ファイルをS3にアップロードする
*/
func S3Upload(dirPath string, searchWord string, bucketName string, partSize int, credentialsPath string) {
	filePaths := getFilePaths(dirPath, searchWord)

	creData := getCredentials(credentialsPath)
	cre := credentials.NewStaticCredentials(
		creData.AccessKey,
		creData.SecretKey,
		"")
	s3Cli = s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String(region),
	})

	for _, filePath := range filePaths {
		key := filepath.Base(filePath)
		fmt.Println(key)

		file := getFile(filePath)
		err := S3MultipartUpload(key, bucketName, partSize, file)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprintf("Uploaded s3://%s", bucketName))

		if err := os.Remove(filePath); err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("delete file %s", filePath))
	}

}

/*
S3MultipartUpload : ファイルをS3にアップロードする
*/
func S3MultipartUpload(key string, bucket string, partSize int, file *os.File) (err error) {
	var uploader *s3manager.Uploader
	if partSize != 0 {
		uploader = s3manager.NewUploaderWithClient(s3Cli, func(u *s3manager.Uploader) {
			u.PartSize = int64(partSize * 1024 * 1024)
		})
	} else {
		uploader = s3manager.NewUploaderWithClient(s3Cli)
	}

	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	}

	_, err = uploader.Upload(uploadInput)

	return err
}

/*
指定したディレクトリ以下に存在する、指定した検索ワードに該当するファイルパスを取得する
*/
func getFilePaths(dirPath string, searchWord string) (filePaths []string) {
	list, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	for _, finfo := range list {
		if finfo.IsDir() || -1 == strings.Index(finfo.Name(), searchWord) {
			continue
		}

		fileName := finfo.Name()
		filePath := path.Join(dirPath, fileName)
		fmt.Println(filePath)

		filePaths = append(filePaths, filePath)
	}
	return
}

/*
ファイルパスからファイルを取得する
*/
func getFile(filePath string) (file *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	return
}

/*
認証鍵情報を取得する
*/
func getCredentials(credentialsPath string) (creData Credentials) {
	// ファイル読み込み
	file, err := readFile(credentialsPath)
	if err != nil {
		return
	}

	// ファイル(json文字列)からCredentialsデータに変換
	err = json.Unmarshal(file, &creData)
	if err != nil {
		fmt.Printf("json Unmarshal error: %v\n", err)
		os.Exit(1)
	}

	return
}

/*
指定されたファイルを読み込む
*/
func readFile(filePath string) (file []byte, err error) {
	// ファイル読み込み
	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return
	}

	return
}
