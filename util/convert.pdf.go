package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	tokenURL   = "https://api.aspose.cloud/connect/token"
	baseAPIURL = "https://api.aspose.cloud/v4.0/words"
)

type AsposeClient struct {
	ClientID     string
	ClientSecret string
	HTTPClient   *http.Client
}

func NewAsposeClient(clientID, clientSecret string) *AsposeClient {
	return &AsposeClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (a *AsposeClient) GetToken(ctx context.Context) (string, error) {

	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", a.ClientID)
	form.Set("client_secret", a.ClientSecret)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		tokenURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error obteniendo token: %s", string(body))
	}

	var tr tokenResponse

	if err := json.Unmarshal(body, &tr); err != nil {
		return "", err
	}

	if tr.AccessToken == "" {
		return "", fmt.Errorf("token vacío")
	}

	return tr.AccessToken, nil
}

func (a *AsposeClient) UploadFile(ctx context.Context, token, localFile, remoteName string) error {

	file, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	u := fmt.Sprintf("%s/storage/file/%s", baseAPIURL, url.PathEscape(remoteName))

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u, file)
	if err != nil {
		return err
	}

	req.ContentLength = info.Size()
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error subiendo archivo: %s", string(body))
	}

	return nil
}

type saveAsRequest struct {
	SaveFormat string `json:"SaveFormat"`
	FileName   string `json:"FileName"`
}

func (a *AsposeClient) ConvertToPDF(
	ctx context.Context,
	token string,
	remoteDocx string,
	remotePDF string,
	localPDF string,
) error {

	reqBody := saveAsRequest{
		SaveFormat: "pdf",
		FileName:   remotePDF,
	}

	data, _ := json.Marshal(reqBody)

	u := fmt.Sprintf("%s/%s/saveAs", baseAPIURL, url.PathEscape(remoteDocx))

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		u,
		bytes.NewBuffer(data),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", mime.TypeByExtension(".json"))

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error convirtiendo archivo: %s", string(body))
	}

	return a.DownloadFile(ctx, token, remotePDF, localPDF)
}

func (a *AsposeClient) DownloadFile(
	ctx context.Context,
	token string,
	remoteFile string,
	localFile string,
) error {

	u := fmt.Sprintf("%s/storage/file/%s", baseAPIURL, url.PathEscape(remoteFile))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error descargando PDF: %s", string(body))
	}

	if err := os.MkdirAll(filepath.Dir(localFile), 0755); err != nil {
		return err
	}

	out, err := os.Create(localFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

type DrawingObjectInsert struct {
	RelativeHorizontalPosition string  `json:"RelativeHorizontalPosition,omitempty"`
	RelativeVerticalPosition   string  `json:"RelativeVerticalPosition,omitempty"`
	Width                      float64 `json:"Width,omitempty"`
	Height                     float64 `json:"Height,omitempty"`
	WrapType                   string  `json:"WrapType,omitempty"`
}

func (a *AsposeClient) InsertImageAtBookmark(
	ctx context.Context,
	token string,
	remoteDocx string,
	bookmarkName string,
	localImagePath string,
) error {
	u := fmt.Sprintf("%s/%s/bookmarks/%s/drawingObjects", baseAPIURL, url.PathEscape(remoteDocx), url.PathEscape(bookmarkName))

	file, err := os.Open(localImagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add JSON part for DrawingObject
	drawingObj := DrawingObjectInsert{
		WrapType: "Inline",
		Width: 200,
		Height: 100,
	}
	drawingData, _ := json.Marshal(drawingObj)
	
	part, _ := writer.CreateFormField("DrawingObject")
	part.Write(drawingData)

	// Add file part
	filePart, err := writer.CreateFormFile("File", filepath.Base(localImagePath))
	if err != nil {
		return err
	}
	io.Copy(filePart, file)

	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error insertando imagen en marcador: %s", string(respBody))
	}

	return nil
}
