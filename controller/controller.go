package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/filecloud/config"
	"github.com/huoxue1/filecloud/model"
	"github.com/huoxue1/filecloud/service"
	"github.com/huoxue1/filecloud/utils"
)

// GetFilesFromPath
/*
 * Method: GET
 * path: /files
 * params: parent_id string query
 * example: /files?parent_id=
 * return []*model.FileInfo
 */
func GetFilesFromPath() gin.HandlerFunc {
	return func(context *gin.Context) {
		parentID, b := context.GetQuery("parent_id")
		if !b {
			context.JSON(200, fail("the params not found", errors.New("error")).Error)
		}
		infos, err := model.QueryByParent(parentID)
		if err != nil {
			context.JSON(200, fail("select the content fail", err.Error()))
			return
		}
		context.JSON(200, ok("success", infos))
	}
}

// GetFile
/*
 * Method: GET
 * path: /file
 * params: id query
 * example: /file
 * return model.FileInfo
 */
func GetFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("id")
		if !b {
			context.JSON(200, fail("the params not found", nil))
			return
		}
		f := new(model.FileInfo)
		err := model.QueryByID(id, f)
		if err != nil {
			context.JSON(200, fail("select the file error", err.Error()))
			return
		}
		context.JSON(200, ok("success", f))
	}
}

// GetRoot
/*
 * Method: GET
 * path: /root
 * params: null
 * example: /root
 * return model.FileInfo
 */
func GetRoot() gin.HandlerFunc {
	return func(context *gin.Context) {
		parent, err := model.QueryByParent("")
		if err != nil {
			context.JSON(200, fail("select the root error", err.Error()))
			return
		}
		context.JSON(200, ok("success", parent[0]))
	}
}

// CheckFile
/*
 * Method: POST
 * path: /check_file
 * params: {name:"",md5:""} body
 * example: /check_file {"name":"123","md5":"123"}
 * return ok bool
 */
func CheckFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		type params struct {
			Name     string `json:"name"`
			MD5      string `json:"md5"`
			ParentID string `json:"parent_id"`
		}
		param := new(params)
		err := context.BindJSON(param)
		if err != nil {
			context.JSON(200, fail("bind the params fail", err.Error()))
			return
		}
		f := new(model.FileInfo)
		err = model.QueryByMd5(param.MD5, f)
		if err != nil {
			context.JSON(200, fail("bind the params fail", err.Error()))
			return
		}
		if f.ID == "" {
			context.JSON(200, ok("the file not found", map[string]bool{"ok": false}))
			return
		}
		f.FileName = param.Name
		f.ParentID = param.ParentID
		f.UploadTime = time.Now().Unix()
		err = model.InsertFile(*f)
		if err != nil {
			context.JSON(200, ok("insert the file err:"+err.Error(), map[string]bool{"ok": false}))
			return
		}
		context.JSON(200, ok("success", map[string]bool{"ok": true}))
	}
}

// Upload
/*
 * Method: POST
 * path: /upload
 * params: file
 * example: /upload
 * return null
 */
func Upload() gin.HandlerFunc {
	return func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(200, fail("fail", err.Error()))
			return
		}
		parentID := context.Query("parent_id")
		f, err := file.Open()
		if err != nil {
			context.JSON(200, fail("open the file error", err.Error()))
			return
		}
		files, err := service.WriterToFiles(f)
		if err != nil {
			log.Errorln(err.Error())
			context.JSON(200, fail("save the file error", err.Error()))
			return
		}
		md5 := utils.GetFileMd5(f)

		info := model.FileInfo{
			ID:         "",
			FileName:   file.Filename,
			IsDir:      0,
			Content:    files,
			UploadTime: time.Now().Unix(),
			ParentID:   parentID,
			MD5:        md5,
			Size:       file.Size,
		}
		err = model.InsertFile(info)
		if err != nil {
			context.JSON(200, fail("save the file error", err.Error()))
			return
		}
		context.JSON(200, ok("success", nil))
	}
}

// Mkdir
/*
 * Method: POST
 * path: /mkdir
 * params: {parent_id:"",name:""}
 * example: /mkdir
 * return null
 */
func Mkdir() gin.HandlerFunc {
	type params struct {
		ParentID string `json:"parent_id"`
		Name     string `json:"name"`
	}
	return func(context *gin.Context) {
		p := new(params)
		err := context.BindJSON(p)
		if err != nil {
			context.JSON(200, fail("the params not found", err.Error()))
			return
		}
		repeat := model.CheckDirRepeat(p.ParentID, p.Name)
		if repeat {
			context.JSON(200, fail("the dir id exist", "repeat"))
			return
		}
		info := model.FileInfo{
			ID:         "",
			FileName:   p.Name,
			IsDir:      1,
			Content:    []string{},
			UploadTime: time.Now().Unix(),
			ParentID:   p.ParentID,
			MD5:        "",
		}
		err = model.InsertFile(info)
		if err != nil {
			context.JSON(200, fail("inset the dir to db error", err.Error()))
			return
		}
		context.JSON(200, ok("success", nil))
	}
}

// DeleteFile
/*
 * Method: POST
 * path: /delete_file
 * params: id query
 * example: /delete_file
 * return null
 */
func DeleteFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("id")
		if !b {
			context.JSON(200, fail("the params not found", ""))
			return
		}
		err := model.DeleteByID(id)
		if err != nil {
			context.JSON(200, fail("delete the file fail", err.Error()))
			return
		}
		context.JSON(200, ok("success", nil))
	}
}

// DownloadFile
/*
 * Method: GET
 * path: /download
 * params: id query
 * example: /download
 * return null
 */
func DownloadFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, b := context.GetQuery("id")
		if !b {
			context.JSON(403, fail("the id not found", nil))
			return
		}
		info := new(model.FileInfo)
		err := model.QueryByID(id, info)
		if err != nil {
			context.JSON(404, fail("select the file error", err.Error()))
			return
		}
		if info.IsDir == 1 {
			context.JSON(401, fail("the file id a dir", nil))
			return
		}

		context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", info.FileName))
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.Writer.Header().Set("Transfer-Encoding", "chunked")
		context.Writer.WriteHeader(http.StatusOK)

		for _, s := range info.Content {
			file, err := os.OpenFile(config.GetConfig().StoragePath+"/"+s, os.O_RDWR, 0666)
			if err != nil {
				context.JSON(404, fail("open"+s+" the file error", err.Error()))
				return
			}
			_, err = io.Copy(context.Writer, file)
			if err != nil {
				context.JSON(404, fail("open"+s+" the file error", err.Error()))
				return
			}
			context.Writer.(http.Flusher).Flush()
		}
	}
}
