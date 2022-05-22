package main

type VFile struct { /*Virtual file for emulating linux' /dev/nil or /dev/proc/1234/ file system*/
	Write func([]byte) (int, error)
	Read  func([]byte) error
}
