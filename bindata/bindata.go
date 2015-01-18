package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

func static_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x92,
		0x31, 0x53, 0xc4, 0x20, 0x10, 0x85, 0xeb, 0xbb, 0x5f, 0xb1, 0xd2, 0x1f,
		0x24, 0x6a, 0xe1, 0x44, 0x72, 0x8d, 0x5a, 0x6b, 0x61, 0x63, 0x89, 0xb0,
		0x06, 0xce, 0x10, 0x22, 0x6c, 0xf4, 0xf2, 0xef, 0x25, 0xc4, 0xc2, 0x51,
		0xc7, 0x46, 0xab, 0xb7, 0x2c, 0xc3, 0xf7, 0xe6, 0xbd, 0x41, 0x9e, 0x5c,
		0xdf, 0x5e, 0xdd, 0x3f, 0xdc, 0xdd, 0x80, 0x25, 0xdf, 0xef, 0xb7, 0x72,
		0x95, 0x8d, 0xb4, 0xa8, 0x4c, 0xd6, 0x8d, 0x24, 0x47, 0x3d, 0xee, 0xa5,
		0x58, 0x75, 0xd9, 0x78, 0x24, 0x05, 0xda, 0xaa, 0x98, 0x90, 0x5a, 0x36,
		0xd1, 0xd3, 0xee, 0x82, 0x81, 0x28, 0x57, 0x49, 0x47, 0x37, 0x12, 0xa4,
		0xa8, 0x5b, 0x66, 0x89, 0xc6, 0x46, 0x08, 0x1d, 0x0c, 0xf2, 0xc3, 0xcb,
		0x84, 0x71, 0xe6, 0x3a, 0x78, 0xb1, 0x8e, 0xbb, 0x9a, 0xd7, 0x35, 0xaf,
		0xb8, 0x77, 0x03, 0x3f, 0x24, 0x96, 0xf9, 0xeb, 0xd3, 0x5f, 0x29, 0xd6,
		0x75, 0x76, 0xf1, 0xa5, 0x54, 0x48, 0x89, 0x82, 0x7e, 0x16, 0xcb, 0xb2,
		0x4c, 0x7f, 0xe4, 0xf8, 0x60, 0xa6, 0x1e, 0x93, 0xc0, 0xe3, 0x18, 0x22,
		0xb9, 0xa1, 0xfb, 0xca, 0x93, 0xe2, 0xa3, 0x13, 0xf9, 0x18, 0xcc, 0x5c,
		0x0c, 0x8c, 0x7b, 0x05, 0x67, 0x5a, 0xa6, 0xc3, 0x40, 0xca, 0x0d, 0x18,
		0x6b, 0x06, 0x89, 0xe6, 0x1e, 0x5b, 0x96, 0x83, 0xed, 0xde, 0x9c, 0x21,
		0xdb, 0xc0, 0x59, 0x5d, 0x8d, 0xc7, 0x4b, 0xb0, 0x98, 0x6d, 0xa9, 0x81,
		0xf3, 0xaa, 0x1c, 0xbd, 0x8a, 0x9d, 0x1b, 0x1a, 0xa8, 0x40, 0x4d, 0x14,
		0x16, 0xa7, 0x4c, 0xfb, 0x99, 0x7a, 0xfa, 0x1f, 0xd4, 0xcf, 0x65, 0x78,
		0xf5, 0xbd, 0x76, 0x29, 0xd6, 0x58, 0x39, 0x66, 0xf9, 0x02, 0xdb, 0xf7,
		0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x27, 0x8d, 0x1c, 0x1b, 0x02, 0x00,
		0x00,
	},
		"static/index.html",
	)
}

