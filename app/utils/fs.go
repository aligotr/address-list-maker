package utils

import (
	"log"
	"os"
	"path/filepath"
)

/*
Создать файл и записать в него строку
  - Проверить существование папки, если нет, создать
  - Проверить существование файла, если нет, создать
  - Перезаписать содержимое файла
*/
func CreateOrTruncFileAndWriteString(dir string, fileName string, content string) {
	CreateDir(dir)

	filePath := filepath.Join(dir, fileName)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalln(err)
	}
}

/* Создать папку, включая её полный путь */
func CreateDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

/* Очистить содержимое папки */
func RemoveDirContent(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}
