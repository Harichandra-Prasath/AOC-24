package calendar

import "os"

func Extract(fp string) string {

	data, _ := os.ReadFile(fp)
	return string(data)

}
