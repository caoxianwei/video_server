package dbops

import "testing"

//init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)
func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	m.Run()
}