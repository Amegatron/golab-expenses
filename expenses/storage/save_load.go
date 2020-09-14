package storage

import (
	"encoding/json"
	"io"
	"os"
)

// Usage:
//var pathToFile = "file.json"
//
//func ExampleSaveToFile() {
//	value := struct {
//		Name string `json:"name"`
//		Value string `json:"value"`
//	}{
//		Name: "blahBlahBlah",
//		Value: "LaLaLaLa",
//	}
//
//	file, err := NewFileSaver(pathToFile)
//	if err != nil {
//		log.Printf("error creating file: %s\n", err.Error())
//		return
//	}
//	defer file.Close()
//	if err := file.Save(&value); err != nil {
//		log.Printf("error writing file: %s\n", err.Error())
//		return
//	}
//	log.Printf("Value saved: %+v\n", value)
//}
//func ExampleLoadFromFile() {
//	var value struct {
//		Name string `json:"name"`
//		Value string `json:"value"`
//	}
//
//	file, err := NewFileLoader(pathToFile)
//	if err != nil {
//		log.Printf("error opening file: %s\n", err.Error())
//		return
//	}
//	defer file.Close()
//	if err := file.Load(&value); err != nil {
//		log.Printf("error reading file: %s\n", err.Error())
//		return
//	}
//	log.Printf("value read: %+v\n", value)
//}

type Saver interface {
	Save(src interface{}) error
}
type Loader interface {
	Load(dst interface{}) error
}

type IOSaver struct {
	dst io.Writer
}

func NewIOSaver(writer io.Writer) *IOSaver {
	return &IOSaver{dst: writer}
}
func (saver *IOSaver) Save(src interface{}) error {
	jsonEncoder := json.NewEncoder(saver.dst)
	return jsonEncoder.Encode(src)
}

type IOLoader struct {
	src io.Reader
}

func NewIOLoader(reader io.Reader) *IOLoader {
	return &IOLoader{src: reader}
}
func (loader *IOLoader) Load(dst interface{}) error {
	jsonDecoder := json.NewDecoder(loader.src)
	return jsonDecoder.Decode(dst)
}

type FileSaver struct {
	file *os.File
}

func NewFileSaver(path string) (*FileSaver, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &FileSaver{file: file}, nil
}
func (fs *FileSaver) Save(src interface{}) error {
	return NewIOSaver(fs.file).Save(src)
}
func (fs *FileSaver) Close() error {
	return fs.file.Close()
}

type FileLoader struct {
	file *os.File
}

func NewFileLoader(path string) (*FileLoader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &FileLoader{file: file}, nil
}
func (fl *FileLoader) Load(dst interface{}) error {
	return NewIOLoader(fl.file).Load(dst)
}
func (fl *FileLoader) Close() error {
	return fl.file.Close()
}
