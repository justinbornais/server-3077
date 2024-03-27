package handlers

import (
	"strings"

	"github.com/justinbornais/server-3077/utilities"
)

func GetUserTypes() ([]utilities.UserType, error) {
	db := utilities.GetDB()

	rows, err := db.Query("SELECT id, name FROM user_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userTypes []utilities.UserType
	for rows.Next() {
		var ut utilities.UserType
		if err := rows.Scan(&ut.ID, &ut.Name); err != nil {
			return nil, err
		}
		ut.Name = strings.ToUpper(ut.Name)
		userTypes = append(userTypes, ut)
	}

	return userTypes, nil
}
