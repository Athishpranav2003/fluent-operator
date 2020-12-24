package utils

import "strings"

func ContainString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func ConcatString(slice []string, sep string) string {

	if slice == nil || len(slice) == 0 {
		return ""
	}

	ns := ""
	for _, s := range slice {
		ns = ns + sep + s
	}

	return strings.TrimSuffix(ns, sep)
}
