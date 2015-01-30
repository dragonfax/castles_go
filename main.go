package main

func main() {
	initWindow()
	go MuxEvents()
	go NewGame().run()
	MainQueueLoop()
}
