package main

type VFile struct { /*Virtual file for emulating linux' /dev/nil or /dev/proc/1234/ file system*/
	Write func([]byte) (int, error)
	Read  func([]byte) error
	name  [128]byte
}

func GetFileByNID(targetNID NID) *File {
	for nid, vfile := range VFileNIDMap {
		if nid == targetNID {
			return &File{vfile, &nid, vfile.name, nil}
		}
	}
	// get file not vfile
	return nil
}
