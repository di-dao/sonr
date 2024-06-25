package fs

type Folder string

func (f Folder) Path() string {
	return string(f)
}
