package model

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/logoove/sqlite" // 导入sqlite
	log "github.com/sirupsen/logrus"

	"github.com/huoxue1/filecloud/config"
)

type FileInfo struct {
	ID         string   `json:"id"`
	FileName   string   `json:"file_name"`
	IsDir      int      `json:"is_dir"`
	Content    []string `json:"content"`
	UploadTime int64    `json:"upload_time"`
	ParentID   string   `json:"parent_id"`
	MD5        string   `json:"md5"`
	Size       int64    `json:"size"`
}

var db *sql.DB

func Init() {
	DB, err := sql.Open("sqlite3", config.GetConfig().DBPath)
	if err != nil {
		log.Panicln("the database init error:", err.Error())
		return
	}
	db = DB
	db.Exec(`create table if not exists file(
    id        TEXT not null
        constraint file_pk
            primary key,
    file_name TEXT not null,
    is_dir    integer default 0,
    content   blob,
    upload_time INTEGER,
    parent_id TEXT default '',
    md5 TEXT,
    size INTEGER default 0
);`)
}

func InitRoot() {
	parent, err := QueryByParent("")
	if err != nil {
		log.Errorln("查询root出现错误")
		log.Errorln(err.Error())
		return
	}
	if len(parent) < 1 {
		err := InsertFile(FileInfo{
			ID:         "",
			FileName:   config.GetConfig().StoragePath,
			IsDir:      1,
			Content:    []string{},
			UploadTime: time.Now().Unix(),
			ParentID:   "",
			MD5:        "",
		})
		if err != nil {
			log.Errorln("插入root节点出现错误")
			log.Errorln(err.Error())
			return
		}
	}
}

func InsertFile(info FileInfo) error {
	info.ID = uuid.New().String()

	data, _ := json.Marshal(info.Content)
	_, err := db.Exec(`insert into file values (?,?,?,?,?,?,?,?)`, info.ID, info.FileName, info.IsDir, data, info.UploadTime, info.ParentID, info.MD5, info.Size)
	if err != nil {
		return err
	}
	return nil
}

func QueryByID(id string, f *FileInfo) error {
	row := db.QueryRow("select * from file where id=?", id)
	var data []byte
	err := row.Scan(&f.ID, &f.FileName, &f.IsDir, &data, &f.UploadTime, &f.ParentID, &f.MD5, &f.Size)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &f.Content)
	if err != nil {
		return err
	}
	return err
}

func QueryByMd5(md5 string, f *FileInfo) error {
	row := db.QueryRow("select * from file where md5=?", md5)
	var data []byte
	err := row.Scan(&f.ID, &f.FileName, &f.IsDir, &data, &f.UploadTime, &f.ParentID, &f.MD5, &f.Size)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &f.Content)
	if err != nil {
		return err
	}
	return err
}

func QueryByParent(parentID string) ([]*FileInfo, error) {
	var infos []*FileInfo
	var fileInfos []*FileInfo
	rows, err := db.Query("select * from file where parent_id=?", parentID)
	if err != nil {
		return infos, err
	}
	defer rows.Close()
	for rows.Next() {
		f := new(FileInfo)
		var data []byte
		err := rows.Scan(&f.ID, &f.FileName, &f.IsDir, &data, &f.UploadTime, &f.ParentID, &f.MD5, &f.Size)
		if err != nil {
			return infos, err
		}
		err = json.Unmarshal(data, &f.Content)
		if err != nil {
			return infos, err
		}
		if f.IsDir == 1 {
			infos = append(infos, f)
		} else {
			fileInfos = append(fileInfos, f)
		}
	}
	infos = append(infos, fileInfos...)
	return infos, err
}

func QueryByName(name string) ([]*FileInfo, error) {
	var infos []*FileInfo
	rows, err := db.Query(`select * from file where file_name like '%` + name + `%';`)
	if err != nil {
		return infos, err
	}
	defer rows.Close()
	for rows.Next() {
		f := new(FileInfo)
		var data []byte
		err := rows.Scan(&f.ID, &f.FileName, &f.IsDir, &data, &f.UploadTime, &f.ParentID, &f.MD5, &f.Size)
		if err != nil {
			return infos, err
		}
		err = json.Unmarshal(data, &f.Content)
		if err != nil {
			return infos, err
		}
		infos = append(infos, f)
	}
	return infos, err
}

func CheckDirRepeat(parentID string, name string) bool {
	row := db.QueryRow("select count(*) from file where file_name=? and parent_id=? and is_dir=1;", name, parentID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return true
	}
	return count > 0
}

func DeleteByID(id string) error {
	f := new(FileInfo)
	err := QueryByID(id, f)
	if err != nil {
		return err
	}

	// 删除对象为文件夹，判断是否为空文件夹，如果是就删除，否则抛出异常
	if f.IsDir == 1 {
		infos, err := QueryByParent(f.ID)
		if err != nil {
			return err
		}
		if len(infos) > 0 {
			return errors.New("the dir has file,can't delete the dir")
		}
	}

	for _, file := range f.Content {
		err := os.Remove(fmt.Sprintf("%v/%v", config.GetConfig().StoragePath, file))
		if err != nil {
			return err
		}
	}

	_, err = db.Exec("delete from file where id=?;", id)
	if err != nil {
		return err
	}
	return err
}