func static_main_js() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xec, 0x55,
		0x5d, 0x6b, 0xe3, 0x38, 0x14, 0x7d, 0xb6, 0x7f, 0x85, 0xd6, 0x94, 0xc6,
		0x66, 0xb3, 0xce, 0x57, 0x0b, 0x8b, 0x4b, 0x28, 0xdd, 0x2e, 0x74, 0x18,
		0x98, 0xb6, 0x90, 0xc0, 0x3c, 0x94, 0x3c, 0x28, 0xf6, 0x75, 0x62, 0xaa,
		0x48, 0x46, 0x92, 0x9b, 0xa4, 0x43, 0xfe, 0xfb, 0x5c, 0x49, 0xb6, 0x1b,
		0x4f, 0x29, 0xb4, 0xc3, 0x30, 0x4f, 0xf3, 0x12, 0x1c, 0xdd, 0x73, 0xee,
		0xe7, 0x91, 0xee, 0x13, 0x95, 0x24, 0x5d, 0x53, 0xa9, 0x47, 0x17, 0x7e,
		0xfb, 0x3d, 0xbe, 0xf0, 0xfd, 0x93, 0x30, 0xaf, 0x78, 0xaa, 0x0b, 0xc1,
		0xc3, 0x88, 0x7c, 0xf3, 0x3d, 0x63, 0xdc, 0x91, 0x29, 0xe1, 0xb0, 0x25,
		0xff, 0x53, 0x0d, 0x61, 0x84, 0x20, 0xef, 0x53, 0xb1, 0x5a, 0x5b, 0x8a,
		0x8a, 0x15, 0xe8, 0xbb, 0xd2, 0xe0, 0x55, 0x88, 0x70, 0x6f, 0xc5, 0xc4,
		0x92, 0xb2, 0xc4, 0x50, 0x3d, 0x4f, 0x17, 0x1b, 0x78, 0x16, 0x1c, 0xee,
		0xf2, 0x1c, 0x61, 0x09, 0xd9, 0xc5, 0x2b, 0xd0, 0xf3, 0xce, 0x61, 0x18,
		0x21, 0xf0, 0xe0, 0x7b, 0x87, 0x08, 0xdd, 0x9e, 0x18, 0xfb, 0xe7, 0xd9,
		0xdd, 0x6d, 0xd8, 0x1b, 0x64, 0xb0, 0xac, 0x56, 0x03, 0x17, 0x64, 0x90,
		0x51, 0x4d, 0x2f, 0x53, 0xca, 0xd8, 0x92, 0xa6, 0x8f, 0xd3, 0xcb, 0x5e,
		0x9f, 0xb4, 0x59, 0x1a, 0x93, 0xcd, 0xd4, 0x73, 0xf5, 0xd4, 0xb9, 0x1e,
		0x65, 0x38, 0xd3, 0x22, 0x7d, 0xbc, 0x36, 0xdf, 0x36, 0x43, 0x87, 0xab,
		0x33, 0xf4, 0x24, 0xf0, 0x0c, 0xe4, 0x5c, 0x24, 0xa4, 0x97, 0x0a, 0xae,
		0x69, 0xc1, 0x41, 0x8e, 0x7a, 0x7d, 0x6b, 0x7b, 0x16, 0x62, 0x33, 0xdf,
		0x97, 0x80, 0xb6, 0x5d, 0xcf, 0x9c, 0x1c, 0xfa, 0xae, 0x2a, 0xcd, 0xa0,
		0xe1, 0x6b, 0xd8, 0xa1, 0xaf, 0xde, 0xcd, 0x35, 0x29, 0x69, 0xa5, 0x40,
		0x1d, 0xe1, 0xf6, 0x57, 0xbb, 0x42, 0xb5, 0xb8, 0x63, 0x52, 0xc3, 0xba,
		0xa5, 0x5c, 0x28, 0xc0, 0xb8, 0x99, 0xe3, 0x99, 0x46, 0x34, 0x6c, 0x95,
		0x4a, 0x61, 0xea, 0x95, 0x0d, 0x09, 0x38, 0x5d, 0x32, 0xc8, 0x12, 0x92,
		0x53, 0xa6, 0xe0, 0x05, 0x28, 0x29, 0x5f, 0xc1, 0x0c, 0x18, 0xa4, 0x5a,
		0xb4, 0xe0, 0x65, 0xa5, 0x35, 0x8e, 0x24, 0x21, 0x0f, 0x4d, 0x44, 0x57,
		0x87, 0x8b, 0x56, 0xd7, 0xe7, 0xa5, 0xa2, 0xe2, 0x98, 0xc7, 0x79, 0xbf,
		0x93, 0xd5, 0x79, 0x93, 0x4c, 0x9f, 0xbc, 0x83, 0x3c, 0x19, 0x76, 0xd9,
		0x93, 0xe1, 0x5b, 0xf4, 0x4d, 0xc1, 0x2b, 0x0d, 0x3f, 0xd0, 0x47, 0x5d,
		0xf6, 0x68, 0xf3, 0x06, 0x19, 0x67, 0xdf, 0xeb, 0x42, 0xaf, 0xf0, 0xc4,
		0x61, 0x17, 0xce, 0xa0, 0x6c, 0x0f, 0x4c, 0x87, 0x26, 0x47, 0x6d, 0x04,
		0x59, 0xc0, 0x4b, 0x1f, 0x38, 0xdd, 0xa0, 0xb7, 0xa0, 0x1d, 0x57, 0xe0,
		0xa8, 0x46, 0x44, 0x09, 0x31, 0xbf, 0xf1, 0x4d, 0x7a, 0x6f, 0x2d, 0xce,
		0xd0, 0x44, 0x97, 0x40, 0xeb, 0xf0, 0x5a, 0x08, 0xa6, 0x8b, 0xb2, 0x1d,
		0xe5, 0x13, 0x65, 0x15, 0xcc, 0xaa, 0x3c, 0x2f, 0x76, 0x08, 0xe4, 0x9d,
		0x41, 0x2e, 0x8c, 0xb6, 0xf1, 0xc2, 0xd4, 0x9a, 0x1b, 0xff, 0x02, 0x6d,
		0x8e, 0x3f, 0xae, 0xcd, 0x2f, 0xb0, 0x11, 0x72, 0x4f, 0xb0, 0x83, 0x22,
		0xc5, 0x1b, 0x9c, 0x7d, 0x48, 0xa2, 0xff, 0xed, 0x35, 0xfc, 0x11, 0xe7,
		0x6f, 0x15, 0xe7, 0x55, 0x33, 0xa8, 0xd7, 0xe2, 0xb4, 0xd3, 0x68, 0xed,
		0x3f, 0x27, 0xd1, 0xe5, 0x6b, 0x85, 0xfa, 0x56, 0xa5, 0xbe, 0xd7, 0x3c,
		0xaa, 0x64, 0xab, 0x2a, 0xc9, 0xdc, 0x02, 0xb0, 0x1b, 0x80, 0xa1, 0x72,
		0xb7, 0x05, 0xcf, 0xc4, 0x36, 0xb6, 0xb1, 0x11, 0x63, 0x54, 0x2d, 0x41,
		0x57, 0x92, 0x93, 0x30, 0x64, 0x71, 0x29, 0x05, 0x4a, 0x59, 0x20, 0x70,
		0x3a, 0x25, 0xc1, 0x5a, 0xeb, 0x52, 0x25, 0x41, 0x44, 0x2e, 0x49, 0xb0,
		0x55, 0x2a, 0x19, 0x0c, 0x02, 0x92, 0x98, 0x4f, 0xf3, 0x15, 0x91, 0xbf,
		0x09, 0x8b, 0xd7, 0x42, 0x69, 0x53, 0x30, 0xfe, 0x09, 0xad, 0x03, 0x21,
		0x35, 0xf9, 0x6b, 0x4a, 0xfe, 0x1d, 0x46, 0xe4, 0xf4, 0x94, 0x1c, 0x9d,
		0x9c, 0x9d, 0x4d, 0x22, 0xeb, 0x29, 0x09, 0x2c, 0xd3, 0x9e, 0xa3, 0x37,
		0xeb, 0x28, 0x78, 0xbd, 0x28, 0xfe, 0xc9, 0x01, 0x7b, 0x87, 0xf9, 0x1d,
		0xb0, 0xa4, 0xad, 0xaa, 0x2f, 0xdd, 0x57, 0x58, 0xce, 0xf0, 0xae, 0xe1,
		0xc6, 0xa9, 0x8b, 0x33, 0xf7, 0x72, 0xab, 0x62, 0xc1, 0x45, 0x09, 0x1c,
		0x41, 0x6d, 0xf1, 0x75, 0xdd, 0xd6, 0xb6, 0x01, 0xa5, 0xe8, 0x0a, 0x3a,
		0x66, 0x78, 0xd2, 0x0e, 0x61, 0x5b, 0x63, 0x22, 0xa2, 0xd9, 0x6c, 0xad,
		0xb8, 0xa4, 0x52, 0x81, 0xb1, 0xc7, 0x76, 0x2b, 0x99, 0x16, 0x79, 0x45,
		0x4e, 0xc2, 0xe3, 0x87, 0xc5, 0x54, 0x34, 0xac, 0xf9, 0xf5, 0xca, 0x8a,
		0x9d, 0x0e, 0x1e, 0x86, 0x8b, 0x98, 0x66, 0xd9, 0xbd, 0x28, 0xb8, 0x0e,
		0x1f, 0x2c, 0x67, 0xae, 0xfa, 0x9d, 0x57, 0x69, 0xd1, 0x27, 0x5a, 0x56,
		0xe0, 0x3c, 0x1f, 0xda, 0x07, 0x63, 0xfc, 0x0e, 0x0f, 0x5d, 0xe9, 0x1c,
		0x3b, 0x32, 0xdb, 0xf7, 0xc2, 0x47, 0x15, 0x7c, 0x0f, 0x00, 0x00, 0xff,
		0xff, 0x46, 0xc4, 0x57, 0x98, 0x15, 0x08, 0x00, 0x00,
	},
		"static/main.js",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"static/index.html": static_index_html,
	"static/main.js": static_main_js,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
		"main.js": &_bintree_t{static_main_js, map[string]*_bintree_t{
		}},
	}},
}}