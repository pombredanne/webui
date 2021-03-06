package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_base_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x54\xdd\x4e\xdb\x4c\x10\xbd\xcf\x53\x0c\xbe\x01\x24\x6c\x2b\x1f\xa0\x8f\x22\x3b\x12\xa5\xb4\x4d\x45\xa1\xe5\xa7\x3f\x42\x5c\xac\xbd\xe3\x78\xc3\x7a\xd7\x78\xd7\x09\x11\xca\xbb\x77\x36\x76\xea\x04\x01\xe2\xaa\xb9\x81\x59\xcf\x9e\x99\x39\xe7\xcc\x3e\x3e\x02\xc7\x4c\x28\x04\x2f\x61\x06\x3d\x98\xcf\xa3\x8d\x0f\xe7\xc7\x57\xbf\xbf\x9d\x40\x6e\x0b\x39\xe8\x45\xcb\x3f\xc8\x38\x48\xa6\x46\xb1\x87\xca\x1b\xf4\x80\x7e\x51\x81\x96\x41\x9a\xb3\xca\xa0\x8d\xbd\xeb\xab\x8f\xfe\xc1\xda\xa7\xdc\xda\xd2\xc7\xfb\x5a\x4c\x62\xef\x97\x7f\x7d\xe4\x1f\xeb\xa2\x64\x56\x24\x92\x6a\xa5\x5a\x59\x54\x74\x6f\x78\x12\x23\x1f\xe1\xda\x4d\xc5\x0a\x8c\xbd\x89\xc0\x69\xa9\x2b\xbb\x92\x3c\x15\xdc\xe6\x31\xc7\x89\x48\xd1\x5f\x04\x3b\x20\x94\xb0\x82\x49\xdf\xa4\x4c\x62\xdc\x27\xa0\x06\xc9\x0a\x2b\x71\xf0\xa3\x96\x8a\x27\x30\xc5\x04\x6a\x83\x15\x65\x5b\xac\x32\x96\x62\x14\x36\x09\x6d\xf6\x86\xef\xc3\x29\xb3\x68\x2c\x55\x2b\x4a\x21\x91\x03\x53\x1c\x0a\x82\xcf\x04\x05\xc7\x97\x97\xe0\xfb\x6d\x97\x52\xa8\x3b\xa8\x50\xc6\x9e\xb1\x33\x89\x26\x47\xa4\x36\xf3\x0a\xb3\xd8\x73\x63\x9b\xc3\x30\x2c\xd8\x43\xca\x55\x90\x68\x6d\x8d\xad\x58\xe9\x02\x82\x0e\xff\x1e\x84\xbb\xc1\x6e\xb0\x17\xa6\xc6\x74\x67\x01\xd5\x0b\xe8\xc4\x5b\xed\xeb\xbc\xb4\x42\x2b\x26\xc1\xe6\x58\xe0\x3f\xe8\xc2\x5f\x14\x7a\xb6\x97\xcf\x57\x5f\x4f\xf7\xc1\xe4\xa2\x58\xd0\x73\x81\xa6\xd4\x8a\x07\x63\x03\x99\xae\x60\x78\x72\x00\xa6\x2e\x9d\x6a\xa0\xb3\x36\x19\x25\x81\x29\x6b\x1a\x3e\x91\x0b\x06\xf7\x35\x56\x02\x4d\x37\x8a\x83\xfe\x79\x74\x71\x36\x3c\xfb\x74\xb8\x0a\xca\x35\x1a\xb5\x69\x61\xaa\xab\x3b\x10\x19\xcc\x74\x0d\xce\x17\x8e\x0a\x28\xd9\x08\x29\x62\x90\x91\x5c\x34\xeb\x1a\xdc\x0d\x65\x4b\x4b\x1d\xc1\xbb\xdb\xf6\xd4\xa4\x95\x28\x2d\x98\x2a\xed\xf8\xd1\xc6\x04\x2d\x47\x8e\x16\xe7\xf7\x7d\x9a\x6e\x42\xb4\xfc\x1f\xfc\xd7\xc5\x0b\x32\xc6\xc4\x45\x14\x36\x30\x6f\xc7\xac\x9a\x71\xc2\x7e\xb0\x47\x88\x6d\xf4\x12\xde\xc6\x0d\x2a\x2e\xb2\x5b\x37\x4a\xaf\x17\x85\x6e\xf3\x68\x01\x13\xcd\x67\x4b\x19\xb8\x98\x40\x2a\x99\x31\xb1\xe7\x16\x83\xd1\x06\x57\x4b\x8d\x9e\x26\x38\x8a\x7c\x87\xb1\x48\x81\x95\x5f\xc4\x5a\xa7\x84\xd4\x43\xde\x7f\x6d\x51\xe8\x6b\x14\xb2\xee\x7a\x14\x52\x85\x95\x82\xf4\x90\x58\x2c\x4a\x49\xeb\x03\x5e\xbb\xab\x1e\x04\xf3\x79\x6f\x35\xbb\x13\x7a\xfc\x9d\xe4\x9f\xc1\x96\xc2\x14\x8d\x61\xf4\xaf\xf3\xce\xfb\xa5\xfb\x36\x0d\x7c\x61\x13\x76\xd9\x30\x5b\xca\x7a\x24\x94\xd9\xee\xc4\x7d\x8e\x72\x36\x66\x0f\xc1\x48\xeb\x91\x44\x56\x0a\xb3\xe0\xdd\x9d\x85\x52\x24\x26\x1c\x3b\xbf\xcd\x48\x80\x7e\x9f\x14\x68\xa2\x17\x05\xa0\x06\x87\x2a\x95\x35\x47\x60\x52\x76\xaf\x41\xdb\x08\x6c\x25\x28\xf5\x74\x7b\x07\xb4\x23\xa9\x49\x14\xa4\xd9\x44\xf0\x9a\x76\xd4\x79\x91\xac\x6e\x40\x21\x72\xba\xf6\x6a\xdb\x6f\xdd\xce\xf1\xd3\x27\x62\xbd\x71\xf2\x49\x63\x90\x68\x61\xd8\x01\x09\x42\x2e\xa2\xf7\xfc\x4f\x00\x00\x00\xff\xff\x78\x50\x7c\xbc\xe3\x05\x00\x00")

