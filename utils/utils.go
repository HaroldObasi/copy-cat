package utils

import (
	"os"
)

// GetFilesInDirectory returns a list of files in a directory
func GetFilesInDirectory(rootPath, dirPath string) ([]string, error) {
    var files []string

    entries, err := os.ReadDir(rootPath + dirPath)
    if err != nil {
        return nil, err
    }

    for _, entry := range entries {
        if entry.Name() == "node_modules" && entry.IsDir() {
            continue
        }

        if entry.IsDir() {
            nestedFiles, err := GetFilesInDirectory(dirPath + "/", entry.Name())
            if err != nil {
                return nil, err
            }
            files = append(files, nestedFiles...)
            continue
        }
        
        var name string
        if rootPath == "./" {
            name = entry.Name()
        } else {
            name = dirPath + "/" + entry.Name()
        }

        files = append(files, name)
    }

    return files, nil
}