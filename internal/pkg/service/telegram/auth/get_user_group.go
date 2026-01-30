package auth

import (
	"os"
	"strconv"
)

type UserCategiry string

const (
	SuperAdmin     UserCategiry = "SuperAdmin"
	LibrarianAdmin UserCategiry = "LibrarianAdmin"
	SimpleUser     UserCategiry = "SimpleUser"
)

// GetUserCategory ... TODO подумать как сделать отдельный подсервис
func GetUserCategory(id int64) UserCategiry {
	userID := strconv.FormatInt(id, 10)

	superAdminID := os.Getenv("SUPER_ADMIN_ID")

	switch {
	case userID == superAdminID:
		return SuperAdmin
	default:
		return SimpleUser
	}
}