func templates_base_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_base_html,
		"templates/base.html",
	)
}

func templates_base_html() (*asset, error) {
	bytes, err := templates_base_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/base.html", size: 1507, mode: os.FileMode(420), modTime: time.Unix(1428176593, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_vuln_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x4d\x73\xd3\x30\x10\xbd\xfb\x57\x68\xcc\x05\x0e\xb6\x86\xf4\xc2\x74\x5c\x5d\x4a\x60\xb8\x42\x81\xb3\xb0\xd6\xb6\x40\x96\x32\x2b\xb9\x69\xa7\xd3\xff\x8e\xa4\xf8\x33\x38\x9d\x26\x33\xf4\xd0\x24\xfb\xbc\xbb\xef\xbd\xdd\x75\xf2\xf4\x44\x04\x54\x52\x03\x49\x4b\xa3\x1d\x68\x97\x92\xe7\xe7\x24\xc4\x65\x45\xf2\x2d\x62\xf8\x59\x34\x57\xa4\x54\xdc\xda\x9b\xd4\xc1\x83\xcb\x04\xd7\x35\x60\xca\x3c\x6c\x90\x98\xb2\x43\x10\xd7\xc4\xe7\xf4\x09\x05\x6d\xae\x58\xa8\x01\xca\x42\xac\x97\x24\x85\x90\xf7\x43\x11\x34\xfb\x94\x25\xc4\xff\xcd\xa3\xa5\x51\x59\x2b\xb2\xf7\x9b\x1e\x8b\x38\x27\x0d\x42\x75\x93\x52\x0b\xf7\x80\xd2\x3d\xd2\xd0\xe6\x47\xa7\x74\xfe\xad\x8f\xf8\x06\x29\x2b\xec\x8e\xeb\xa1\x92\xe2\xbf\x40\x91\xf8\x3f\xf3\x8f\x0f\xa9\x6b\x79\x6b\xd5\x0a\x1a\x8a\xb1\x82\xf2\x89\x48\x6f\x48\x7c\xf4\x8e\xd7\x36\xa8\x9a\x61\x18\x1c\x39\x01\x4f\x1a\x1c\xaf\x23\xfd\x97\x19\xef\x50\xb6\x1c\x1f\x0f\xdc\x8e\xe8\xcc\x7b\x82\x16\x47\x2c\x66\x91\x82\x7a\x67\x59\xd2\x7f\xf8\x09\x6e\x26\xa9\x77\xd2\xa9\x30\x16\x52\xd8\x96\x2b\xc5\xde\x8c\xc8\x17\x71\xe8\x17\xc3\x7e\x8a\x1b\x36\xae\x42\xc4\x3f\x82\x2d\x51\xee\x9c\x34\x3a\xee\xc5\x6b\x67\xfa\xc1\xcb\xf5\x2b\x31\x4b\x8f\x2b\xb2\x24\xf9\xfa\x6a\x24\x6e\xe1\xef\xce\x3a\x59\x1d\x7c\xf2\x86\xfd\x11\x66\xaf\x57\x79\x2e\xdb\x4c\x3e\x2d\xa4\x7d\x92\x0f\x97\x48\x0a\x69\x75\x27\xfd\x45\x94\x30\xf8\xf9\x76\xf4\xd3\xa3\xf9\xb6\xaa\x0c\xba\x60\x77\x2b\x75\xe7\xc0\xbe\x9b\x19\xfc\x1f\x3d\x08\xbd\x3f\x0f\xcc\x5e\x30\xe1\x5c\xc5\x5f\xa1\x02\x04\x5f\xd4\xae\xf0\x9f\x3b\x7a\xbb\x8f\xb7\x7f\x74\x1d\x7d\xf4\xac\xae\xc3\x01\x35\xce\xed\xae\x29\x2d\xf7\x90\xb7\xd2\x21\xe4\x06\x6b\x2a\xb8\xe3\x34\xbe\xc4\x64\x98\xb7\x1d\x2e\x2c\x6f\x5c\xab\x52\x76\xfb\x73\x9b\x8d\x87\xc4\x19\x39\xe5\xc3\x89\xb5\x98\xd4\xae\x68\x59\x82\x17\x49\x0a\xd4\xbe\xa3\x1a\xdf\x45\xc3\x6d\x9e\xc1\xf5\x9f\xaf\x7f\x03\x00\x00\xff\xff\x92\x79\x82\x29\xd6\x05\x00\x00")

