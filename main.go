package main

func main() {
	config := getConfig()
	u := Upload{}
	u.startUpload(&config)
}
