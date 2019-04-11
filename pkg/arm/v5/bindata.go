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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x51\x6f\x13\x39\x10\x7e\xc6\xbf\x62\xd8\x14\x4a\x7b\x75\xb6\x2d\x3c\xa0\x40\x90\x4a\x29\x52\xa5\x1e\x45\x2d\x27\x1e\x00\x55\xce\x7a\x76\xd7\x64\xd7\xde\xb3\xc7\x49\x4b\xc9\x7f\x3f\xd9\x9b\xa4\x49\x9a\xe6\x2a\xc1\x43\xd8\xda\x9f\x67\xc6\xdf\xf7\xd9\x9e\xce\xd3\x74\xa0\x74\x3a\x10\xae\x04\x8e\xd7\x8c\xa9\x1c\x9e\x42\x61\xb1\x81\x74\x24\x6c\x5a\xa9\x41\x2a\x4d\x36\x44\x0b\x29\x52\x96\xe6\x8e\xc4\xe0\x0d\x50\x89\x9a\x01\xb8\x1b\x47\x58\x67\x54\x81\x23\xd3\x40\x0b\xec\x3a\xb4\x23\x95\x21\x03\xa8\x87\xb9\xeb\x5e\xe7\x0e\x78\x0e\xa9\xc4\x51\x2a\x95\x1b\xa6\xe2\x97\xb7\x98\x5a\x74\xc6\xdb\x0c\x79\x23\x2c\x1d\x30\x00\xcc\x4a\x03\xdb\x9b\x61\x70\xaf\x2a\x08\xe1\xa1\xb0\xcd\xbf\xde\x90\x00\xd8\x87\xfd\x6d\x78\xf7\xee\xae\xd8\x50\x86\xf1\x9a\x56\x57\x32\x00\x8b\x8e\x8c\xc5\xcc\x68\xe0\x17\x6b\xe6\x33\x41\xd0\x46\x6a\x87\x52\x29\xb0\x36\xba\xfb\xd3\x19\x0d\x6f\xdf\x6e\x9f\x9c\x7f\xdc\x66\xb7\x0c\x20\xa9\x4c\xc1\xa5\x55\x23\xb4\x49\x0f\x92\x9f\xc6\x5b\x2d\x2a\x99\xb0\x09\x3b\x39\xff\xb8\x42\x94\xb0\xb4\xca\x54\xae\x18\x9b\xee\xa7\xf1\x55\x05\xb7\xb7\xd0\x3d\x36\x3a\x57\x45\xf7\xb4\x16\x05\xba\xee\x27\x23\x11\x26\x13\x78\xfe\x2e\x12\xa4\x03\xea\xf9\x5a\xb5\x90\x32\xb9\x4e\xab\xb9\x16\xab\x0c\xbb\xcc\xa9\x83\xb4\xf2\x7a\x1f\x7e\xff\x06\xb2\x1e\x1f\x14\x63\x01\xba\x92\xb0\x95\x41\x62\x2e\x7c\x45\xee\x51\x32\x84\x75\x0f\x8b\x10\x67\x03\x2f\xb9\xb1\x20\x1d\x81\xd2\x40\x59\xb3\xf7\xfa\xd5\xab\x57\x6f\x40\x1a\xf6\xa4\xb1\x86\x4c\x7f\xeb\x56\x3a\x7a\xf6\x6c\x6f\x77\xc2\x9e\x34\xc6\x52\x3b\xd0\xe9\xec\xee\x4d\xd8\x13\xd5\x90\x18\x54\xe8\x80\x1f\xc1\xf9\xe5\xd5\xc7\xd3\x8b\x93\xaf\x47\x67\x67\x57\x47\x67\x67\xe7\x5f\x81\x37\xb0\x15\x83\x00\xaf\x83\x2e\x84\xc0\x79\xfb\xff\xa7\x93\xaf\x61\x70\x36\xcd\x65\x08\x0d\x5b\xf1\x97\xff\x84\xa3\xe3\xe3\x93\xcf\x5f\x80\x8f\x99\x34\x1a\x19\x9b\xe5\xe1\x4e\x8c\x70\x6a\x19\x77\xe3\xb2\x28\x61\x3a\x9b\x65\xac\x03\xe3\x12\x75\xeb\x01\xa5\x0b\xd0\x41\xd5\xb1\x10\x05\x6a\x02\xa1\x25\x68\xa4\xb1\xb1\x43\xf0\xa4\x2a\x45\x0a\x1d\x14\x06\x1d\x28\x4d\x06\xac\xc8\x10\x32\xa3\xa5\x22\x65\x74\x97\x75\x40\xe5\xf3\xc5\xd6\x6b\x07\x03\xcc\x8d\x45\x90\xda\x81\x72\x30\xd4\x66\xac\x81\x4c\x30\xc0\x34\x13\x02\x6a\x09\xbe\x81\xb1\xa2\x12\xb0\x6e\xe8\x06\x1c\x59\xa5\x0b\x36\x2e\x55\x85\xf0\xed\x1b\x6c\xbd\x28\x8d\x23\x2d\x6a\x04\x2e\x77\xa0\xdf\x87\x24\x81\x1f\x3f\x02\xe7\xe0\x2a\xc4\x06\x0e\xc2\x77\xd8\x76\xbb\xe6\x29\x6c\xb6\xee\x65\xd8\xad\x6f\x60\x32\x89\xba\xc1\x2c\x4a\xcb\x9d\x43\x82\xbf\xae\x19\x5e\x47\x6e\x2f\x8f\x2e\xff\xb9\x38\xed\x6f\x2f\x44\xf9\x5b\x38\x42\x3b\x0d\xd2\xce\xc3\x64\xb2\x1d\x17\xf2\xeb\xd9\xb1\xb1\x5e\x03\xe7\x8d\x55\x23\x55\x61\x81\x12\x38\xb7\x35\x70\x3e\x23\x34\xec\x09\xf8\x08\xd2\x5e\x1a\x3e\x7b\xbf\x80\xe3\x34\xdb\xc6\x92\x99\xd7\x21\x51\x8b\x64\xcc\x37\x52\x10\xf2\x4c\x70\xb2\xde\x11\x63\x2e\xa4\x52\xc0\x2d\x42\xe2\x3a\x2f\x60\x37\x1c\x69\xb4\x3d\xd8\xe9\xee\x76\xbe\x1f\x94\x44\x8d\xeb\xa5\xe9\x1d\xa9\x3b\x9d\xa4\x3d\x9d\xc6\xaa\x42\xe9\xb4\x8e\xdb\x4b\x4d\x83\xda\x95\x2a\x27\xde\x0e\x74\x87\x7e\x80\xad\x7f\xfe\x3c\x47\x10\x3f\xfe\x2c\x46\x65\xb7\xb7\x3c\x38\x48\x23\x6c\x75\xdf\x8b\x6c\xe8\x9b\xf7\x95\x19\x7c\x0a\xc2\x27\x49\xd8\x7a\x65\x8a\x02\x2d\x70\x82\xb6\x26\xee\x5a\x5a\xba\xae\x84\x64\xee\xe1\x70\x74\x47\x68\x6f\xc0\xe8\x05\xef\xec\x24\xc1\xed\x8e\x82\xd0\x50\x20\x45\x13\x0e\x62\x16\x16\x84\xb9\xc8\x97\x8f\x7a\xba\xcb\x08\xeb\x26\xd4\xf1\x41\xd9\xfe\xf2\xdc\x74\x5d\x3d\x94\xca\xc2\xd6\x02\x8e\x6d\xae\x51\x9a\xb1\xae\x8c\x90\xa1\xcc\x36\x46\xf2\x48\xcf\x9e\x50\x26\x5b\x4e\x1e\xb0\xed\x92\xed\xee\x3b\xed\x3b\x83\xe8\xb6\x7b\x42\xf7\xee\x0f\xad\x03\x67\x95\xf1\xb2\xb1\x66\xa4\x24\xda\xb4\x97\x5e\x49\x41\x22\xbd\x32\x7e\x1e\x7a\x91\x86\x5e\x6a\x7c\xb0\x74\x98\xfa\x9f\xbd\xb4\xcb\x07\x95\x19\x04\x95\xfa\x01\xbd\xa2\xfd\x0c\x23\xd1\x91\xd2\x22\xdc\x35\xfd\x10\x7f\xaa\x42\x57\x0e\x60\xc6\xeb\x66\xf6\x5b\xfc\x1c\x8c\xf2\xce\x12\x87\xb3\x5b\x7f\x73\x84\x16\x14\xd4\x73\x5a\x34\xae\x34\xf4\x58\xfd\xda\x4b\x23\xec\xfc\xcf\xf5\x0b\x16\xec\xcd\xbf\xe6\x53\x8b\x0e\xed\x2d\xff\xd5\x2a\xc1\x11\x4e\xbe\x1c\x7f\x38\xfe\x72\x76\x75\xf4\xf9\xb4\x9f\xbc\x4c\x1e\x10\x68\xa9\xd8\x88\x09\x51\x62\xb3\x30\xdd\xf6\x8c\xae\x25\xd5\x17\x04\x89\xe9\x78\xf0\x08\x0f\x47\x64\xf9\xf4\x68\x1c\x4f\x01\xf1\x52\x5f\x38\xa3\xd3\x61\xa5\x15\x29\x51\xf1\xac\xf2\xd1\x8f\xc9\x54\x8a\xfd\xf8\xaf\x3f\xbb\x5f\x96\x46\x7b\x87\x2f\x5f\xef\xef\x2d\x0e\x1d\xac\x05\x1e\xdc\x07\x1e\xae\x05\x1e\x46\x60\xb2\xbe\x24\x4e\x66\x88\x3a\xd2\xc2\x73\x63\x79\xec\x46\x56\xa0\x42\x8e\xd0\x92\x72\xc8\x1b\x44\xcb\xbd\xad\x1c\xac\xb9\x1a\x63\x1a\xc6\xea\xd1\x7d\x96\xd2\xdd\x95\xb1\x78\x53\xd9\xd5\x9b\x2a\xf0\xb9\x74\x09\x2d\x75\x30\x2b\x71\x1f\x63\x70\x8c\xcf\x69\x12\x2f\xe4\xf0\x3e\x4f\x26\x8c\x91\xd7\x28\xb9\x90\x35\x34\xd6\xe4\xc1\xf2\x77\x2f\x44\x66\x34\x59\x53\xf1\xa6\x12\xe1\x15\xee\x80\x36\x84\x3d\x10\x64\x6a\x95\xf1\x3b\x5c\x7c\xf3\x33\x1b\x5a\xfb\xca\x98\xc6\x81\xd7\xa4\xaa\x69\x1d\xa1\x43\xf0\x0d\xbb\x6b\x4c\x51\x87\x26\x65\x7d\x94\x79\xa3\xba\xda\xc7\x6e\x44\x87\x06\xb5\x03\x52\x39\x31\xa8\xe2\x4b\xe1\x6e\x5c\x65\x0a\x70\x4a\x67\xb1\x19\xa9\x85\x16\x05\x02\x86\xe7\x83\xca\x00\xa1\xd2\x1a\x5f\x94\x30\xeb\xa0\x17\x12\xb6\x71\x70\x16\x65\x6d\x49\xa6\xb9\x37\xfd\x5f\x00\x00\x00\xff\xff\x98\x50\xee\x8a\xe1\x0c\x00\x00")

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

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x4f\x8f\xdb\xb6\x13\xbd\xf3\x53\xbc\xf5\xfe\x10\xff\x8a\x82\x56\x72\xdd\x64\x17\x28\x8a\x04\xc8\xa5\x01\x76\x51\xf4\x10\xe4\x40\x4b\x23\x89\x31\x35\xa3\x90\x43\xdb\x1b\x43\xdf\xbd\xa0\x65\x37\xdb\xdd\xc0\xbd\x09\xe0\xfb\x33\xf3\xf8\xa8\xeb\xab\x6a\xed\xb9\x5a\xbb\xd4\xc3\xd2\xde\x18\xdf\xe2\x0a\x5d\xa4\x11\xd5\xd6\xc5\x2a\xf8\x75\xd5\x48\xbd\xa1\x88\x8a\xb4\xae\xda\xa4\x6e\xfd\x16\xda\x13\x1b\x20\x3d\x26\xa5\xa1\xd6\x80\xa4\x32\x62\x06\xae\x12\xc5\xad\xaf\xc9\x00\xc3\xa6\x4d\xab\x7d\x9b\x60\x5b\x54\x0d\x6d\xab\xc6\xa7\x4d\xe5\xbe\xe7\x48\x55\xa4\x24\x39\xd6\x64\x47\x17\xf5\x8d\x01\xa8\xee\x05\xcb\xcb\x30\xbc\x98\x0a\x45\x1e\x5d\x1c\xbf\x65\x51\x07\xbc\xc6\xeb\x25\xee\xee\x7e\x0c\x5b\xc6\x90\xcc\xfa\x9c\x69\x80\x48\x49\x25\x52\x2d\x0c\x7b\xff\xe2\xfc\x70\xb0\xf0\x2d\xe8\x1b\x56\xf7\x12\x08\x0b\xcf\x6d\x74\x0b\x4c\x93\x01\x6a\xa7\x98\x4d\x66\x74\xd5\x38\x1a\x84\x57\x5f\x93\x30\xde\xbd\x5b\xbe\xff\xf4\x61\x69\x0e\x06\x58\x04\xe9\x6c\x13\xfd\x96\xe2\xe2\x06\x8b\xaf\x92\x23\xbb\xd0\x2c\xcc\x64\xde\x7f\xfa\x70\x34\x21\x6e\x66\xd1\xa7\x71\xba\xa8\xcf\xf3\x6c\xbd\x31\xa7\xad\xc7\x1c\x02\x0e\x07\xac\x7e\x17\x6e\x7d\xb7\xfa\x38\xb8\x8e\xd2\xea\x0f\x69\x08\xd3\x84\x57\x77\xc7\x18\xb9\xa0\x5e\x19\x73\x8d\x5d\x4f\x3c\x8b\x7a\xee\xc0\x05\xb6\x73\xae\x23\x56\x38\x6e\xc0\xa4\x3b\x89\x1b\x64\xf5\xc1\xab\xa7\x84\x4e\x28\xc1\xb3\x0a\xa2\xab\x09\xb5\x70\xe3\xd5\x0b\xaf\xcc\x75\x09\xe5\x4c\x8e\x99\x13\xd6\xd4\x4a\x24\x34\x9c\xe0\x13\x36\x2c\x3b\x86\x4a\xe9\xc8\xc9\x89\x8e\x2b\xe6\x11\x3b\xaf\x3d\x68\x18\xf5\x11\x49\xa3\xe7\xce\xec\x7a\x1f\x08\x9f\x3f\xe3\x7f\xff\xef\x25\x29\xbb\x81\x60\x9b\x5f\x70\x7b\x8b\xc5\x02\x5f\xbe\xbc\x45\x23\x48\x81\x68\xc4\x9b\xf2\xcd\x64\x4e\x9c\x2b\x5c\xce\xe2\xa1\x6c\x9b\x47\x4c\x53\xe1\x95\x74\x67\x15\x73\x14\x49\xa4\xf8\x75\x6f\x68\x3f\x4a\x54\x3c\xfc\xf6\xf0\xe7\xfd\xc7\xdb\xe5\x13\x95\xbf\x24\x6e\x28\x9e\x44\xe6\x73\x4c\xd3\xf2\x48\xb4\xfb\xf3\x3d\xc4\xcc\xb0\x76\x8c\x7e\xeb\x03\x75\xd4\xc0\xda\x38\xc0\xda\x73\xa0\x65\x27\xd8\x2d\xaa\x9b\xaa\x7c\xde\x7c\x87\xa5\x93\xdb\xc5\x91\x4d\xe6\x62\x34\x23\x8d\xc9\x63\xe3\x94\x6c\xed\xac\xc6\x9c\xd4\x5c\xea\xa6\x66\xa6\xc6\xba\x66\xc0\x18\xa5\x2d\x49\xc9\x48\x9c\x7a\xdf\xaa\xad\x85\x35\x4a\xb0\x63\x70\x4c\x73\xf7\x42\xa2\xff\x62\x95\x4b\x7c\x5a\x54\x73\x0d\x16\xa5\x1b\x38\x95\xc1\xd7\xf6\xdf\x48\xd4\xb1\xfc\x4c\x82\xc8\x98\x90\x59\x7d\xc0\xe0\x92\x52\x2c\xe5\xc8\xa3\xf9\x51\x72\x62\xb7\x0e\xf4\x73\x95\x7f\x4a\xff\xfc\x4d\x5c\x44\xcf\x65\x6f\x7c\x72\xeb\x50\x8a\x1e\xd3\x63\x0a\xd2\x21\x79\xae\x8f\x3d\x1c\x1c\xbb\x8e\x40\x5b\x8a\x8f\xda\x17\x88\xf6\x51\x72\xd7\xe3\xfc\x30\x9f\x18\xce\x3a\x74\x56\xf9\xe9\x48\x32\xbe\x38\xfe\x3b\x00\x00\xff\xff\x6b\x17\x8a\x82\x53\x05\x00\x00")

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