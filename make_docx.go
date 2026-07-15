package main
import (
	"archive/zip"
	"os"
)
func main() {
	f, _ := os.Create("valid.docx")
	zw := zip.NewWriter(f)
	w, _ := zw.Create("word/document.xml")
	w.Write([]byte(`<?xml version="1.0"?><w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"><w:body><w:p><w:r><w:t>Hello {nombre}</w:t></w:r></w:p></w:body></w:document>`))
	zw.Close()
	f.Close()
}
