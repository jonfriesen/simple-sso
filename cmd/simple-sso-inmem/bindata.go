// Code generated by go-bindata.
// sources:
// ../../templates/footer.html
// ../../templates/header.html
// ../../templates/login.html
// DO NOT EDIT!

package main

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

var _TemplatesFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcb\xcf\x2f\x49\x2d\x52\x52\xa8\xad\xe5\x52\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x56\x48\xcd\x4b\x51\xa8\xad\x05\x04\x00\x00\xff\xff\x62\xd2\xa2\x66\x31\x00\x00\x00")

func TemplatesFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesFooterHtml,
		"../../templates/footer.html",
	)
}

func TemplatesFooterHtml() (*asset, error) {
	bytes, err := TemplatesFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/footer.html", size: 49, mode: os.FileMode(436), modTime: time.Unix(1573314683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x5d\x6f\xda\x30\x14\x7d\x2e\xbf\xe2\x2a\x68\xd2\x26\x91\x92\xa4\x49\x45\x3d\x86\xc4\xd8\xd0\x1e\xb6\xa7\xed\x6d\xea\x83\x83\x4d\x62\x61\xec\xc8\x76\x4a\x28\xe2\xbf\x4f\xce\x17\x1f\x4d\xd6\x4d\x9b\xac\x58\xc1\xf7\x72\x73\x7d\xce\xf1\xf1\xe1\x00\x84\xae\x99\xa0\xe0\xa4\x14\x13\xaa\x1c\x38\x1e\x07\xd3\xd4\x6c\xf9\x6c\x00\x30\xb5\x8b\xf6\x05\x60\x6a\x98\xe1\x74\xf6\x55\x26\x4c\x4c\xc7\xd5\x8f\x41\x15\xd1\x66\xcf\x29\x98\x7d\x46\x3f\x38\x86\x16\x66\xbc\xd2\xda\x99\x0d\x6e\x6c\xf0\x56\x6f\xb1\x32\x6e\xa2\x28\x15\x70\xa8\xd6\xb6\x58\x25\x4c\xb8\x9c\xae\x0d\xc2\xb9\x91\xef\x2f\x96\x15\x4b\xd2\x66\xbd\x09\x14\xee\x8e\x11\x93\x22\x88\x3c\x2f\x2b\xea\xfc\x18\xaf\x36\x89\x92\xb9\x20\x08\x86\xcb\x89\x1d\x75\x24\xc3\x84\x30\x91\x20\xb8\xf3\xb2\xa2\x9a\x82\xe6\xad\x4e\x59\x4b\x61\x10\xf8\x41\x56\xc0\x5c\x31\xcc\x47\xf0\x85\xf2\x27\x6a\xd8\x0a\x8f\x40\x63\xa1\x5d\x4d\x15\x5b\xd7\xd9\x2b\xc9\xa5\x42\x30\xbc\xbf\xbf\x6f\x3e\x2e\x15\xa1\xca\x55\x98\xb0\x5c\x23\x88\xda\xc2\xee\x8e\xc6\x1b\x66\xdc\xfe\x84\xad\x7c\xee\x8e\x1e\x07\x37\x37\x17\x78\xa5\x7e\x03\x59\xd5\x6e\x10\x66\x05\x38\x3f\x14\x8d\xf3\x55\x4a\x0d\x7c\xfb\xee\x8c\xfe\xa4\xfd\x16\x8f\x12\x85\x16\x8e\xf0\x04\x07\x61\x3a\xe3\x78\x8f\x20\xe6\x72\xb5\xb9\x20\x04\x81\x5b\x22\x58\xcd\x7e\xfb\x7a\x05\xcd\x72\xb9\xec\xe2\xe5\xe1\xd3\x22\x8c\x9a\x88\x15\x87\xab\x53\x4c\xe4\x0e\x81\x6f\xab\xd5\xcf\xf0\x21\xb4\xa3\x17\xdb\xf2\xa9\x7b\x7f\x15\xe7\xae\xe4\x6e\xcc\x3b\x32\xeb\xa4\x58\x1a\x23\xb7\xc8\xf6\xa6\x25\x67\x04\x86\x93\x87\xf9\x32\x5c\x58\x49\x76\xd0\x34\xd3\x19\x6e\xe5\xdd\x89\xa5\x25\xd0\xd5\xec\x99\x22\xf0\xa3\x0e\xf0\x16\x8b\xc5\xa2\x94\xc0\x75\x71\x8e\x63\xca\x7f\x5b\xba\xa1\xa9\xe1\xb6\x47\x4c\x65\xa1\x8b\x46\xd7\x5c\x62\x83\xc0\x9e\xc3\xcb\x23\x68\x64\x86\x4a\xa6\xaf\xda\x8c\x3e\xdb\xd1\x55\x9c\x89\x2c\x37\x3f\x4f\x16\xe0\x3c\x8e\xa0\x37\x21\xc3\x5a\xef\xa4\x22\xce\x63\xd3\x49\xfb\x81\x28\xaa\x3f\x99\xd2\xd2\x06\xce\x8f\x2c\x67\x82\xba\xf5\xfa\x19\x86\xb5\x33\xf8\x9e\xf7\xe6\x5a\xee\xde\x89\xde\xf3\xed\x9c\xef\x32\xb8\x62\xbe\xd2\x65\xcd\x79\xb9\xdd\xa8\xd3\x6e\x3e\xda\x51\x47\x64\x6e\x6c\x6f\x08\xbc\x17\xda\x2c\x5a\xb9\x33\xa1\xa9\x69\x05\x6f\x7d\x47\x25\x31\x7e\x1b\xdc\x4d\x46\x70\x9a\xbc\xdb\xe0\x5d\xdb\xcf\x3f\xfc\xb9\xf2\x0b\x21\xd5\x16\x73\xf0\xc3\xac\x18\xdb\xe9\x35\xaf\x78\x41\xeb\x6d\x9c\x1b\x23\x5b\xc1\x9c\x20\x70\x1b\xc6\x2e\xce\xf7\xff\x31\xc5\xde\xa4\x86\x21\x21\x05\xbd\xa6\xba\xf4\xa5\x20\x6a\x1c\x2a\xe8\x3a\x63\x7f\x61\x43\x7d\x48\xa0\x54\x3e\x51\xd5\x8b\xc7\x70\xe2\xcd\x83\x70\x5e\x16\x00\x98\x8e\xcb\x9b\xb1\xbc\x48\xc7\xcd\x4d\x3a\x8d\x25\xd9\xcf\x06\x87\x03\x50\x41\xe0\x78\xfc\x15\x00\x00\xff\xff\x42\x9f\x78\x8b\x82\x07\x00\x00")

func TemplatesHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesHeaderHtml,
		"../../templates/header.html",
	)
}

func TemplatesHeaderHtml() (*asset, error) {
	bytes, err := TemplatesHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/header.html", size: 1922, mode: os.FileMode(436), modTime: time.Unix(1573314683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xd2\x4f\x8b\xd5\x40\x0c\x00\xf0\xf3\xf6\x53\x84\x1c\xbc\xf9\xca\x5e\xb5\xd3\x9b\x0b\x82\xe8\x4a\x51\xf0\x24\xd3\x4e\xde\xeb\xc0\xfc\x33\x33\xb3\xba\x94\x7e\x77\xe9\xbf\xf7\xac\xba\xed\xad\xc9\x24\xf9\x11\x32\x0c\x90\xc8\x06\x23\x13\x01\xf6\x24\x15\x31\xc2\x38\x16\x45\x75\xf6\x6c\x41\x76\x49\x7b\x27\xb0\x8c\xd1\x23\x58\x4a\xbd\x57\x02\x83\x8f\x09\xa1\x33\x32\x46\x81\xd1\x4a\x4e\xaf\x2f\x4c\xe4\xb0\x2e\x00\x00\xaa\xfe\xbe\x6e\x9a\x4f\xf0\xc1\x5f\xb4\x83\x07\xcf\x76\x0e\xdf\x0d\x03\xe8\x33\x9c\xde\x31\x7b\x9e\x66\xc0\xfa\xdd\x55\x31\x48\x57\xbf\x77\x4f\xd2\x68\x05\x1d\x93\x22\x97\xb4\x34\xf1\x54\x95\x73\xea\x5a\x4f\x4e\x6d\x95\x55\xd9\xdf\xaf\xf3\x8c\x6c\xc9\xd4\xd7\x7e\x4b\xbb\x2f\x91\x18\x3e\x4a\x4b\xf0\xe6\xcf\x2e\xf3\x03\xed\x42\x4e\xa0\x95\x40\x27\x2d\x21\xa4\xe7\x40\x02\x13\xfd\x4a\x08\x53\x44\x60\x8e\xc4\x4b\x2e\x18\xd9\x51\xef\x8d\x22\x16\xf8\xcd\x67\x86\x87\x6c\xcc\xdc\x19\xa1\x5c\x05\xe5\x8d\x70\x40\x7a\x94\x31\xfe\xf4\xac\x0e\x45\x64\xa5\x36\x1b\x29\xac\x15\x1b\xeb\xf6\xbf\x63\x7d\x9d\x17\xf7\x78\x4d\xfe\xc5\x5a\x4c\x7b\xd4\xba\xf4\x57\xae\x8d\xe1\xed\x01\xe7\x47\x26\x7e\xfe\x1e\x13\x6b\x77\xd9\x54\xbd\x56\x8a\xdc\x66\xda\xbf\x78\x92\x26\x93\x18\x06\x38\x7d\x9e\xe2\xcd\x1c\x86\x71\xfc\xc7\xb4\xad\xea\x85\x5d\x1d\xc2\x16\x46\xcc\xad\xd5\xb7\x43\x6c\x73\x4a\xde\x6d\x02\x6c\xd6\xec\xff\xe6\x56\xe5\x74\xdd\x75\x51\xec\xce\xff\xec\x7d\x5a\xcf\xff\x77\x00\x00\x00\xff\xff\xf0\x13\x98\x77\x17\x03\x00\x00")

func TemplatesLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesLoginHtml,
		"../../templates/login.html",
	)
}

func TemplatesLoginHtml() (*asset, error) {
	bytes, err := TemplatesLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../templates/login.html", size: 791, mode: os.FileMode(436), modTime: time.Unix(1573314683, 0)}
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
	"../../templates/footer.html": TemplatesFooterHtml,
	"../../templates/header.html": TemplatesHeaderHtml,
	"../../templates/login.html": TemplatesLoginHtml,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"templates": &bintree{nil, map[string]*bintree{
				"footer.html": &bintree{TemplatesFooterHtml, map[string]*bintree{}},
				"header.html": &bintree{TemplatesHeaderHtml, map[string]*bintree{}},
				"login.html": &bintree{TemplatesLoginHtml, map[string]*bintree{}},
			}},
		}},
	}},
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

