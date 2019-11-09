package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/defs"
	"video_server/untils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, e := dbConn.Prepare("INSERT INTO users (login_name, pwd) values (?, ?)")
	if e != nil {
		return e
	}
	_, e = stmtIns.Exec(loginName, pwd)
	if e != nil {
		return e
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, e := dbConn.Prepare("select pwd from users where login_name = ?")
	if e != nil {
		log.Printf("%s", e)
		return "", e
	}
	var pwd string
	e = stmtOut.QueryRow(loginName).Scan(&pwd)
	if e != nil && err != sql.ErrNoRows {
		return "", e
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, e := dbConn.Prepare("delete from users where login_name=? and pwd=?")
	if e != nil {
		log.Printf("Deleteuser error: %s", e)
		return  e
	}
	_, e = stmtDel.Exec(loginName, pwd)
	if e != nil {
		return nil
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error)  {
	// create uid
	vid, e := untils.NewUUID()
	if e != nil {
		return nil, e
	}

	now := time.Now()
	ctime := now.Format("Jan 02 2006, 15:04:05")
	stmtIns, e := dbConn.Prepare("insert into video_info (id, author_id, name, display_ctime) value (?, ?, ?, ?)")
	if e != nil {
		return nil, e
	}
	_, e = stmtIns.Exec(vid, aid, name, ctime)
	if e != nil {
		return nil, e
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil
}

