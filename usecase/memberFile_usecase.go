package usecase

import (
	"bukuduit-go/usecase/viewmodel"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

type MemberFileUseCase struct {
	*UcContract
}

func (uc MemberFileUseCase) Add(input echo.Context) (res viewmodel.BookFileVM, err error) {

	name := input.FormValue("name")
	uploadedFile, errr := input.FormFile("file")
	if errr != nil {
		return res, errr
	}
	src, err := uploadedFile.Open()
	if err != nil {
		return res, err
	}

	defer src.Close()
	dir, Myerror := os.Getwd()

	if Myerror != nil {
		return res, Myerror
	}

	filename := uploadedFile.Filename
	newName := name + ".png"
	os.Rename(filename, newName)
	fileLocation := filepath.Join(dir, "member_file", newName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return res, err
	}

	defer targetFile.Close()

	if _, errrr := io.Copy(targetFile, src); err != nil {
		return res, errrr
	}

	body := viewmodel.BookFileVM{
		Name: filename,
		Path: fileLocation,
	}

	return body, nil
}
