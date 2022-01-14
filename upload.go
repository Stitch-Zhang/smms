package smms

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type uploadResp struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		FileID    int    `json:"file_id"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Filename  string `json:"filename"`
		Storename string `json:"storename"`
		Size      int    `json:"size"`
		Path      string `json:"path"`
		Hash      string `json:"hash"`
		URL       string `json:"url"`
		Delete    string `json:"delete"`
		Page      string `json:"page"`
	} `json:"data"`
	RequestID string `json:"RequestId"`
}

// SMMSENDPOINT SM.MSP V2 API ENDPOINT
const (
	smmsEndpoint  = "https://sm.ms/api/v2"
	acceptFormats = "jpe|jpg|jpeg|gif|png|bmp|ico|svg|svgz|tif|tiff|ai|drw|pct|psp|xcf|psd|raw|webp"
)

// UploadImg upload image to SMMS
// if operation have completed return uploaded file url
// accept formats:
// jpe,jpg,jpeg,gif,png,bmp,ico,svg,svgz,tif,tiff,ai,drw,pct,psp,xcf,psd,raw,webp
func UploadImg(filePath, token string) (string, error) {
	if support := strings.ContainsAny(filepath.Ext(filePath), acceptFormats); !support {
		return "", errors.New("file format is unsuppoted : " + filepath.Ext(filePath))
	}
	fo, err := os.Open(filePath)

	if err == nil {
		defer fo.Close()

		formBuffer := bytes.NewBuffer(nil)
		form := multipart.NewWriter(formBuffer)
		form.WriteField("format", "json")
		fw, _ := form.CreateFormFile("smfile", filepath.Base(filePath))
		io.Copy(fw, fo)
		form.Close()

		req, err := http.NewRequest(http.MethodPost, smmsEndpoint+"/upload", formBuffer)
		if err != nil {
			return "", err
		}

		req.Header.Set("Content-Type", form.FormDataContentType())
		req.Header.Set("Authorization", token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}

		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		var respJSON uploadResp
		if err = json.Unmarshal(body, &respJSON); err == nil && respJSON.Success {
			return respJSON.Data.URL, nil
		}
	}
	return "", err
}
