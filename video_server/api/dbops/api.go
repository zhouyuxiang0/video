package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/defs"
	"video_server/utils"
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
	vid, e := utils.NewUUID()
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

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}

