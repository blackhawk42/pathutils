package pathutils

import(
	"fmt"
	"path/filepath"
	"os"
	"strings"
)

// Golang: "A simple 'isFile' or 'FileExists' function? Fuck that, you filthy causual!"
func FileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	
	if err != nil {
		if os.IsNotExist(err) { // File doesn't exists
			return false, nil
		} else {
			return false, err // an error ocurred
		}
	}
	
	return true, nil
}

// Split filename and extension, where path = filename + extension.
// Yes, It's basically Python's os.path.splitext...
func Splitext(path string) (string, string) {
	return strings.TrimSuffix(path, filepath.Ext(path)), filepath.Ext(path)
}

// Dealing with already existing files with the same name. Returns a string
// with the Windows-popularized format of "filename (number).extension".
func RepeatedFilenames(path string) (string, error) {
	base, ext := Splitext(path)
	files_this_name := 0
	filename := fmt.Sprintf("%s%s", base, ext)
	
	file_exists, err := FileExists(filename)
	if err != nil {
		return filename, err
	}
	
	for file_exists{
		files_this_name++
		filename = fmt.Sprintf("%s (%d)%s", base, files_this_name, ext)
		
		file_exists, err = FileExists(filename)
		if err != nil {
			return "error", err
		}
	}
	return filename, nil
}
