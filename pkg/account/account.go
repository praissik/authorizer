package account

import "authorizer/pkg/database"

func IsEmailExists(email string) (bool, error) {
	db, err := database.GetMySQL()
	if err != nil {
		return false, err
	}

	row := db.QueryRow(`SELECT true FROM accounts WHERE Email = ?`, email)

	found := false
	err = row.Scan(&found)
	if err != nil {
		return false, err
	}

	return found, nil
}
