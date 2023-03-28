package helpers

import "strings"

func GetDescription(arr []string) []string {
	var description string
	if len(arr) == 4 {
		description = arr[1]
	} else if len(arr) == 5 {
		description = arr[1] + arr[2]
	} else if len(arr) == 6 {
		description = arr[1] + arr[2] + arr[3]
	} else {
		description = ""
	}

	return strings.Split(description, "/")
}

// txn id, mode, destination
func TxnDescription(str []string) (string, string, string) {
	if len(str) == 4 {
		return str[3], "ATM", str[0]
	}
	if len(str) == 5 {
		return str[2], str[0], str[1]
	}
	if len(str) == 6 {
		return str[2], str[0] + " " + str[1], str[3]
	}
	return "", "", ""
}
