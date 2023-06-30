package utility

import (
	"regexp"
	"strconv"
	"strings"
)

var envfilepath = ".env"

func IsUsernameInvalid(username string) bool {
	minUsername, _ := strconv.Atoi(DotEnv("MIN_UNAME", envfilepath))
	maxUsername, _ := strconv.Atoi(DotEnv("MAX_UNAME", envfilepath))

	if len(username) < minUsername || len(username) > maxUsername {
		return true
	}

	return false
}

func IsPhoneInvalid(phoneNumber string) bool {
	minPhoneNum, _ := strconv.Atoi(DotEnv("MIN_PHONE_NUM", envfilepath))
	maxPhoneNum, _ := strconv.Atoi(DotEnv("MAX_PHONE_NUM", envfilepath))
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")

	if len(phoneNumber) < minPhoneNum || len(phoneNumber) > maxPhoneNum {
		return true
	}

	var validPhone = regexp.MustCompile(`^[0-9]+$`).MatchString(phoneNumber)

	return !validPhone
}

func IsEmailInvalid(email string) bool {
	for _, c := range email {
		if c < 31 || c > 127 || strings.ContainsAny(string(c), `()<>,;:\\"[]`) {
			return true
		}
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return !emailRegex.MatchString(email)
}

func IsPasswordInvalid(password string) bool {
	minPassword, _ := strconv.Atoi(DotEnv("MIN_PASS", envfilepath))
	maxPassword, _ := strconv.Atoi(DotEnv("MAX_PASS", envfilepath))

	if len(password) <= minPassword || len(password) >= maxPassword {
		return true
	} else if strings.ContainsAny(password, ` ^*+=-_()<>,;:\\"[]`) {
		return true
	}

	return false
}
