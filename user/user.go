
package user

import "strings"

func CheckUserName(username string) bool {
	if len(username) < 6 || strings.Contains(username, "admin") {
		return false
	}
	return true
}
