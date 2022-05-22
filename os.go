package main

type os1212121212 struct{}

var os = os1212121212{}

func (so os1212121212) OpenFile(path string) *File {
	return nil // do this later :D
}

func main() {
	InitMemp()
	for {
		println(len(memp))
	}
}
