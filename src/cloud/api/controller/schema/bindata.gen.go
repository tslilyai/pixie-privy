// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema.graphql
package schema

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x57\x5f\x6f\xdb\x38\x12\x7f\xd7\xa7\x98\x22\x0f\xdb\x00\x89\x81\x5b\xec\x15\x07\xbd\x69\x6d\xa5\x15\x9a\xc8\xbe\x48\x4e\xaf\x57\x14\xc1\x58\x1c\x5b\x44\x24\x52\x25\xa9\x24\xbe\x43\xbf\xfb\x81\xa4\x24\x53\x49\x7a\xfb\x64\x6a\x38\x7f\x7e\x9c\xf9\xcd\x90\x3e\x83\xb2\xe6\x1a\xf6\xbc\x21\x60\xa4\x2b\xc5\x77\xa4\xc1\xd4\x04\xba\xaa\xa9\x45\xd8\x2b\xd9\xba\xef\x64\x93\x81\x26\xf5\xc8\x2b\x5a\x44\x67\xd1\x19\x64\xe6\x37\x0d\x42\x1a\xe0\x8c\xb0\xb9\x80\x5d\x6f\xe0\x89\x40\x10\x31\x30\x12\x5a\x14\x3d\x36\xcd\x11\x0e\x24\x48\xa1\x21\x30\xc7\x8e\x34\xec\xa5\x72\xfe\xca\x63\x47\x45\xa5\x78\x67\x60\x9b\x45\x67\xf0\x54\x93\x00\x33\x81\xe1\x1a\xfa\x8e\xa1\x21\xb6\xf0\x10\x2b\x14\xb0\x23\x60\x52\x10\xec\x8e\xa0\x7a\x21\xb8\x38\xc4\xd1\x19\xc0\x41\x61\x57\xff\x68\x2e\x3d\xe4\x4b\x17\xc7\x7b\x1e\x63\x5f\x1a\x3d\x1c\x68\x31\x28\xc3\xe5\xa5\xec\x4d\xd7\x9b\x51\xce\x16\x46\x3b\x18\xbc\xaa\xe1\x89\x37\x4d\x00\xbc\x26\x18\x94\xad\x6f\x0f\xd0\xd4\x68\xbc\xde\x8e\xa0\xe3\xd5\x03\x31\xe8\x3b\x0b\xcd\xaa\x6f\xb3\x45\x34\xe4\x36\xf0\xef\x2c\x35\xe8\x5a\xf6\x0d\x03\x7a\xe6\xda\x00\x17\x3e\xdd\xd8\x12\x30\xae\xa8\x32\x52\x1d\x01\xc3\x22\x4c\x98\xad\xf9\x22\x8a\x86\xd2\xfc\x37\x02\xf8\xd1\x93\x3a\xc6\xf0\x4f\xfb\x13\x01\xb4\xbd\x41\xc3\xa5\x88\xe1\x66\x58\x45\x3f\xa3\xc8\x81\xde\x6a\x52\x99\xd8\x4b\x67\xc6\x59\x0c\xd9\xea\x5d\x04\x20\xb0\xa5\x18\x0a\xa3\xb8\x38\xd8\x6f\x6a\x91\x37\xa1\xa0\xe3\x95\xe9\xd5\x4c\x47\xaa\x43\x3e\x33\xfb\x19\x45\x24\xfa\x16\x12\x65\xf8\x1e\x2b\x63\x6b\xeb\xe2\x00\x24\xe5\xfd\x36\xff\x9c\xaf\xbf\xe4\xe3\xe7\x75\x96\x6f\xff\x75\x9f\xdc\xac\x3e\xfc\x31\x8a\x56\xc9\xed\x97\x2c\x9f\xcb\x96\xeb\xbc\x4c\xb2\x3c\xbd\xbd\x2f\xd2\xf2\xfe\x6b\x72\x73\x5d\xbc\xbd\x15\xfa\x9b\x80\xf4\x46\x56\xb2\xed\x1a\x32\x94\x54\x36\x0f\x13\xa4\x64\x86\xe8\x0c\x12\x01\xc4\xb8\x01\x74\x6a\x20\xab\xaa\x57\x1a\xf8\x1e\x10\x7a\x4d\x0a\x6a\xd4\xd0\x4a\xc6\xf7\xdc\xf2\xba\x26\xe0\xc2\x11\x81\x9e\x8d\x2d\x36\x17\x9a\x94\xe1\xe2\x00\x52\x01\xa3\x86\xdc\xba\xaa\x51\x61\x65\x48\xe9\x85\x0b\xe2\x88\xc0\x45\xd5\xf4\xcc\xb6\xd7\xb1\x73\x06\xbe\xf2\x0f\x74\xdc\x49\x54\x0c\x50\x30\xe8\x50\x7b\x07\xb2\x6d\x51\x30\x67\x6e\x11\xa7\xab\xac\xf4\x70\x41\x53\x43\xd5\x09\xaf\x68\x8e\x6f\x83\xae\x6a\xa9\x49\x00\x0a\xc0\x20\x1b\xa0\xfb\xc3\x81\xb4\xb5\x5d\x8c\xb0\x18\xaf\xd0\x58\x5c\xd2\x85\xb0\xa0\x66\x26\x8e\xea\xdc\x8c\xbc\x6d\xe5\xa3\xef\x09\x1b\xea\x37\x0d\x36\xb6\x6d\x6a\xe9\x84\xc2\x26\x06\xbb\x4e\xc9\x4e\x71\xd7\x3d\xb8\x1b\x4f\x51\xa4\xd7\xe9\xb2\x7c\xb3\x4a\xa9\x30\xdc\x1c\x3f\x73\xc1\x7c\x95\xd2\xcf\x41\x95\xec\xd7\x66\xbd\x1a\x56\xc5\xdd\x72\x5c\x2d\x6f\xb3\x4d\x39\x7c\xe4\xc9\x4d\x5a\x6c\x92\x65\x3a\x51\xde\x75\x85\x73\x67\x91\xc6\x53\x0b\x58\x0e\x57\x4d\xaf\x8d\x15\x2e\xfd\xe2\x85\x7c\x29\x85\xa0\xca\xb7\xd2\xf2\xa5\xe8\xa4\xcb\x47\xc6\xbf\xc7\x80\xfa\xf1\xac\x11\xce\x63\x58\x5e\x67\xa3\xc4\xda\x8d\xba\x7a\xb2\x0a\xdb\xe9\xfc\x64\xae\xc7\x48\xf4\x6c\x2c\x9f\xee\xb8\xbe\xea\x45\xe5\xc4\xef\xfd\x88\x9b\x9a\xf0\x02\xf6\xbd\xa8\xac\x23\x1d\xc3\xb7\x41\xf8\xfd\x3c\x86\xd0\xc8\x45\x0f\xb2\xfe\xde\x91\x79\xf4\x71\x31\x94\x72\x23\x75\x0c\x99\x30\x17\x03\xc9\xe2\x5f\xf4\xd3\xf9\x7c\xe3\x96\x74\xdf\x98\x77\x53\xfa\x5f\xef\xb9\x5a\xd8\x6e\x20\x7b\x0d\xb4\x68\x0c\xb1\xa1\x9f\xb8\x0e\x9a\x4b\x0f\x70\xfc\x30\xb6\x64\xde\x11\x09\xe8\x50\x69\x62\xe3\x88\x9d\x53\x54\x4e\x3c\xf6\x1c\xc6\x5d\x61\x64\x07\x9d\xd4\xdc\xc2\x75\x8d\x34\xc5\xcc\xc2\x53\x3b\xfd\x2f\x35\x99\x9a\xd4\x2b\x0c\x16\x17\xc2\x23\x36\x9c\x5d\x00\x3d\x53\xd5\x1b\xdc\x35\x34\xf6\xa7\xf5\xca\x75\x3a\xc9\x63\xf8\x53\xca\x86\x50\xf8\x5e\x6d\x9a\xa0\xdd\xfc\xd5\x47\x58\xd5\x20\xf7\x2e\xd0\x00\xd2\x61\xb3\xeb\x93\x6a\x0c\xdf\xca\x50\xf0\x7d\x4a\xea\x4c\x1c\xe4\x93\x0b\x46\xcf\x81\x63\xdf\xb4\xa6\x26\x4d\x33\x0c\xa8\x5c\xee\x87\x90\x99\xb5\x72\xb5\x9e\x65\xc1\x8f\x18\x7b\x7c\x0c\x8c\x87\xab\xdb\x56\x0a\x77\x43\x40\x77\x01\xb6\xf8\xe0\x07\xc2\x90\x95\x20\x51\x0b\x47\xde\xf1\x2b\xd9\x1b\x52\x85\x73\x1e\x66\x4a\xcf\x0e\x1e\xd2\xe6\xad\x0c\xbc\xbd\xef\x52\x61\xc7\x47\xfc\x8b\xb1\x12\x01\x84\x5d\x16\x01\xac\x86\x57\x82\x63\xf8\x20\x1d\x47\xd3\xd0\xf3\x85\x41\xd3\xeb\x20\xcd\x8c\xf6\x68\x89\xac\x8d\x1d\x6d\x7c\x6f\x1f\x40\xf5\xc0\x93\x07\x21\x9f\x84\x3d\xf0\xdd\xbf\xef\x8b\xf9\x25\x63\x4d\x07\x13\x0d\x35\x61\x63\xea\xa3\xb5\xae\x09\x95\xd9\x11\x1a\x5f\x18\x45\x15\xf1\x47\x62\xf6\x6a\x50\x74\xe8\x1b\x54\xc0\x85\x21\xf5\x88\x8d\x76\xf7\x83\xa9\x3d\xbf\x87\x31\x65\xdd\x29\xd2\x9d\x14\xcc\x82\x30\x12\x14\xfd\xe8\x49\x1b\x7d\xc2\xf1\x29\x4d\xae\xcb\x4f\x5f\x5f\xe0\xf0\x2f\x1c\x69\x4d\x18\xd7\x95\x9f\x6d\xb6\x1b\x3d\x83\x3e\xde\x6e\x96\x50\x4d\x13\x0f\x76\x8a\xf0\x41\x2f\x9c\x83\x5a\x76\xe4\xfb\x15\xcd\x74\x61\x8c\x80\x9c\xdf\x4a\xb6\x04\x3b\xac\x1e\xec\xf5\xc4\x05\x39\xe8\x8a\x74\xdf\x5a\xa2\xc2\x80\xc8\x23\x39\x01\x5d\x65\xc5\x72\x9d\xe7\xe9\xb2\x4c\x57\x53\xb5\xef\xf8\x7f\xb8\x9b\xbd\x7b\x7e\x70\x75\xe8\x50\x6b\x53\x2b\xd9\x1f\xea\x54\x58\x52\xb1\x13\x95\x46\xa3\x60\xa6\xbf\x78\xee\x68\x57\xd0\x78\x5e\x5f\xbb\xd1\xa0\x36\x9f\xc6\x6a\xdc\xe8\x18\xae\x1a\x89\x6e\x5a\x3f\x06\x08\xe2\x19\x9e\x77\x2f\x03\xce\xaf\x08\x1f\xba\x4b\x18\x53\xa4\x75\xf8\x80\x32\xf2\x81\xc4\xec\xf9\xe4\xbc\x8c\x6f\x36\x67\xb8\x54\x84\x86\x96\x6f\x5c\x53\x11\xc0\xd6\x3d\x8d\x43\x2c\xef\x87\x02\x64\x2b\x77\xd4\x8b\xff\x97\xa7\xf3\x69\x15\x9c\xe0\x74\x45\x0d\x2f\xb7\x5e\xcd\xde\x81\x00\xba\xc6\xdf\xff\xfe\xe1\x35\xec\xd9\x6d\xe5\x0f\x6d\xa8\x75\xad\x3c\xec\x7c\x7f\xa5\xeb\xd4\x1e\x49\xe9\xa0\xf9\xdc\x95\x5a\xa3\x38\x50\x23\x0f\xb3\x74\xf1\x96\xb4\xc1\xb6\x0b\xea\x32\x3a\x5c\xc9\xca\x2b\xde\x60\x97\x0a\x33\xdc\xf7\xe3\x55\x18\x7a\x61\xa3\xe6\xeb\x13\xdc\x71\x5d\x74\x54\x0d\xa0\x0e\x68\x3f\x7e\xa9\xf5\x97\x81\x1e\xbd\x5e\x3c\x1a\x58\x07\x7e\xaa\xac\xd0\xe0\xf4\x10\x5d\x25\x65\x72\x5f\x7e\xdd\xa4\xc1\xa4\xf8\x73\xbd\xbe\x4e\x13\xbb\xca\xf2\xd2\xbd\x85\xb7\x59\x5e\xfe\xed\xf7\x7f\x44\x00\x57\xd7\xeb\xc4\xcb\x8a\xf2\x36\xcb\x3f\x46\x00\x65\x76\x93\x7e\xf8\x23\xb7\x0f\xe3\xd5\xf6\x36\x29\xb3\x75\xee\x3e\xc7\x29\x56\x50\x8b\xc2\xf0\x6a\x8a\xe9\xc6\x52\xb1\x49\x97\xd9\x55\x96\xae\xbc\x20\x5f\xe7\xa9\x5f\x25\x1f\xd3\xbc\xbc\xdf\x66\xc3\xc6\x76\x33\xae\x8a\xf4\xf6\x2e\x5b\xa6\xee\x89\x35\xa5\xe3\x4a\x24\xea\xe0\x0e\x98\x28\xdf\x9a\x2f\xff\x4b\xb0\xe1\xb8\xf1\x74\x70\xd7\x82\x01\xa8\x78\x06\xd1\xd9\xf8\x01\x7b\x87\x4d\x4f\xaf\x4b\x70\x8a\xe9\x02\xa2\x3a\x58\x9a\xcd\x90\xbc\xfb\xfe\x52\xfd\x2f\x0b\xb6\xb7\x6a\xbe\x64\x27\x57\xb3\xba\x4f\x0f\x28\xe7\x83\x05\x9c\x8b\xe1\xdb\x2b\x0a\x5a\x08\x13\x0d\xbc\xce\x0b\xee\x78\x8d\xfd\x08\x6f\x3a\x43\xb8\xff\x33\xfa\x5f\x00\x00\x00\xff\xff\x92\x0d\xa2\x0f\x9d\x0f\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 3997, mode: os.FileMode(436), modTime: time.Unix(1586464567, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
