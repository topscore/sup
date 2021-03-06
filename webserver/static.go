// Code generated by go-bindata.
// sources:
// templates/bottom.html
// templates/config.html
// templates/home.html
// templates/top.html
// DO NOT EDIT!

package webserver

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesBottomHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xca\x2f\x29\xc9\xcf\x55\x52\xa8\xad\xe5\x52\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\x02\xaa\x4b\xcd\x4b\x01\xca\x01\x02\x00\x00\xff\xff\x8a\xda\x87\x4b\x31\x00\x00\x00")

func templatesBottomHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesBottomHtml,
		"templates/bottom.html",
	)
}

func templatesBottomHtml() (*asset, error) {
	bytes, err := templatesBottomHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/bottom.html", size: 49, mode: os.FileMode(436), modTime: time.Unix(1441678318, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesConfigHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\xcb\x4e\xc3\x40\x0c\x3c\x27\x5f\x61\xf9\x03\x9a\x0a\x2e\x1c\x36\x7b\x29\x9c\x41\x2a\x3f\xb0\xd9\x38\x74\xa5\xec\x43\x1b\xa7\x80\xa2\xfe\x3b\xde\xb4\x55\x01\xf5\x66\x8f\xc6\x33\xe3\x59\x16\xe8\x69\x70\x81\x00\x6d\x0c\x83\xfb\x40\x38\x9d\xea\x6a\x59\x98\x7c\x1a\x0d\x0b\xce\x31\x21\x6c\x04\xad\x2b\x75\x78\xd0\x2f\xbd\x63\xd8\xad\x5c\xd5\xc8\x5e\x17\x36\xb8\x01\x36\xd3\x6c\x2d\x4d\xd3\x2a\x50\xa9\x04\x76\x34\xd3\xd4\xe2\x05\x46\x2d\xb4\x5f\x1c\xd5\x24\xbd\x9e\x52\xe8\xcb\x49\x91\x1f\x62\xf6\x60\x2c\xbb\x18\x5a\x6c\xae\x81\x3c\xf1\x21\xf6\x2d\xbe\xbd\xee\xdf\x51\x17\x6d\xa6\x2f\x36\x99\x0c\x04\xe3\xa9\xbd\x24\x7f\x36\x6c\x10\x6c\x1c\xc5\xf3\x69\x8b\x90\xe3\xa7\x4c\x8f\xdb\xb3\xf1\x8d\xb3\x7a\x5f\x15\xb4\xea\xb2\xae\x01\x44\xd4\x85\x34\x33\xf0\x77\xa2\x92\xb9\xf3\x8e\x11\x8e\x66\x9c\x65\xdd\x9b\x23\xe1\x4a\x53\x4d\xc9\x28\x4f\x97\x59\xbe\xdf\xcd\x39\x53\xf8\x5f\x88\x4a\x99\xee\xb9\x16\xb8\xfe\xdb\x6e\x17\x99\xa3\x3f\x17\x7c\x2b\xe3\x27\x00\x00\xff\xff\x30\x37\x14\xec\x97\x01\x00\x00")

func templatesConfigHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigHtml,
		"templates/config.html",
	)
}

func templatesConfigHtml() (*asset, error) {
	bytes, err := templatesConfigHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config.html", size: 407, mode: os.FileMode(436), modTime: time.Unix(1442355067, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesHomeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\xcb\x6e\xab\x30\x10\x5d\xdf\xfb\x15\x23\x7f\x00\xe4\xde\xec\x22\xe2\x2a\x4d\xb2\x68\x85\xda\x4a\xf4\x07\x1c\x32\x04\x4b\xc6\x76\xed\x61\x85\xf2\xef\xb5\x71\x52\xa0\x6a\x57\x30\xc7\xe7\xcc\x99\xd7\x30\xc0\x19\x1b\xa9\x11\x58\x6b\x3a\x64\x70\xbd\xfe\xfd\x33\x0c\x84\x9d\x55\x82\x02\x4a\xc6\x32\xc8\x02\x1a\x61\x90\x0d\x68\x43\x90\xa1\x16\x27\x85\xe7\xc0\x86\xa2\x5d\x43\xad\x84\xf7\x5b\x86\xce\x19\xc7\xf8\x7e\x57\x96\x15\x1c\x9e\xaa\xdd\x63\x79\x3c\x14\x79\xbb\xe6\x10\xb4\xa8\x23\x3f\xe4\x29\x2c\x0f\x61\xd6\x3b\x15\xe2\x22\xb7\x3c\x61\xa5\xf0\x04\x9e\x04\xf5\x7e\x03\x85\xb7\x42\xdf\xf3\x26\x63\xfc\x80\x2c\xc4\x54\x8d\x14\xf8\xbf\x5a\x45\x7b\xdf\xd7\x35\x7a\x3f\x1a\x28\x8f\x11\x1a\xcb\x98\x1c\xd9\xe8\x36\x53\x46\xd3\x98\x9e\x7f\xf3\xb6\x52\x5f\x80\x64\x87\x1b\xb8\x2b\xde\x02\xf4\x1e\x90\x45\xa1\xf1\x51\xf7\xdd\xde\x68\x12\x35\xc5\x7c\x60\x5b\x13\x46\x18\xc0\x13\x3a\x0f\x26\x54\x2e\x94\x9a\x14\x85\x80\xd6\x61\xb3\x65\xb9\x47\x3a\xa6\xd9\x3d\xdc\x66\xb8\x4d\xdd\xcd\x46\xba\x9a\x7a\xf9\xb7\xec\x62\xc9\x83\xb3\xf4\xf1\x7f\xd1\xbb\xfe\x42\x46\x59\x91\x0b\xfe\x53\x1d\xb5\xd1\x8d\xbc\x30\x9e\xbe\xbd\x13\x24\x8d\xfe\x8d\x9c\x96\xc2\xf8\x73\xf5\xfa\x72\xdb\xd0\x8c\x3a\xbf\x96\x93\x21\x32\x5d\x3a\x98\x69\xe5\x9f\x01\x00\x00\xff\xff\xd2\x33\xc4\xe8\x65\x02\x00\x00")

func templatesHomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHomeHtml,
		"templates/home.html",
	)
}

func templatesHomeHtml() (*asset, error) {
	bytes, err := templatesHomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/home.html", size: 613, mode: os.FileMode(436), modTime: time.Unix(1444930010, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesTopHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x91\xc1\xee\xda\x30\x0c\xc6\xcf\xf4\x29\xb2\xec\x90\xc3\x28\x6d\x05\x1a\xa3\x4b\xba\xc3\x18\xd7\x21\xb1\x69\xda\x31\x4b\x5d\x1a\x91\x26\x55\x62\x60\x15\xe2\xdd\x97\x52\x86\xf4\x3f\x44\x76\x92\x9f\x3f\x5b\x9f\x6f\x37\x52\x43\xa3\x2d\x10\x8a\xae\xa7\xe4\x7e\x4f\xf8\xbb\xed\xf7\xaf\x3f\x7e\xef\xbf\x91\x16\x3b\x53\x25\xfc\x7f\x00\x59\x57\xc9\x8c\x77\x80\x92\xa8\x56\xfa\x00\x28\xe8\x19\x9b\xf4\x13\x7d\xbd\x5b\xd9\x81\xa0\x17\x0d\xd7\xde\x79\xa4\x44\x39\x8b\x60\x23\x77\xd5\x35\xb6\xa2\x86\x8b\x56\x90\x3e\x2e\x73\xa2\xad\x46\x2d\x4d\x1a\x94\x34\x20\x8a\x87\x8a\xd1\xf6\x44\x5a\x0f\x8d\x60\x2d\x62\x1f\xca\x2c\x6b\xa2\x46\x58\x1c\x9d\x3b\x1a\x90\xbd\x0e\x0b\xe5\xba\x4c\x85\xf0\xa5\x91\x9d\x36\x83\xf8\xe5\xfc\xe9\xc3\x41\xda\x50\xae\xf2\x7c\xbe\xce\x73\x46\x3c\x18\xc1\x02\x0e\x06\x42\x0b\x80\x8c\xe0\xd0\x83\x60\x08\x7f\x71\xac\x64\x63\xa7\xc7\x77\x4c\x66\x7f\x5c\x3d\x90\x5b\x4c\x66\x63\xa7\x74\x52\x2d\x09\x1b\x75\xc9\xa8\xcb\xe6\x24\xc4\x90\x06\xf0\xba\xf9\x1c\xc1\x7b\x3c\x8b\x70\x56\x0a\x42\x98\x2a\x95\x33\xce\x97\xe4\xfd\x72\xbb\xd9\xac\xf3\x17\x03\xde\x3b\xff\x96\xd8\xed\x56\xc5\xf2\xe3\x93\xe0\xd9\x73\x0a\x9e\x4d\xf6\xf2\x71\x98\xd1\xec\xa2\xe2\x72\xf2\x81\x66\xb4\x3a\xfc\xdc\xf3\x4c\x56\x91\x2a\xaa\x24\xee\x0c\x6c\x1d\x57\xf5\x2f\x00\x00\xff\xff\xd1\x1e\xe0\xec\xbd\x01\x00\x00")

func templatesTopHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTopHtml,
		"templates/top.html",
	)
}

func templatesTopHtml() (*asset, error) {
	bytes, err := templatesTopHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/top.html", size: 445, mode: os.FileMode(436), modTime: time.Unix(1442354384, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/bottom.html": templatesBottomHtml,
	"templates/config.html": templatesConfigHtml,
	"templates/home.html": templatesHomeHtml,
	"templates/top.html": templatesTopHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"bottom.html": &bintree{templatesBottomHtml, map[string]*bintree{
		}},
		"config.html": &bintree{templatesConfigHtml, map[string]*bintree{
		}},
		"home.html": &bintree{templatesHomeHtml, map[string]*bintree{
		}},
		"top.html": &bintree{templatesTopHtml, map[string]*bintree{
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

