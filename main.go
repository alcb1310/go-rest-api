package main

import "gitlab.com/0x4149/logz"

const local bool = true

func init() {
	if local {
		logz.VerbosMode()
	}

	logz.Run()
}

func main() {}
