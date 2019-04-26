// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\x6f\x4f\x1b\x3f\x12\x7e\x5d\x7f\x8a\xe9\x86\x5f\x29\x1c\xce\x02\xed\xf5\xaa\xb4\xe9\x89\x52\x2a\x21\x71\xa5\x82\x9e\xfa\xa2\xad\x90\xb3\x9e\xdd\x75\xe3\xb5\x7d\xfe\x93\x40\xe9\x7e\xf7\x93\x77\x37\x21\xff\xe0\x87\x54\x5e\x84\xe0\x19\xcf\x8c\x9f\xe7\xf1\x78\xe8\x3d\x4d\x47\x42\xa5\x23\xe6\x4a\xa0\x78\x4d\x88\xc8\xe1\x29\x14\x16\x0d\xa4\x13\x66\x53\x29\x46\x29\xd7\xd9\x18\x2d\xa4\xe8\xb3\x34\x77\x9e\x8d\xde\x80\x2f\x51\x11\x00\x77\xe3\x3c\x56\x99\x97\xe0\xbc\x36\xd0\x3a\xf6\x1d\xda\x89\xc8\x90\x00\x54\xe3\xdc\xf5\xaf\x73\x07\x34\x87\x94\xe3\x24\xe5\xc2\x8d\x53\xf6\x2b\x58\x4c\x2d\x3a\x1d\x6c\x86\xd4\x30\xeb\x0f\x08\x00\x66\xa5\x86\xed\x87\xdd\x60\xad\x2a\x88\xe1\xa1\xb0\xe6\x7f\x41\x7b\x06\xb0\x0f\xfb\xdb\xf0\xee\xdd\x5d\xb1\xb1\x0c\x1d\x94\x5f\xdd\x49\x00\x2c\x3a\xaf\x2d\x66\x5a\x01\xbd\xd8\x60\xcf\x98\x87\x36\x52\xbb\x94\x72\x86\x95\x56\xfd\x9f\x4e\x2b\x78\xfb\x76\xfb\xe4\xfc\xe3\x36\xb9\x25\x00\x89\xd4\x05\xe5\x56\x4c\xd0\x26\x03\x48\x7e\xea\x60\x15\x93\x3c\x21\x35\x39\x39\xff\xb8\x02\x14\xb3\x7e\x15\xa9\x5c\x10\xd2\x9d\xc7\x04\x29\xe1\xf6\x16\xfa\xc7\x5a\xe5\xa2\xe8\x9f\x56\xac\x40\xd7\xff\xa4\x39\x42\x5d\xc3\xb3\x77\x0d\x40\x2a\x7a\x3d\xdb\xc8\x16\xfa\x8c\x6f\xe2\x6a\xce\xc5\x2a\xc2\x2e\x73\xe2\x20\x95\x41\xed\xc3\xef\xdf\xe0\x6d\xc0\x7b\xc9\x58\x70\x5d\x49\xd8\xd2\xc0\x31\x67\x41\x7a\xf7\x28\x1a\xe2\xbe\xfb\x49\x68\xac\x11\x97\x5c\x5b\xe0\xce\x83\x50\xe0\x33\xb3\xf7\xfa\xe5\xcb\x97\x6f\x80\x6b\xf2\xc4\x58\xed\xf5\x70\xeb\x96\x3b\xff\xd7\x5f\x7b\xbb\x35\x79\x62\xb4\xf5\xed\x42\xaf\xb7\xbb\x57\x93\x27\xc2\x78\x36\x92\xe8\x80\x1e\xc1\xf9\xe5\xd5\xc7\xd3\x8b\x93\xaf\x47\x67\x67\x57\x47\x67\x67\xe7\x5f\x81\x1a\xd8\x6a\x82\x00\xad\x22\x2f\x1e\x81\xd2\xf6\xf7\xa7\x93\xaf\x71\x71\x66\xa6\x3c\x86\x86\xad\xe6\x93\xfe\x84\xa3\xe3\xe3\x93\xcf\x5f\x80\x4e\x09\xd7\x0a\x09\x99\xe5\xa1\x8e\x4d\xb0\x93\x8c\xbb\x71\x59\x43\x61\x3a\xb3\x12\xd2\x83\x69\x89\xaa\xd5\x80\x50\x05\xa8\xc8\xea\x94\xb1\x02\x95\x07\xa6\x38\x28\xf4\x53\x6d\xc7\x10\xbc\x90\xc2\x0b\x74\x50\x68\x74\x20\x94\xd7\x60\x59\x86\x90\x69\xc5\x85\x17\x5a\xf5\x49\x0f\x44\x3e\xdf\x6c\x83\x72\x30\xc2\x5c\x5b\x04\xae\x1c\x08\x07\x63\xa5\xa7\x0a\xbc\x8e\x02\xe8\x32\x21\xa0\xe2\x10\x0c\x4c\x85\x2f\x01\x2b\xe3\x6f\xc0\x79\x2b\x54\x41\xa6\xa5\x90\x08\xdf\xbe\xc1\xd6\xf3\x52\x3b\xaf\x58\x85\x40\xf9\x0e\x0c\x87\x90\x24\xf0\xe3\x47\xc4\x1c\x9c\x44\x34\x70\x10\xbf\xc7\x63\xb7\x7b\x9e\xc2\xc3\xd2\xbd\x8c\xa7\x0d\x06\xea\xba\xe1\x0d\x66\x51\x3a\xec\xb2\x60\x25\xd0\x33\xd8\x2e\xbd\x37\x6e\x90\xa6\xd3\xe9\xb4\xcf\xad\x36\x23\x7d\xdd\xcf\x74\x95\xba\x34\x3f\x44\xf9\xcf\x6b\xfd\xaf\xfd\x57\x93\x57\x53\x6e\xd3\xf2\xc6\xa0\x1d\x87\x11\xfe\x9b\xcb\x61\xd4\x59\xd3\xc3\xe6\xab\x84\x38\xf4\xf0\x8f\x6b\x82\xd7\x0d\x61\x97\x47\x97\xff\xbd\x38\x1d\x6e\x2f\x94\xf6\x1f\xe6\x3c\xda\xae\xb2\xd6\x0e\x75\xbd\xdd\x6c\xa4\xd7\xb3\xbb\x68\x83\x02\x4a\x8d\x15\x13\x21\xb1\x40\x0e\x94\xda\x0a\x28\x9d\xb1\x14\x81\x02\x3a\x81\x74\x90\xc6\xaf\x83\x5f\x40\xb1\xcb\xf6\x20\x0e\xad\x00\x82\x21\x41\xc5\x84\xed\x0e\x42\x82\xe1\xcc\x23\xcd\x18\xf5\x36\x38\x1f\x8f\xc1\x81\x0a\xa0\x16\x21\x71\xbd\xe7\xb0\x1b\xfb\x05\xda\x01\xec\xf4\x77\x7b\xdf\x0f\x66\x80\xdd\x31\xb6\xd3\x4b\xda\xab\xaf\xad\x28\x84\x4a\xab\xe6\x98\xa9\x36\xa8\x5c\x29\x72\x4f\xdb\x85\x7e\x44\xa9\x15\xe7\x9f\xe7\x88\xca\x6a\x3e\x16\xa3\x92\xdb\x5b\x1a\xe5\xa9\x10\xb6\xfa\xef\x59\x36\x0e\xe6\xbd\xd4\xa3\x4f\x51\x55\x49\x02\x75\x4d\xa4\x2e\x0a\xb4\x40\x3d\xb4\x35\xd1\x0e\x93\xbe\x2b\x21\x99\x5f\x90\xd8\x17\x26\x68\x6f\x40\xab\x05\x61\xee\x24\xf1\x2a\x39\x1f\x55\x04\x05\xfa\x46\xe1\xa3\x26\x0b\x89\x04\x5d\xe4\xcb\x7d\x24\xdd\x25\x1e\x2b\x13\xeb\xf8\x20\xec\x70\xd9\xd6\xed\xab\xc6\x5c\x58\xd8\x5a\xf0\x23\x0f\xd7\xc8\xf5\x54\x49\xcd\x78\x2c\xb3\x8d\x91\x3c\xf2\x42\x9c\xf8\x8c\xb7\x98\xdc\x73\x27\x96\xe4\xb7\xae\xb8\xef\x04\x1a\xd5\xad\x11\x3d\x58\x5f\xda\xe4\x9c\x49\x1d\xb8\xb1\x7a\x22\x38\xda\x74\x90\x5e\x71\xe6\x59\x7a\xa5\xc3\x3c\xf4\x22\x0c\x83\x54\x87\x28\xed\x68\xfa\x9b\xb3\x40\x04\xb4\xc5\xa2\x8d\x44\x47\x52\x8f\x22\x63\xc3\xb8\x73\x45\x07\x75\xdd\x39\x71\x74\x5e\x28\x16\xbb\xda\x30\x26\xeb\x28\xe9\xf3\x51\xe7\xc0\xb2\x68\x83\x19\xe2\x0f\xf3\xd2\xe5\x9f\x39\x23\xbf\x13\xcb\xe1\xec\xb1\x79\x38\x42\xeb\x14\x79\x75\x8a\x19\x57\x6a\xff\x58\x66\xdb\xb6\x12\x31\xf9\x73\x66\x23\x96\x83\xf9\xb7\xb9\x69\x51\xbb\x83\xe5\xbf\x5a\x8e\x28\xc2\xc9\x97\xe3\x0f\xc7\x5f\xce\xae\x8e\x3e\x9f\x0e\x93\x17\xc9\x3d\xd4\x2d\x15\xdb\xf8\xc4\x28\xcd\x8c\xd2\x1d\x7b\x06\xd7\x92\x1e\xd6\xd8\x89\xea\xa1\xf1\xf2\x2c\xdf\x2b\x85\xd3\xce\xa1\x79\x4b\x16\x6e\x6f\xb7\x2c\x94\xf0\x82\x49\x9a\xc9\xd0\x28\x35\xe9\xa8\xd8\x6f\x7e\x86\xb3\xce\xb3\xb4\x3a\x38\x7c\xf1\x7a\x7f\x6f\x71\xe9\x60\xa3\xe3\xc1\xba\xe3\xe1\x46\xc7\xc3\xc6\x31\xd9\x5c\x12\xf5\x7a\x8c\xaa\x81\x85\xe6\xda\xd2\x66\x08\x5a\x71\x65\x7c\x82\xd6\x0b\x87\xd4\x20\x5a\x1a\xac\x74\xb0\xa1\x69\x36\x69\x08\xa9\x26\xeb\x28\xa5\xbb\x2b\x6b\x4d\x0f\xb3\xab\x3d\x2c\xe2\xb9\xd4\x9e\x96\x06\xa7\x95\xb8\x8f\x11\x38\x36\xaf\x78\xd2\xb4\xea\x38\x16\xd4\x35\x21\x3e\x28\xe4\x94\xf1\x0a\x8c\xd5\x79\x94\xfc\xdd\xdb\x91\x69\xe5\xad\x96\xd4\x48\x16\xdf\xed\x1e\x28\xed\x71\x00\xcc\xeb\x4a\x64\xf4\xce\xaf\x19\x35\x32\x1b\xff\xa3\x90\x5a\x1b\x07\x41\x79\x21\xbb\x3a\xe2\x60\x12\x0c\xb9\x9b\x87\x51\xc5\xd9\x68\x73\x94\xf9\x7c\xbc\x3a\x3e\x3f\xe8\x1d\xe7\xe2\x1e\x70\xe1\xd8\x48\x36\x6f\x88\xbb\x71\x52\x17\xe0\x84\xca\x9a\x19\xa8\x62\x8a\x15\x08\x18\x1f\x16\x5f\x46\x17\x5f\x5a\x1d\x8a\x12\x66\x83\xfb\x42\xc2\x36\x0e\xce\xa2\x6c\x2c\x49\x9b\x35\xf3\xff\x03\x00\x00\xff\xff\x64\xbb\xe1\x11\x58\x0d\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x61\x6b\x1c\x37\x14\xfc\xae\x5f\x31\x3e\x97\x5c\x4b\xd1\xad\x53\x68\x0a\x4e\xec\x52\x4a\x02\x81\xd2\x80\x4d\xe9\x87\x90\x0f\xba\xd5\xdb\x5d\xe5\xb4\x7a\x8a\xf4\x74\x7b\x8e\xb9\xff\x5e\x74\x6b\x27\xae\x1d\x2e\xdf\x96\xd3\xcc\xbc\xa7\xd1\xcc\x9d\x9e\x34\x6b\x17\x9a\xb5\xc9\x03\x34\xed\x94\x72\x1d\x4e\xd0\x27\x8a\x68\xb6\x26\x35\xde\xad\x1b\xcb\xed\x86\x12\x1a\x92\xb6\xe9\xb2\x98\xf5\x4b\xc8\x40\x41\x01\xf9\x26\x0b\x8d\xad\x78\x64\xe1\x88\x19\xb8\xca\x94\xb6\xae\x25\x05\x8c\x9b\x2e\xaf\x76\x5d\x86\xee\xd0\x58\xda\x36\xd6\xe5\x4d\x63\x3e\x97\x44\x4d\xa2\xcc\x25\xb5\xa4\xa3\x49\xf2\x5c\x01\xd4\x0e\x8c\xe5\x71\x18\x9e\x6c\x85\x2a\x8f\x3e\xc5\x4f\x85\xc5\x00\x67\x38\x5b\xe2\xf2\xf2\xeb\xb2\x75\x0d\x2e\x41\x1e\x33\x15\x90\x28\x0b\x27\x6a\x39\x40\x5f\x3d\x39\xbf\xbd\xd5\x70\x1d\xe8\x13\x56\x57\xec\x09\x0b\x17\xba\x64\x16\xd8\xef\x15\xd0\x1a\xc1\x3c\x64\x46\x37\xd6\xd0\xc8\x61\xf5\x31\x73\xc0\xab\x57\xcb\xd7\xef\xde\x2c\xd5\xad\x02\x16\x9e\x7b\x6d\x93\xdb\x52\x5a\x9c\x63\xf1\x91\x4b\x0a\xc6\xdb\x85\xda\xab\xd7\xef\xde\x1c\x86\x50\xb0\xb3\xe8\x43\x3b\x4d\x92\xc7\x7e\x76\x4e\xa9\xbb\x5b\xc7\xe2\x3d\x6e\x6f\xb1\xfa\x93\x43\xe7\xfa\xd5\xdb\xd1\xf4\x94\x57\x7f\xb3\x25\xec\xf7\x78\x76\x79\xb0\x31\x54\xd4\x33\xa5\x4e\x31\x0d\x14\x66\x51\x17\x7a\x84\x0a\x9b\x8c\xe9\x29\x08\x4c\xb0\x08\x24\x13\xa7\x0d\x8a\x38\xef\xc4\x51\x46\xcf\x94\xe1\x82\x30\x92\x69\x09\x2d\x07\xeb\xc4\x71\x58\xa9\xd3\x6a\xca\x3d\x39\x95\x90\xb1\xa6\x8e\x13\xc1\x86\x0c\x97\xb1\x09\x3c\x05\x08\xd7\x8c\xdc\x4d\xa2\xc3\x15\x4b\xc4\xe4\x64\x00\x8d\x51\x6e\x90\x25\xb9\xd0\xab\x69\x70\x9e\xf0\xfe\x3d\x7e\xf8\x71\xe0\x2c\xc1\x8c\x04\x6d\x7f\xc2\xc5\x05\x16\x0b\x7c\xf8\xf0\x12\x96\x91\x3d\x51\xc4\xf3\xfa\x1d\x48\xdd\x71\x4e\x70\xdc\x8b\xeb\x7a\xdb\x12\xb1\xdf\x57\x5e\x75\x77\x56\x51\xb3\x48\x5b\x92\x87\xfe\x0b\xcb\x41\x24\xe6\xf3\xa6\x99\xa6\x69\x65\x13\xc7\x35\xef\x56\x2d\x8f\x4d\x6e\xba\x5f\xc8\xff\xba\xe3\xdf\xce\x5e\x6c\x5f\x4c\x36\x35\xc3\x4d\xa4\xb4\x29\x6b\xfa\xdd\xfa\x8b\x9a\xb2\x43\x75\xbe\xfc\xaa\x54\x26\xc1\xcf\x3b\x45\xbb\xc8\x49\x70\xfd\xc7\xf5\x3f\x57\x6f\x2f\x96\x0f\x56\xfb\x97\xd3\x86\xd2\xdd\x66\xf3\x39\xf6\xfb\xe5\x81\xa8\x77\xf7\x8f\x9b\x4a\x80\xd6\x31\xb9\xad\xf3\xd4\x93\x85\xd6\x69\x84\xd6\xf7\xaf\x54\x8d\x82\xde\xa2\x39\x6f\xea\xe7\xf9\x67\x68\xba\x9b\x76\xd4\x87\x39\x00\x25\xaa\x12\xea\xc0\x99\xa1\x54\x89\xd6\x08\xe9\xd6\x68\x49\x25\x8b\x3a\x16\x7c\x29\x81\xac\x36\x76\x44\x4c\xdc\xd5\x67\xe0\x48\x21\x0f\xae\x13\xdd\x72\x90\xc4\x5e\x47\x6f\x02\xcd\xc1\xf6\x99\xbe\xc7\xaa\x09\x79\xd8\x02\x75\x8a\xc0\x42\xe7\x30\xc2\xa3\x6b\xf5\xff\x91\x68\x53\xfd\xa7\xf2\xcc\x31\xa3\x04\x71\x1e\xa3\xc9\x42\xa9\x26\xaf\x44\xf5\xb5\x41\x14\xcc\xda\xd3\xb7\x55\xbe\x34\xea\x71\xe1\x8e\xa2\xe7\x26\x59\x97\xcd\xda\xd7\x16\xa5\x7c\x93\x3d\xf7\xc8\x2e\xb4\x87\x90\x8f\x26\x98\x9e\x40\x5b\x4a\x37\x32\x54\x88\x0c\x89\x4b\x3f\xe0\xbe\xf5\x0f\x06\xce\x3a\x74\xaf\xf2\xcd\x95\x38\x3e\x39\xfe\x2f\x00\x00\xff\xff\x15\x92\x14\xfd\xb0\x05\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
