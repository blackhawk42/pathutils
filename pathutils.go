package pathutils

import(
	"fmt"
	"path/filepath"
	"os"
	"strings"
)

// Golang: "A simple 'isFile' or 'FileExists' function? Fuck that, you filthy causual!"
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err){
		return true
	} else {
		return false
	}
}

// Split filename and extension, where path = filename + extension.
// Yes, It's basically Python's os.path.splitext...
func Splitext(path string) (string, string) {
	return strings.TrimSuffix(path, filepath.Ext(path)), filepath.Ext(path)
}

// Dealing with already existing files with the same name. Returns a string
// with the Windows-popularized format of "filename (number).extension".
func RepeatedFilenames(url string) string {
	base, ext := Splitext(url)
	files_this_name := 0
	filename := fmt.Sprintf("%s%s", base, ext)
	for FileExists(filename) {
		files_this_name++
		filename = fmt.Sprintf("%s (%d)%s", base, files_this_name, ext)
	}
	return filename
}
