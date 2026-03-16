package wordnet

import (
	"io/fs"
	"path/filepath"
)

func openFile(dir fs.FS, basedir string, filename string) (fs.File, error) {
	return dir.Open(filepath.Join(basedir, filename))
}
