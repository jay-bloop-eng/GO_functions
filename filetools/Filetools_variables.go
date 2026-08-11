package filetools

type File_cutted struct {
	File_p1     []byte
	File_p2     []byte
	File_p3     []byte
	File_p4     []byte
	File_remain []byte
}
type File_whole struct {
	File []byte
}

const (
	READ_BUFF_SIZE int = 2048
)

type File_cuti interface {
	File_cut() *File_cutted
}
