package tool

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// GetFirstLevelFolderNames 取得目錄第一層資料夾路名
func GetFirstLevelFolderNames(pwd string) []string {
	var names []string

	_ = filepath.Walk(pwd, func(path string, info fs.FileInfo, err error) error {
		if path == pwd {
			return nil
		}

		if !info.IsDir() {
			return nil
		}

		path = strings.ReplaceAll(path, pwd, "")
		paths := strings.Split(path, "/")
		if len(paths) > 2 || len(paths) <= 1 {
			return nil
		}

		if paths[1] == "asset" || paths[1] == "fyne-cross" || paths[1] == "view" || paths[1] == "tool" {
			return nil
		}

		names = append(names, paths[1])

		return nil
	})

	return names
}
