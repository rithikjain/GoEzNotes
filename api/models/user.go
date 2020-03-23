package models

func CreateUser(username, email, password string) (error, int64) {
	insertQ, err := con.Exec("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", username, email, password)
	if err != nil {
		return err, -1
	}
	id, _ := insertQ.LastInsertId()
	return nil, id
}
