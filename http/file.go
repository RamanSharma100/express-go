package http

import (
	"io"
	"os"
	"path/filepath"
)

type FileObject struct {
	Fieldname    string
	Originalname string
	Encoding     string
	Mimetype     string
	Size         int64
	Destination  string
	Filename     string
	Path         string
	Buffer       []byte // TODO: Implement buffer handling
}

func (ctx *Context) saveUploadedFiles(uploadDir string) (map[string][]FileObject, error) {
	req := ctx.Request.r

	if err := req.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	files := make(map[string][]FileObject)

	for field, allFiles := range req.MultipartForm.File {
		for _, fileHeader := range allFiles {
			file, err := fileHeader.Open()
			if err != nil {
				return nil, err
			}
			defer file.Close()

			filePath := filepath.Join(uploadDir, fileHeader.Filename)
			dstFile, err := os.Create(filePath)
			if err != nil {
				return nil, err
			}
			defer dstFile.Close()

			if _, err := io.Copy(dstFile, file); err != nil {
				return nil, err
			}

			files[field] = append(files[field], FileObject{
				Fieldname:    field,
				Originalname: fileHeader.Filename,
				Encoding:     fileHeader.Header.Get("Content-Transfer-Encoding"),
				Mimetype:     fileHeader.Header.Get("Content-Type"),
				Size:         fileHeader.Size,
				Destination:  uploadDir,
				Filename:     fileHeader.Filename,
				Path:         filePath,
				Buffer:       nil, // TODO: Implement buffer handling
			})
		}
	}

	ctx.Request.AdditionalFields["files"] = files
	return files, nil
}

func (ctx *Context) GetUploadedFiles() (map[string][]FileObject, error) {
	if files, ok := ctx.Request.AdditionalFields["files"].(map[string][]FileObject); ok {
		return files, nil
	}
	return nil, os.ErrNotExist

}
