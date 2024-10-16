package cmd

var data = make(map[string]string)

func SetCMD(key, value string) (string) {
	data[key] = value
	return "OK"
}

func GetCMD(key string) (string, string) {
	if len(data[key]) > 0 {
		return data[key], ""
	}else {
		return "", "$-1"
	} 
}