// Code generated by go-bindata.
// sources:
// data/stub_flasher.json
// DO NOT EDIT!

package esp32

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

var _dataStub_flasherJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x79\xcb\xae\x2e\xb9\x6d\xf5\xfc\x7f\x8a\x83\x3d\xee\x01\x49\x91\x14\xe5\x19\x75\x7b\x0d\xa3\xf1\xdb\x43\x3b\x81\xdd\x19\x04\x41\xde\x3d\x20\x55\x55\x5f\xed\xd3\x8e\x9d\x73\x80\xaf\x2e\x52\xe9\x42\x2e\x2e\x2e\x6a\xff\xd7\xd7\xff\xff\xb7\x3f\xfd\xf9\x8f\x7f\xff\xed\xd7\xbf\xfd\xf6\xf5\x87\x1f\x08\x95\x4b\x41\x15\xfa\xe5\x47\xb6\x7c\xfd\xe1\xc7\x17\x18\xc0\xe6\xb2\x21\x6e\xae\x7f\x7b\x96\x0d\xe3\xf5\xcc\x65\xe3\x38\xd7\xbd\xeb\x8e\x77\x78\xbd\x2f\xca\x08\xb0\x70\xb7\xbd\x7b\xfe\x12\x6e\xd9\x7b\xe3\xd6\xbd\xbb\xf7\x35\x80\x00\x80\x0c\x68\xfb\x5c\x83\x46\xdf\x03\x14\xcb\x19\xbb\x41\x83\xca\xee\xad\xb5\xda\x37\x88\x83\x81\x82\x81\x41\xe5\x66\xd0\x0b\xd1\x00\x28\xd0\xda\x00\xab\xde\x80\x5a\xb6\xc1\x80\x66\x0a\x12\x63\xb8\xc1\x60\x5c\x7d\x6f\xec\xd1\x2f\xe6\x2b\x06\x8c\x9d\x9d\x1b\x8c\x42\xec\xf7\x4c\xad\x36\x5b\x1a\x5f\xc4\xe8\xad\x36\x07\x5a\xa4\x03\xa0\x19\x4c\x73\x6b\xd6\x60\x16\x5c\xb4\xf7\x19\xa9\x19\x94\x33\xa7\x0d\x35\x37\x98\x8c\x73\xef\xed\xb5\x73\xdb\x96\xf7\xd7\x1e\x1b\xc4\xe3\x63\xb7\x30\xd2\xb1\x0f\x01\x43\x67\x23\x87\x42\xa2\x85\xa1\x16\x68\xf7\x9b\x22\xea\x0d\x74\x89\x0d\x80\x7c\xc3\xc5\x89\x6c\x6c\x67\xc5\x7a\x86\x72\xf0\xc1\x48\xd2\x36\x80\xa8\x3a\x70\xa1\x70\x53\x01\xa7\x01\x63\x68\x81\x01\xc6\x8a\x78\x4f\xee\xc0\xb0\x59\x14\x78\x2f\xc6\xb5\xf6\x36\x8d\x65\x3a\x38\x6c\x71\x69\x0c\xb0\x86\x3a\x43\x29\xe3\x1a\xa7\xd4\xc2\x6b\xb1\x42\xf9\x8c\x12\xf3\xaa\xb4\x0e\x30\x07\xfb\x7b\x56\xd1\x55\xf6\x02\xdd\xb0\x77\x19\x62\xa6\xb9\xe3\xa1\xc6\x1a\xdf\x96\x51\x8d\xe6\x65\x13\xbb\xf0\xc4\xe7\x4a\x1c\xa6\xe5\xbc\x06\x8e\x94\xe1\xbc\x78\x70\xb6\xd3\x90\x7b\xc3\x7e\x8c\x09\x00\x75\x4c\x07\xe1\x58\x79\x05\x60\x00\x53\x8b\x2b\x8f\xc1\x71\x6d\x57\x7b\x87\xd3\x3e\xf8\x5c\xf7\xd5\x1f\xc7\xec\xd1\x5e\xd4\x10\x90\x49\x91\x8a\x4f\xc0\xcb\x43\xb8\x79\xef\x95\xc1\x50\x06\x89\x68\x38\xbb\x64\x5b\x19\x45\x44\x85\x71\xba\x44\x0c\xf0\x04\x2f\x83\x45\xb4\x3b\x8e\x8a\xad\xed\xad\xb8\xca\xde\x82\x4d\xf7\x56\xd7\x5a\x1b\xa8\x9d\x27\x5c\xb8\xc3\xec\xb5\xe3\x82\xbd\xcd\xad\x8e\x09\x81\x68\x89\x9e\x30\xc2\x25\x60\xe7\x8b\xf5\xac\xc2\x70\x8e\x17\x12\xad\xc5\x8a\xa1\xe1\xf4\xe7\x6d\x03\x6b\x38\xdb\xde\x46\x8e\x80\xe7\xad\x5b\x76\x77\x82\xf3\x6c\x0d\x9a\x93\x03\xc5\x88\xed\x1e\x3d\xa0\xd3\x27\xb0\xd3\x40\x31\xc3\x59\x9f\x79\x0d\x04\x86\xf5\x68\xa9\x16\x76\xc2\x65\xf1\xbb\x5d\x38\x70\x0a\xcb\xa1\xe4\x9b\xc6\x7a\xf8\x23\xde\x10\xb4\xc1\x18\xef\x09\x73\xad\x44\x50\x61\x0b\x28\xf4\xf4\x30\x0e\xd9\x9b\xc1\x4a\xc4\x73\x6f\xa8\x04\x76\x70\x6a\x80\x3e\xa1\xae\xb4\xae\x28\x3b\xa2\x61\xab\x61\x2b\x66\xb8\x30\xe9\x44\x48\xb4\xc4\x34\xfa\x54\x47\x60\x9a\x8c\x60\xf1\xbe\x88\x8a\x01\xd5\x6b\x6d\x84\x68\x2d\xde\x63\x03\xb3\x01\x15\xaa\x19\xd4\x58\xb3\xb5\xec\x73\x6c\x53\x89\x24\x2c\xdd\x69\xef\x6c\xdd\x9f\x08\xc8\x51\x96\xd1\x34\x04\xab\x85\xdb\x62\xdd\xe5\xb6\x7d\x35\x90\x5a\xbb\xed\xfa\x9a\x53\xbb\x81\x5e\x33\x7a\xb5\x1c\x51\x2b\xba\xed\x1d\x48\x0b\xaf\x1b\x7a\xbf\x2d\x7d\x56\x12\xdf\x24\x1e\x26\x50\xda\x8d\x55\x22\x64\x72\xfd\x1c\x3b\x85\xf6\x9a\xa3\x12\x4d\x42\xf8\x69\x96\x18\x17\xef\x71\x5f\x2b\xa4\x59\xf7\xfd\xdc\x40\x0c\xdb\xbc\x7b\x55\x03\x8d\x76\x04\xa9\xbd\x02\x07\xfe\x20\xe3\xd8\xb0\x04\x67\xb4\x68\x0f\x1b\x04\x1a\x4e\xcf\x83\x89\x9c\x71\x91\x94\x91\x16\x44\x5c\x25\x46\x20\xd0\xbe\xf6\xa6\x01\x6e\xd8\x9e\xd5\xe4\x3e\x80\xd0\x68\xef\x3e\x6f\x94\x36\xd2\x60\x24\x1a\xa8\x66\x68\xe3\xc1\x65\xee\xb4\xc4\xa8\x8d\x06\xf4\x02\x91\x06\x0e\x4a\x29\x31\x07\x41\x3f\x52\x62\x9f\xc9\x6a\x65\x88\x3c\xfc\x23\x0f\xff\x08\x4d\x90\xe0\x9f\x08\xf7\xf8\x27\xd7\xb5\x5e\x7c\xc2\xd0\x4a\xfc\x06\x4f\x13\x2a\xee\x19\x18\x98\x88\x80\xae\x2a\x0d\xb4\x61\x0d\x16\x1a\xa0\x4a\xe1\x46\x19\x64\x5c\x7b\x03\x32\x0d\xdc\x78\xfa\x53\x1f\x8b\xb2\x62\xce\xae\x07\xe3\x20\xac\x65\x4c\x90\xd8\x75\x78\x57\xc5\x6a\x60\x77\x38\x68\x9f\x20\x81\x09\xc3\xb5\x1f\xbf\x7d\xfa\xc6\xb8\x7a\xbf\x17\x27\x12\x57\x2d\xb5\x28\x6c\xc1\xa5\xb1\x2a\xee\xe2\x8e\x2a\x25\xb2\x4c\xc4\x8c\x42\xe6\x00\x49\x0e\x31\xa2\x8a\x8e\x2e\x22\x06\x62\x60\x3a\x40\x6a\xb1\x0e\xa0\xbb\x06\xc3\x9e\x1e\x56\x8b\xd2\xd5\xbb\xd4\x62\x63\x38\xce\x9e\x9e\xaa\xe8\xee\x86\xf2\x70\x86\x27\x4b\xe0\x80\x7e\x5a\xa6\xdd\x2d\x91\x81\xee\x38\x95\x51\xae\x7c\x20\x83\xaf\x7c\x20\x43\x9e\x7c\xd0\x2f\xfb\x17\x1d\x69\x77\x8e\xd5\x0e\x67\x21\x07\x25\xae\x5d\xd3\xbe\xc9\xe9\xca\x04\x82\xa3\x25\x7f\x86\x5d\xc4\xb0\xec\x97\xbd\x6b\x60\x28\x7c\x66\x58\x1e\x0c\x29\xd4\xa2\xa5\xc2\xa8\x04\x08\x3d\xd0\x17\xec\x01\x2a\xb5\x03\xd8\x50\xcf\xb9\xca\x78\x72\x6c\xc4\xdf\x65\xf9\x98\xe1\xe1\x42\x1f\x95\x9f\xb6\x7b\x9e\xf2\xb1\x47\x78\xa2\x1f\xcc\x80\xbb\xfd\xdc\x5a\xde\xd6\x0a\x8e\xb9\x22\xcf\x89\x6a\xe6\xde\xcc\xa3\xad\xa3\x8f\xbd\x1d\xe7\x4c\xdd\x84\xa7\x3f\xf9\xcb\xee\xed\x3d\x92\x3f\x1e\x91\x09\xe5\xc6\x79\x68\xa5\xa3\xe7\x4e\xce\xb4\xeb\x79\x5f\x19\xb5\xa8\x62\x7a\x8a\x0c\xb7\xef\x4d\x0d\x90\xf2\x2e\x12\xc6\x9d\x2d\x42\xc9\x50\x60\x03\x1e\xf5\xd6\x0c\x88\x30\x78\x99\xf6\xb6\xd4\x78\x81\xbe\x4f\x66\x32\x0c\x5e\x4c\x34\x47\x42\x8a\x8c\x13\xfc\xfd\x64\xb2\xb3\xba\x58\x51\xe8\xcb\xcc\xbf\xd0\xc9\x21\x6c\x3a\x10\x91\x04\x7b\x60\x18\x3d\x3d\xbf\x00\x28\x36\xbf\x8d\x20\x5a\x89\x11\xa1\xab\x19\xd0\x50\x0a\x8e\x48\x85\x49\x60\xe6\x40\x14\x6f\x1c\x72\x31\x70\xc6\x4d\x6f\x60\xf8\x1b\x33\x9b\x23\xa9\x9a\x8b\x77\x43\x74\x03\x24\xe9\x25\x34\x51\xb4\x86\x55\x53\x05\x48\x60\x81\xb3\xaf\x3b\xf3\x30\xc2\xbb\xbf\x4a\x4f\x0d\x85\xf4\xbb\xf1\x25\xb0\x1d\xdf\x74\x27\x9f\x56\xf0\xfd\x1d\x49\x78\xd4\x32\xff\xbd\xbf\x73\x09\x0d\x4c\x81\x41\x8c\xd9\x60\x3e\x5f\x85\xa1\x81\x64\xd5\xfc\x2e\x95\xd1\xd9\xbf\x29\x34\x80\x25\x1b\x77\xe4\xa8\x8f\xd6\x93\x50\x99\x9e\x7d\x20\x55\x71\x8c\xfb\xcd\xbe\xd7\x08\x2e\x90\x96\x0d\x9b\xb6\x0b\xf6\xd8\x4f\x8f\x25\x30\xef\x36\xa3\x41\xdb\x8f\x0e\xef\x61\x6f\x32\x8d\x08\xfc\xa6\x6f\x1b\x03\x6f\x36\x46\x6e\xb7\x5f\x9f\xfa\xe1\x02\x20\x46\x3a\x8e\xd8\x87\x51\x8e\x7e\x3b\x7a\x4d\x47\x6c\x9c\x61\x18\xf5\xb8\x16\x2b\x94\x57\x36\x39\x7a\xce\xe8\xf0\xb2\x66\x2e\x0a\x6c\x11\xc6\xbe\x21\x32\x4b\xa0\xd6\x12\xc3\xfa\x89\x8f\xd0\xa1\x23\x99\x14\xc2\x9f\x30\xc6\x80\x01\x83\x08\x0d\xfb\xd8\xeb\x61\xa9\xf0\x75\x7f\x73\x6e\xda\xe0\xbc\x5b\xf7\xbb\x8d\x11\x7a\x13\x3b\xee\x35\xd0\x74\xaf\x8e\x4b\xf6\x5e\xa1\xb2\x20\x57\x81\xeb\x89\x41\x8c\x2c\x84\x9e\x75\x83\xe1\xaa\x9f\x4c\xd6\x89\xb2\xb7\xa7\x66\xf1\xac\x17\x2c\x59\xe1\xf4\x48\x1f\xf1\xc1\xc4\x12\x78\x3c\xef\x40\x4e\xd1\xdf\xa2\xff\x7c\xf2\x0b\xa9\x52\xf0\xd3\x1c\xef\xc8\x3a\x75\x46\x9f\x10\x6a\x96\x0c\x49\xbe\xc7\x9d\x73\x7b\xd9\x93\xc9\x83\xac\x33\xdb\x5d\x6c\x39\xc1\xb9\x36\xdf\x5c\xbc\x10\x4f\xa0\x9a\x7d\x58\x43\x91\xa7\x07\x1e\x9e\xb7\x09\x5e\x23\x53\xcc\x60\xd0\xd9\xa5\x36\x47\xf7\x06\x98\xcf\xc3\x70\xd3\x4b\x47\xa2\x54\x74\xb0\x6c\x9b\x79\x1f\x6a\x8e\x73\x1c\x35\x62\x06\xc0\xce\xc1\xf1\x3c\x42\x80\x94\x11\x99\xe3\xc2\x26\x08\x5d\x6b\x5d\xcf\xfc\x14\xeb\xd4\x0f\xe3\x9d\x3a\x72\xd1\xf6\xb2\x11\xe8\xe2\x3f\x7a\xf1\x20\x58\x07\xc2\xc4\xd7\xc5\x7f\x68\x87\x9f\x76\xea\xe6\xd0\x9e\x7b\x4c\xf3\xdc\x65\xcf\xdf\x91\xbf\x59\x8c\xc2\xad\xa6\x8c\xd4\xef\xfb\x41\xda\xef\xfb\xa8\x25\x45\x6b\x7b\x29\x75\x9f\x26\xd4\xa0\x4d\x0c\x85\x8e\xcd\x00\xa1\xae\x36\xdd\x4e\x86\x1c\xd2\xaf\xcc\x3f\xdf\xb6\x82\xba\x6c\x9f\xbb\xe0\xe8\x06\xe5\xec\x31\x16\x99\x3c\x6e\xe5\xaa\xbb\xaf\xb8\x98\xf7\x9c\xc5\xc0\x0a\x14\x08\xb5\x98\x1d\x93\xd5\xaf\xb6\xe4\xc7\xbb\x35\xb2\xd3\xf7\x8c\x98\x5a\x60\x60\xd4\x07\xc5\x72\xc6\x7a\xd9\xf3\x2a\xda\x5e\x79\x23\xf4\x57\xb2\xef\x08\xf6\x8d\xef\x5d\x76\x4f\x8c\xe2\xde\xe5\x3d\x6b\x03\xfb\xac\x4d\xb4\x96\x98\x85\x0a\xeb\xe1\xe9\xf7\x1a\x48\x22\xc2\x2d\x99\xcf\x2a\xc4\xc4\x03\xe8\xe3\x5f\xfb\xe6\xcf\x80\xf2\x37\xff\xd2\xc1\x3f\x83\x00\xb2\x0d\x92\xd8\x4d\x7c\x2d\x19\x8b\x16\x0a\x9f\x0d\x94\x16\x31\xa8\x86\xe8\x2d\xb6\x86\x80\x4b\xae\x5f\x22\xb3\x6e\x05\xa5\x4b\x0b\xa8\x84\xf7\x45\x9b\xef\x69\x88\xa1\xdd\x93\x7d\x0c\x8c\x10\xb8\xba\xd2\x6a\xc9\x16\xa1\xc1\xa3\x22\x5b\x4f\x9d\x76\x50\x52\xc0\x0b\x65\x3d\x0f\x56\x23\x6f\x79\x37\x70\xee\xee\xa7\x57\x6f\xd0\xb8\xb7\x56\x68\x14\x02\xa6\xc1\x2b\x54\x6e\xe4\x5c\x85\xac\xa7\x1a\x03\x9e\xd1\x9f\x1a\xce\x19\x1c\x90\x19\x38\xb4\x71\x37\xb6\x3c\xd5\x60\x60\xa8\x6c\x13\x5a\x8b\x6a\x0e\x9a\xb5\x72\xea\x3c\x9e\x50\x38\xab\x84\x02\xa5\x39\xc0\x53\xd9\xb8\x01\xb3\xb7\xcb\x43\x9e\xab\x61\x8e\xf8\xdb\x81\x3d\x3e\xb5\x57\xc3\xc9\xef\x4a\xf3\xd6\xed\x81\x27\x51\x0d\x4d\x96\x36\x54\x59\x7d\x6f\x57\xe4\xc8\x14\xa1\x10\x4f\x06\xba\x3d\x78\xfc\xe3\xd8\x65\x6f\x92\xe5\xa1\x98\x1c\x8a\x5c\x56\x7a\x69\x7f\x23\x8a\xbc\x38\x63\x81\x66\x59\xaa\x95\x6f\x59\x27\x01\x79\x8d\x37\x81\x3c\xd9\x32\x18\x8b\xd0\xa2\xb2\xce\xdf\x07\x81\x44\x19\xe3\x1f\x5d\x73\xdd\xdd\x3b\xa0\xa8\x46\x9d\x73\x07\x12\xde\xbc\x01\xf7\xc1\x17\xc3\xad\x9b\x28\x2b\xa7\x83\xd8\x25\xeb\xe8\x1d\x0a\x5d\xb8\x12\x2f\x65\x02\x05\xf2\x02\x37\xc0\x1c\xd9\x3f\x48\x1c\x08\x28\x18\x3e\x74\x5a\x64\x2e\x89\x2c\x41\xa8\x7d\x6f\x4e\x85\x74\xd6\xc4\x59\x0d\xdd\xf7\x91\x36\x44\x95\xf7\x3e\x36\x52\x59\x10\x5f\x05\x5f\x1c\xd6\x78\xb8\x86\x62\xe4\x3c\x1c\xf8\x20\x85\xfa\x99\x3d\xf2\xbf\x44\x96\x74\x20\xb6\xf2\xdd\x1f\x2a\x0b\xef\x4c\x23\xa9\xbc\x4a\x70\x62\xec\xaa\x2c\x89\x0c\x63\x28\x07\x05\xfb\xbc\xfd\x19\x09\x9f\xc8\x56\x89\x8a\x21\xc7\xbf\x62\x2f\x63\x7c\xd0\x4b\x7f\xce\x4d\xf0\x7e\xbe\xb1\x28\x27\xc7\x3c\xeb\xe2\xf4\x59\x79\x79\xce\x0c\xb8\x40\x89\x2a\x1c\x83\xad\x1e\xff\xf2\x65\xaf\x60\xbf\x54\xa5\x05\x98\x11\xee\xfb\xc2\xb7\x4d\x82\x8d\x0a\xaa\xec\xcd\x86\xc5\x8c\x4a\x04\xc0\x20\xae\x11\x7e\x30\x90\x59\x43\x1b\xb9\x44\xc6\xe5\x09\x2e\x1a\x02\x34\x7a\x1e\xfb\x12\x05\x8f\x44\x3c\xe4\x99\xce\x3a\x8a\xd8\x50\xd7\x1b\x5f\x86\xe5\xf8\x49\xff\x39\xea\x8e\x57\x5d\x82\x35\x19\x88\xe9\x95\xc7\x39\x63\xb7\xe4\xef\xf9\x52\xce\xfe\x05\x3f\xbb\x61\xc6\xa9\xd9\x4b\xbf\xf7\x12\x89\xfd\xcf\xb3\x7f\x79\xed\x9f\x0b\xb2\xec\x1d\xc5\x0d\x5b\x29\x0c\x22\x83\xe4\xde\xbf\x58\x52\xbc\xcb\x84\xc0\x84\x5c\x98\xd8\x9b\x23\x22\x4b\x70\x8e\x11\x87\xa2\x61\x51\xa1\xbd\xee\x1d\x8c\x1e\xfb\xad\xbc\x97\x21\x62\xc4\x47\x78\x80\x90\xfb\x83\xec\x8c\xc3\x3e\xf7\xfa\xd7\x16\x21\xc9\x4a\x0f\x48\x08\x3e\x71\x7f\xe2\x4f\xb1\xeb\xde\x92\xbf\xcf\xf9\x86\x0a\x48\x45\xa8\xd8\x9f\x48\x8a\xba\xfd\xdc\xa9\x41\x95\xac\xba\x04\x62\x88\xfb\x5e\xf4\xb6\x8b\xb4\xe8\x41\x11\x91\xae\x5c\x0d\x45\x6b\xaf\x50\x6e\x14\x57\x63\xa9\x20\xb4\x28\xe6\x89\xfa\xb9\x33\xd8\x98\xc0\xb7\xda\x82\xb4\x9b\xc0\x10\x82\x02\xbd\x7c\xf2\x6a\x58\x46\xd4\x7c\x4f\x01\xe3\x01\x35\x98\x98\x42\x6b\xe5\x29\x4b\x9e\x43\x44\x7d\x5a\xa0\x77\x8f\xbc\x11\x9c\x96\xfb\x4f\xdd\xa4\xdd\x77\xc7\xae\x5a\x5d\x24\x3e\xa8\x2c\xd0\x89\x91\x42\x4d\xd6\xa2\x73\x0b\x0c\x1e\xd7\xf7\x4b\xc6\xc9\xe3\x25\xb6\x16\xc1\xdc\x5e\xcc\x79\x52\xee\x15\x78\x57\x96\xbb\x2e\x50\x07\x5d\xf9\x74\xf5\xf7\x7b\x81\xfe\x7a\x8e\x0f\xf5\x7a\x2e\xf7\xfb\xfb\x7a\xd4\x3d\xd0\x78\xe5\x65\xbe\xe3\x59\x50\x34\x58\x64\xed\x57\x3c\x8b\x5e\xf9\xd3\x2d\xfc\x52\xfb\xde\x15\x47\xdd\xab\xd5\xe6\xa2\xaf\x7e\x79\x7a\x73\xf9\xb3\x05\x59\xbc\x46\xa8\x3b\xd2\x76\x48\xb8\x40\xe0\xfe\xb9\x97\x06\x3f\x61\x7d\xf9\x7f\xfd\xde\xff\x56\x09\xea\x7d\x26\x95\xc7\x24\xe4\x51\x2b\xf7\x5a\x35\xa3\xa9\xc6\x7f\x0e\xff\xbf\xbf\xd2\x3c\x85\x4d\x39\x01\xb3\xae\xac\x91\x58\x5e\x9a\x30\x56\xde\x7e\xbf\xf2\x6c\xc9\xbc\x1a\xed\xf4\x6a\xaf\x75\xdb\x7d\xee\x61\xb9\x26\x45\x89\x32\x64\xb8\x7d\x76\x5d\xd3\x9a\x15\xea\x63\x95\xfa\x8c\x9d\xf1\x80\xc2\xb1\xc7\x5a\x63\xbf\x23\xef\xbf\xad\x5c\x35\x4f\x61\x87\x0a\xbc\x57\x6b\xb1\xa6\xf5\x53\x1e\xf9\x58\x5b\x05\x39\x23\xd3\xd2\x5f\x3d\xef\xc5\xde\xe3\xda\x77\x3b\x6a\x4d\x3b\x42\x37\xab\x69\xc7\x53\x4f\x9e\xaf\x6a\x54\x56\xff\x3a\x5e\xc9\xa1\xf7\x9f\xe3\xd5\xf4\xe0\x2f\x90\xe5\x79\xee\xee\x73\x6f\x0e\xeb\x99\xe6\x39\xca\x4a\x7f\x9c\xfb\x8f\x7d\x15\x69\x25\x7f\xd4\x9b\x89\xec\xf0\x47\xe2\x87\xc6\x3f\x5c\x8f\x49\x9e\xc2\xdc\xf8\x39\xf7\xff\x1c\x3f\x35\x90\x93\x7b\x3f\xc8\x39\x28\xfa\x19\x3f\x43\xe9\xc9\xf7\xff\x08\xe7\x9a\xa3\xce\x51\xf4\x73\x96\xab\x88\x07\xf1\xf5\x1b\xae\xae\xd3\xf5\x6f\xa3\x61\xf3\xbd\x4e\x6c\x7c\x7a\x1a\x22\xec\xfd\xf3\x7a\xe1\xff\xb4\x5e\x45\x48\xf6\xad\xe3\x67\xf6\x0d\x6f\x82\xfc\x6f\xde\x84\x57\xf4\xc1\x13\x7d\x8a\x9f\x7a\xe9\x70\xb0\xee\xb7\xaf\x3c\x9e\x17\xb9\x04\x27\xbb\xa6\x92\xc9\x33\x6a\x1a\x4a\x03\x60\x49\xfe\x6d\x00\x83\xb4\x56\x46\x78\x19\x5b\x06\xd6\xb9\xd7\x24\x0f\xa6\x9b\x60\x3a\x61\x8d\x13\x4f\x50\x09\x91\xab\x4b\x19\x2d\xb5\xd7\x99\x6b\x36\x68\x4f\x14\x8d\x67\xfe\xcf\x0e\xb7\x68\x8d\x7a\x3c\xd7\x13\x9c\xdf\xf2\x14\xa7\x6a\xe8\xf1\xa7\x57\x2a\xfb\xda\xa0\x73\xef\x3d\x34\x3d\xa7\xa6\x1f\xa9\xe9\x19\x38\x94\x71\xb0\x30\xca\xc7\x7a\xcf\x7c\xf3\xa5\xfc\x16\xe9\x73\x7a\x4f\x47\x5d\xa5\x46\xd5\x97\xee\x0f\xb6\x96\x47\xf7\x47\x25\x2c\x33\x92\xae\xf4\xf0\xb8\x50\x2b\x14\x5a\x33\xf4\x4d\xe8\x7e\xf1\xeb\x94\x1c\x42\x2e\xf3\xe3\x77\x91\x06\xc4\x9d\x88\x52\xf5\x9f\x13\xa5\xa3\xe1\x59\x43\x7b\xc0\x80\xa8\xb3\xf5\x55\x87\xd9\x39\x9f\x19\xa3\xc1\x7d\x6e\x0e\xdf\x4e\x12\x06\xb8\x23\x16\xe8\xc8\xe8\xde\x7b\xe4\xa6\xa8\x6a\x3c\x2a\x94\x54\x74\xef\x73\x89\xaf\x5f\x7e\x7c\xfd\xf9\xaf\xbf\xfd\xed\x3f\x9f\x3f\x50\x17\x52\xfb\xe5\xc7\xd7\x5f\xff\xe3\x2f\x7f\xfc\xf7\x5f\xff\xf6\xeb\x5f\xfe\x1e\x2d\xbf\xfc\xf8\x3a\x0f\xbf\xfb\x6b\x36\x47\xe7\x3f\xfd\xfa\xdb\xaf\x5f\x7f\xf8\xf1\xa5\x1b\x24\xca\xa6\xa0\x9b\xb8\xb6\xeb\xb9\xf7\x73\x9d\xeb\x5c\xd7\xd5\xbe\xe1\xfb\xf5\xeb\x1a\xeb\x3d\x4b\x91\x1a\xc2\xf4\xbf\xff\xdf\xff\x04\x00\x00\xff\xff\xff\x3d\x3f\x81\x4f\x1f\x00\x00")

func dataStub_flasherJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataStub_flasherJson,
		"data/stub_flasher.json",
	)
}

func dataStub_flasherJson() (*asset, error) {
	bytes, err := dataStub_flasherJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/stub_flasher.json", size: 8015, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"data/stub_flasher.json": dataStub_flasherJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"stub_flasher.json": &bintree{dataStub_flasherJson, map[string]*bintree{}},
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