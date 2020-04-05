package dataloaders

import (
	"github.com/cflynn07/jide-server/database"
	"github.com/cflynn07/jide-server/types"
)

// GetNodeByID returns a database record for a given id
func GetNodeByID(id int) (types.User, error) {
	var user types.User
	database.DB.Where("idusers = ?", "2").First(&user)
	return user, nil
}
