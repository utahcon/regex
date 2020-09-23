package ipv4

import "regexp"

func Validate(address string) bool {
	re := regexp.MustCompile(`^(([0-9]{1}|[0-9]{2}|1[0-9]{2}|2([0-4][0-9]|5[0-5])).){3}([0-9]{1}|[0-9]{2}|1[0-9]{2}|2([0-4][0-9]|5[0-5]))($|\/([0-9]{1}|([1-2][0-9]|3[0-2]){1,2})$)`)
	return re.Match([]byte(address))
}
