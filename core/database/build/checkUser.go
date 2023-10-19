package Build

import (
	Database "Rain/core/database"
	CNC "Rain/core/config/admin"
)

func NewMongo() bool {
	_, error := Database.GetUser(CNC.NewMongo_Username)
	if error {
		return true
	} else {
		return false
	}
}