func templates_vuln_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_vuln_html,
		"templates/vuln.html",
	)
}

func templates_vuln_html() (*asset, error) {
	bytes, err := templates_vuln_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/vuln.html", size: 1494, mode: os.FileMode(420), modTime: time.Unix(1428176986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_vulns_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x52\xc1\x4e\xc3\x30\x0c\xbd\xf7\x2b\x4c\xee\x59\xc5\x76\x41\xa8\xcb\x8d\x03\x67\x10\xf7\xac\xf1\xd6\x4a\x21\x45\x49\x36\x40\xd3\xfe\x9d\x38\x4d\xd6\x74\x0c\x21\x24\x76\x58\x13\xbf\xf8\xf9\x3d\xdb\xd5\xf1\x08\x0a\xb7\xbd\x41\x60\xed\x60\x3c\x1a\xcf\xe0\x74\xaa\x28\xde\x6f\x61\xf1\x60\x2d\x5d\x9b\x6e\x05\xad\x96\xce\xad\x99\xc7\x0f\xcf\x95\x34\x3b\xb4\x4c\x04\x78\xb0\x30\xb4\x7b\x8b\xea\x1e\x42\x4e\x4a\x68\xea\x6e\x25\x88\x03\xb5\xc3\x92\xef\xb9\xf7\x3a\x06\x1a\xd5\x1f\x32\xa5\x1d\xde\x99\xa8\x20\xfc\xca\x68\x3b\x68\xfe\xaa\xf8\xed\x32\x61\x11\xef\x96\x82\xaa\x64\x9a\x50\x67\x59\xa0\x1b\x0b\x75\x22\xaa\x03\x93\xa8\xd2\x87\x84\x18\x95\x75\x58\x12\x0f\x8b\x97\xbd\x36\xee\x4f\x52\xee\x4a\x25\x12\x3a\x8b\xdb\x35\xab\x0f\x81\xa7\x26\x51\x8f\x54\x81\x89\xf3\x11\x38\xcc\xb5\xca\x22\xfd\x86\xf3\xc6\xbd\x49\x93\x4b\x68\xb9\x41\x0d\xf1\x9f\x87\x2c\x87\x07\xb4\xbd\xff\x84\xc5\x53\x3e\x65\xee\x22\xd0\xd4\x44\x21\x38\x9f\x13\xe7\x5e\xcb\x1d\xf9\xbb\x82\xa6\x0e\xfc\xf0\x60\xb2\xe6\xe5\x2e\x3a\x8b\xb5\x67\x72\x47\x29\x93\x04\x32\x47\x3c\x94\x7f\xa5\xe0\xd8\xfd\x5f\x81\x34\xae\x3c\xb6\x28\xe6\x72\x38\xe7\xb7\x17\xe0\x34\xa3\x6f\x7e\xfe\xbb\xcf\x91\x34\x4a\xcc\x96\x8b\xdb\xb8\x83\xd3\xc6\xcd\x77\x6f\x3c\x7e\x05\x00\x00\xff\xff\x4e\xa2\x80\x2b\x76\x03\x00\x00")

func templates_vulns_html_bytes() ([]byte, error) {
	return bindata_read(
		_templates_vulns_html,
		"templates/vulns.html",
	)
}

func templates_vulns_html() (*asset, error) {
	bytes, err := templates_vulns_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/vulns.html", size: 886, mode: os.FileMode(420), modTime: time.Unix(1428176902, 0)}
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
	"templates/base.html":  templates_base_html,
	"templates/vuln.html":  templates_vuln_html,
	"templates/vulns.html": templates_vulns_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"base.html":  &_bintree_t{templates_base_html, map[string]*_bintree_t{}},
		"vuln.html":  &_bintree_t{templates_vuln_html, map[string]*_bintree_t{}},
		"vulns.html": &_bintree_t{templates_vulns_html, map[string]*_bintree_t{}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
