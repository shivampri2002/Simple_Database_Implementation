package database

import (
	// "encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}

	mutex := d.GetOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	fnlPath := filepath.Join(dir, resource+".bson")
	tmpPath := fnlPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// b, err := json.MarshalIndent(v, "", "\t")
	b, err := bson.Marshal(v)
	if err != nil {
		return err
	}

	// b = append(b, byte('\n'))
	if err = os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}
	return os.Rename(tmpPath, fnlPath)
}

func (d *Driver) Read(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - unable to read!")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource - unable to read record (no name)!")
	}

	record := filepath.Join(d.dir, collection, resource)

	if _, err := Stat(record); err != nil {
		return err
	}

	b, err := os.ReadFile(record + ".bson")
	if err != nil {
		return err
	}
	// return json.Unmarshal(b, &v)
	return bson.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([][]byte, error) {
	if collection == "" {
		return nil, fmt.Errorf("Missing collection - unable to read!")
	}
	dir := filepath.Join(d.dir, collection)
	if _, err := Stat(dir); err != nil {
		return nil, err
	}
	files, _ := os.ReadDir(dir)
	var records [][]byte
	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		records = append(records, b)
	}
	return records, nil
}

func (d *Driver) Delete(collection, resource string) error {
	path := filepath.Join(collection, resource)
	mutex := d.GetOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)
	switch fi, err := Stat(dir); {
	case fi == nil || err != nil:
		return fmt.Errorf("Unable to find file or directory named %v\n", path)
	case fi.Mode().IsDir():
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".bson")
	}
	return nil
}
