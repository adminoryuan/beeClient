package main

import "sync"

func main() {

	syncLock := sync.WaitGroup{}
	syncLock.Add(1)

	config := getConfig()
	u := Upload{}
	u.startUpload(&config)

	syncLock.Wait()
}
