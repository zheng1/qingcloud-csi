package block

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"os"
	"path"
)

type blockVolume struct {
	VolName string           `json:"volName"`
	VolID   string           `json:"volID"`
	VolSize int              `json:"volSize"`
	Zone    string           `json:"zone"`
	Sc      qingStorageClass `json:"storageClass"`
}

func persistVolInfo(volumeId string, persistentStoragePath string, volInfo *blockVolume) error {
	file := path.Join(persistentStoragePath, volumeId+".json")
	fp, err := os.Create(file)
	if err != nil {
		glog.Errorf("failed to create persistent storage file %s with error: %v\n", file, err)
		return fmt.Errorf("create err %s/%s", file, err)
	}
	defer fp.Close()
	encoder := json.NewEncoder(fp)
	if err = encoder.Encode(volInfo); err != nil {
		glog.Errorf("failed to encode volInfo: %+v for file: %s with error: %v\n", volInfo, file, err)
		return fmt.Errorf("encode err: %v", err)
	}
	glog.Infof("successfully saved volInfo: %+v into file: %s\n", volInfo, file)
	return nil
}

func loadVolInfo(volumeId string, persistentStoragePath string, volInfo *blockVolume) error {
	file := path.Join(persistentStoragePath, volumeId+".json")
	fp, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("open err %s/%s", file, err)
	}
	defer fp.Close()

	decoder := json.NewDecoder(fp)
	if err = decoder.Decode(volInfo); err != nil {
		return fmt.Errorf("decode err: %v.", err)
	}

	return nil
}

func deleteVolInfo(volumeId string, persistentStoragePath string) error {
	file := path.Join(persistentStoragePath, volumeId+".json")
	glog.Infof("Deleting file for Volume: %s at: %s resulting path: %+v\n", volumeId, persistentStoragePath, file)
	err := os.Remove(file)
	if err != nil {
		if err != os.ErrNotExist {
			return fmt.Errorf("error removing file: %s/%s", file, err)
		}
	}
	return nil
}