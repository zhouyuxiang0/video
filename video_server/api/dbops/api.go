package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, e := dbConn.Prepare("INSERT INTO users (login_name, pwd) values (?, ?)")
	if e != nil {
		return e
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, e := dbConn.Prepare("select pwd from users where login_name = ?")
	if e != nil {
		log.Printf("%s", e)
		return "", e
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, e := dbConn.Prepare("delete from users where login_name=? and pwd=?")
	if e != nil {
		log.Printf("Deleteuser error: %s", e)
		return  e
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}

