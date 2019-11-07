package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUserCredential(t *testing.T) {
	e := AddUserCredential("avenssi", "123")
	if e != nil {
		t.Errorf("Error of AddUser: %v", e)
	}
}

func testGetUserCredential(t *testing.T) {
	pwd, e := GetUserCredential("avenssi")
	if pwd != "123" || e != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	e := DeleteUser("avenssi", "123")
	if e != nil {
		t.Errorf("Error of DeleteUser: %v", e)
	}
}

func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleteing user test failed")
	}
}
