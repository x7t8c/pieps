package main

var NIDPathMap = map[NID]string{
	NID(SingleNID(0)): "/dev/nil",
	NID(SingleNID(1)): "/dev/out",
	NID(SingleNID(2)): "/dev/proc/main/cout",
	NID(SingleNID(3)): "/dev/proc/main/cinp",
}

func BytesFromString(s string) []byte {
	return []byte(s)
}

func FileNameFromBytes(b []byte) [128]byte {
	d := [128]byte{}
	for x := 0; x < len(b); x++ {
		if x > 127 {
			break
		}
		d[x] = b[x]
	}
	return d
}

var VFileNIDMap = map[NID]*VFile{
	NID(SingleNID(0)): &VFile{
		func(b []byte) (int, error) {
			return len(b), nil
		},
		func(b []byte) error {
			return nil
		},
		FileNameFromBytes(BytesFromString("/dev/nil")),
	},
}

func SingleNID(b byte) [128]byte {
	return [128]byte{b}
}

type NID [128]byte

type File struct {
	vfile  *VFile
	nid    *NID // filesystem index id
	name   [128]byte
	parent *NID
}

type Folder struct {
	children []*NID
	parent   *NID
}

type FileReader struct {
	f   *File
	pos int64
}

func (F *File) Reader() *FileReader {
	return &FileReader{
		F, 0,
	}
}

func (F *FileReader) Read() (byte, bool) {
	// eof => 0x0, true
	return 0x1, false
}

func (F *File) Size() int64 { // -1 => too big //\\ -2 => not readable
	return 0
}

func (F *File) ReadBytes() []byte {
	return []byte{0x0} // first byte <=> too big or not
}
