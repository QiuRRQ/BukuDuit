package aws

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type AWSS3 struct {
	AWSConfig aws.Config
	Bucket    string
	Directory string
}

func (awss3 AWSS3) UploadManager(fileToBeUploaded *multipart.FileHeader) (s3path string,fileName string, err error) {
	session, err := awsSession.NewSession(&awss3.AWSConfig)
	if err != nil {
		return s3path,fileName, err
	}

	file, err := fileToBeUploaded.Open()
	if err != nil {
		return s3path,fileName, err
	}
	size := fileToBeUploaded.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	fileName = bson.NewObjectId().Hex() + filepath.Ext(fileToBeUploaded.Filename)

	tempFile := awss3.Directory + "/" + fileName
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		Bucket:               aws.String(awss3.Bucket),
		ContentDisposition:   aws.String("attachment"),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		Key:                  aws.String(tempFile),
	})
	if err != nil {
		return s3path,fileName,err
	}

	s3path = os.Getenv("HDRIVE_S3_BASE_URL")+awss3.Directory

	return s3path,fileName, err
}
