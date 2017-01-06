package main

import (
	"flag"
	"path"

	"github.com/scarviz/goS3Uploader/logic"
)

const credentialsPath = "./.aws/credentials"

func main() {
	dirPath, searchWord, bucketName, partSize, credir := flagParse()
	if credir == "" {
		credir = credentialsPath
	} else {
		credir = path.Join(credir, credentialsPath)
	}
	logic.S3Upload(dirPath, searchWord, bucketName, partSize, credir)
}

/*
パラメータのパース処理
*/
func flagParse() (dirPath string, searchWord string, bucketName string, partSize int, credir string) {
	pathFlag := flag.String("dirpath", "", "uploade file dir path")
	searchWordFlag := flag.String("searchword", "", "upload file search word")
	bucketFlag := flag.String("bucket", "", "bucket name")
	partSizeFlag := flag.Int("partsize", 0, "part size (MB)")
	credirFlg := flag.String("credir", "", "dir path where .aws dir is stored")
	flag.Parse()

	dirPath = *pathFlag
	searchWord = *searchWordFlag
	bucketName = *bucketFlag
	partSize = *partSizeFlag
	credir = *credirFlg

	return
}
