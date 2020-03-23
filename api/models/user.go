package models

func CreateUser(username, email, password string) error {
	insertQ, err := con.Query("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", username, email, password)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}
