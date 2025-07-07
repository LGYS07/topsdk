package util

type FileItem struct {

	FileName string

	Content []byte

}

func (s *FileItem) SetFileName(v string) *FileItem {
	s.FileName = v
	return s
}
func (s *FileItem) SetContent(v []byte) *FileItem {
	s.Content = v
	return s
}

func NewFileItem(fileName string, content []byte) *FileItem {
	return &FileItem{FileName: fileName,Content: content}
}
