package engines

import (
	"github.com/Tsuki124/crawlab-sdk/internal/constants"
	"github.com/Tsuki124/crawlab-sdk/internal/interfaces"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/crawlab-team/go-trace"
	"github.com/tidwall/gjson"
	"github.com/go-resty/resty/v2"
	"os"
	"path/filepath"
	"strings"
)

type SeaweedFS struct {
	interfaces.SeaweedFS

	_ADDR   string
	_CLIENT *resty.Client
}

func NewSeaweedFS(prePath string) interfaces.SeaweedFS {
	client := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	fsURL := strings.TrimRight(os.Getenv(constants.ENV_SEAWEED_FS_FILER_URL), "/")
	prePath = strings.Trim(prePath, "/")
	addr := fmt.Sprintf("%s/%s", fsURL, prePath)

	engine := &SeaweedFS{
		_ADDR:   addr,
		_CLIENT: client,
	}
	return engine
}

func (my *SeaweedFS) _URL(path string) string {
	path = strings.Trim(path, "/")
	return my._ADDR + "/" + path
}

func (my *SeaweedFS) _ListJsonURL(path string) string {
	return fmt.Sprintf(constants.KEY_SEAWEED_FS_FILE_LIST_JSON, my._URL(path))
}

func (my *SeaweedFS) _InfoJsonURL(path string) string {
	return fmt.Sprintf(constants.KEY_SEAWEED_FS_FILE_INFO_JSON, my._URL(path))
}

func (my *SeaweedFS) List(dirpath ...string) ([]interfaces.SeaweedFile, error) {
	var path string
	if len(dirpath)>0 {
		path = dirpath[0]
	}

	jsonURL := my._ListJsonURL(path)
	req := my._CLIENT.SetHeader("Accept", "application/json").R()

	resp, err := req.Get(jsonURL)
	if err != nil {
		return nil, trace.Error(err)
	}
	if resp.IsError() {
		return nil, trace.Error(errors.New(fmt.Sprintf("the statuscode is %d", resp.StatusCode())))
	}

	files := make([]interfaces.SeaweedFile, 0)
	bodyJson := gjson.ParseBytes(resp.Body())
	prePath := bodyJson.Get("Path").Str
	for _, entityJson := range bodyJson.Get("Entries").Array() {
		file := &SeaweedFile{}
		fullPath := entityJson.Get("FullPath").Str
		name := strings.TrimLeft(fullPath,prePath)
		mime := entityJson.Get("Mime").Str

		file.name = name
		file.path = fullPath
		if mime=="" {
			file.isDir = true
		}

		files = append(files,file)
	}

	return files, nil
}

func (my *SeaweedFS) Download(path string) ([]byte, error) {
	url := my._URL(path)
	req := my._CLIENT.R()

	resp,err := req.Get(url)
	if err!=nil {
		return nil,trace.Error(err)
	}
	if resp.IsError() {
		return nil,trace.Error(errors.New(fmt.Sprintf("the statuscode is %d",resp.StatusCode())))
	}

	return resp.Body(), nil
}

func (my *SeaweedFS) Upload(path string, content []byte) error {
	url := my._URL(path)
	req := my._CLIENT.R()
	req.SetHeader("Accept-Encoding","gzip, deflate")

	req.SetBody(content)
	resp,err := req.Put(url)
	if err!=nil {
		return trace.Error(err)
	}
	if resp.IsError() {
		return trace.Error(errors.New(fmt.Sprintf("the statuscode is %d",resp.StatusCode())))
	}

	return nil
}

func (my *SeaweedFS) Delete(path string) error {
	url := my._URL(path)
	req := my._CLIENT.R()

	resp,err := req.Delete(url)
	if err!=nil {
		return trace.Error(err)
	}
	if resp.IsError() {
		return trace.Error(errors.New(fmt.Sprintf("the statuscode is %d",resp.StatusCode())))
	}

	return nil
}

func (my *SeaweedFS) Info(path string) (interfaces.SeaweedFile,error) {
	url := my._InfoJsonURL(path)
	req := my._CLIENT.R()

	resp,err := req.Get(url)
	if err!=nil {
		return nil,trace.Error(err)
	}
	if resp.IsError() {
		return nil,trace.Error(errors.New(fmt.Sprintf("the statuscode is %d",resp.StatusCode())))
	}

	infoJson := gjson.ParseBytes(resp.Body())
	fullPath := infoJson.Get("FullPath").Str
	name := filepath.Base(fullPath)
	mime := infoJson.Get("Mime").Str

	file := &SeaweedFile{name: name,path: fullPath}
	if mime=="" {
		file.isDir = true
	}

	return file,nil
}
