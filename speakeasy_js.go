// based on https://code.google.com/p/gopass
// Author: johnsiilver@gmail.com (John Doak)
//
// Original code is based on code by RogerV in the golang-nuts thread:
// https://groups.google.com/group/golang-nuts/browse_thread/thread/40cc41e9d9fc9247

//go:build js
// +build js

package speakeasy

import (
	"os"

	"golang.org/x/term"
)

// getPassword gets input hidden from the terminal from a user. This is
// accomplished by turning off terminal echo, reading input from the user and
// finally turning on terminal echo.
func getPassword() (password string, err error) {
	passwd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	return string(passwd), nil
}
