// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html (11.726kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x71\x73\xdb\xb6\x92\xff\xdb\xf9\x14\x5b\x5e\xfc\x24\x9d\x4d\x52\xb6\x93\x3c\x9f\x44\xaa\x93\x97\xd7\xd7\xcb\xcd\x5d\xdb\x69\xd3\xb9\x7b\xd3\xd7\xb9\x01\x89\x95\x88\x18\x04\x58\x00\x94\xac\x7a\xf4\xdd\x6f\x00\x90\x14\x29\xc9\x6e\xd2\xe4\xde\x34\x7f\x38\x24\xb0\xd8\x5d\xec\xfe\x16\xbb\x58\x2a\xf9\xe2\xaf\xdf\xbe\x79\xf7\xf7\xef\xbe\x82\xc2\x94\x7c\xf1\x2c\xb1\xff\x01\x27\x62\x95\x06\x28\x82\xc5\xb3\xb3\xa4\x40\x42\x17\xcf\xce\xce\x92\x12\x0d\x81\xbc\x20\x4a\xa3\x49\x83\xda\x2c\xc3\xdb\x60\x3f\x51\x18\x53\x85\xf8\x4b\xcd\xd6\x69\xf0\x3f\xe1\x8f\xaf\xc3\x37\xb2\xac\x88\x61\x19\xc7\x00\x72\x29\x0c\x0a\x93\x06\x6f\xbf\x4a\x91\xae\xb0\xb7\x4e\x90\x12\xd3\x60\xcd\x70\x53\x49\x65\x7a\xa4\x1b\x46\x4d\x91\x52\x5c\xb3\x1c\x43\xf7\x72\x09\x4c\x30\xc3\x08\x0f\x75\x4e\x38\xa6\x57\xc1\xe2\x99\xe5\x63\x98\xe1\xb8\x78\x78\x88\xbe\x41\xb3\x91\xea\x6e\xb7\x9b\xc1\xeb\xda\x14\x28\x0c\xcb\x89\x41\x0a\x7f\x23\x75\x8e\x26\x89\x3d\xa5\x5b\xc4\x99\xb8\x83\x42\xe1\x32\x0d\xac\xea\x7a\x16\xc7\x39\x15\xef\x75\x94\x73\x59\xd3\x25\x27\x0a\xa3\x5c\x96\x31\x79\x4f\xee\x63\xce\x32\x1d\x9b\x0d\x33\x06\x55\x98\x49\x69\xb4\x51\xa4\x8a\x6f\xa2\x9b\xe8\xcf\x71\xae\x75\xdc\x8d\x45\x25\x13\x51\xae\x75\x00\x0a\x79\x1a\x68\xb3\xe5\xa8\x0b\x44\x13\x40\xbc\xf8\x7d\x72\x97\x52\x98\x90\x6c\x50\xcb\x12\xe3\x17\xd1\x9f\xa3\xa9\x13\xd9\x1f\x7e\x5a\xaa\x15\xab\x73\xc5\x2a\x03\x5a\xe5\x1f\x2c\xf7\xfd\x2f\x35\xaa\x6d\x7c\x13\x5d\x45\x57\xcd\x8b\x93\xf3\x5e\x07\x8b\x24\xf6\x0c\x17\x9f\xc4\x3b\x14\xd2\x6c\xe3\xeb\xe8\x45\x74\x15\x57\x24\xbf\x23\x2b\xa4\xad\x24\x3b\x15\xb5\x83\x9f\x4d\xee\x63\x3e\x7c\x7f\xe8\xc2\xcf\x21\xac\x94\x25\x0a\x13\xbd\xd7\xf1\x75\x74\x75\x1b\x4d\xdb\x81\x63\xfe\x4e\x80\x75\x9a\x15\x75\x16\xad\x51\x59\xe4\xf2\x30\x47\x61\x50\xc1\x83\x1d\x3d\x2b\x99\x08\x0b\x64\xab\xc2\xcc\xe0\x6a\x3a\x3d\x9f\x9f\x1a\x5d\x17\x7e\x98\x32\x5d\x71\xb2\x9d\xc1\x92\xe3\xbd\x1f\x22\x9c\xad\x44\xc8\x0c\x96\x7a\x06\x9e\xb3\x9b\xd8\x39\x99\x95\x92\x2b\x85\x5a\x37\xc2\x2a\xa9\x99\x61\x52\xcc\x2c\xa2\x88\x61\x6b\x3c\x45\xab\x2b\x22\x8e\x16\x90\x4c\x4b\x5e\x1b\x3c\x50\x24\xe3\x32\xbf\xf3\x63\x2e\x9a\xfb\x9b\xc8\x25\x97\x6a\x06\x9b\x82\x35\xcb\xc0\x09\x82\x4a\x61\xc3\x1e\x2a\x42\x29\x13\xab\x19\xbc\xaa\x9a\xfd\x40\x49\xd4\x8a\x89\x19\x4c\xf7\x4b\x92\xb8\x35\x63\x12\xfb\x83\xeb\xd9\x59\x92\x49\xba\x75\x3e\xa4\x6c\x0d\x39\x27\x5a\xa7\xc1\x81\x89\xdd\x81\x34\x20\xb0\xe7\x10\x61\xa2\x9d\x1a\xcc\x29\xb9\x09\xc0\x09\x4a\x03\xaf\x44\x98\x49\x63\x64\x39\x83\x2b\xab\x5e\xb3\xe4\x80\x1f\x0f\xf9\x2a\xbc\xba\x6e\x27\xcf\x92\xe2\xaa\x65\x62\xf0\xde\x84\xce\x3f\x9d\x67\x82\x45\xc2\xda\xb5\x4b\x02\x4b\x12\x66\xc4\x14\x01\x10\xc5\x48\x58\x30\x4a\x51\xa4\x81\x51\x35\x5a\x1c\xb1\x05\xf4\x8f\xbf\x47\x4e\xbf\xe2\xaa\xd5\x2b\xa6\x6c\xdd\x6c\xab\xf7\x78\xb0\xc3\xc7\x37\x71\x0b\xcd\x83\x5c\x2e\x35\x9a\xb0\xb7\xa7\x1e\x31\x13\x55\x6d\xc2\x95\x92\x75\xd5\xcd\x9f\x25\x6e\x14\x18\x4d\x83\x5a\xf1\xa0\x39\xfe\xdd\xa3\xd9\x56\x8d\x29\x82\x6e\xe3\x52\x95\xa1\xf5\x84\x92\x3c\x80\x8a\x93\x1c\x0b\xc9\x29\xaa\x34\xf8\x41\xe6\x8c\x70\x10\x7e\xcf\xf0\xe3\xf7\xff\x09\x8d\xcb\x98\x58\xc1\x56\xd6\x0a\xbe\x32\x05\x2a\xac\x4b\x20\x94\x5a\xb8\x46\x51\xd4\x53\xc4\x61\xf7\x58\xd5\x30\x33\x62\x4f\x75\x96\x64\xb5\x31\xb2\x23\xcc\x8c\x80\xcc\x88\x90\xe2\x92\xd4\xdc\x00\x55\xb2\xa2\x72\x23\x42\x23\x57\x2b\x9b\xe9\xfc\x26\xfc\xa2\x00\x28\x31\xa4\x99\x4a\x83\x96\xb6\xf5\x21\xd1\x95\xac\xea\xaa\xf1\xa2\x1f\xc4\xfb\x8a\x08\x8a\xd4\xfa\x9c\x6b\x0c\x16\x5f\xb3\x35\x42\x89\x7e\x2f\x67\x87\x90\xc8\x89\x42\x13\xf6\x99\x1e\x01\x23\x89\xbd\x32\x7e\x4b\xd0\xfc\x4b\x6a\xde\x72\xea\xb6\x50\xa2\xa8\x61\xf0\x16\x2a\x7b\xae\x04\x8b\x87\x07\x45\xc4\x0a\xe1\x39\xa3\xf7\x97\xf0\x9c\x94\xb2\x16\x06\x66\x29\x44\xaf\xdd\xa3\xde\xed\x06\xdc\x01\x12\xce\x16\x09\x79\x0a\xde\x20\x45\xce\x59\x7e\x97\x06\x86\xa1\x4a\x1f\x1e\x2c\xf3\xdd\x6e\x0e\x0f\x0f\x6c\x09\xcf\xa3\xef\x31\x27\x95\xc9\x0b\xb2\xdb\xad\x54\xfb\x1c\xe1\x3d\xe6\xb5\xc1\xf1\xe4\xe1\x01\xb9\xc6\xdd\x4e\xd7\x59\xc9\xcc\xb8\x5d\x6e\xc7\x05\xdd\xed\xac\xce\x8d\x9e\xbb\x1d\xc4\x96\xa9\xa0\x78\x0f\xcf\xa3\xef\x50\x31\x49\x35\x78\xfa\x24\x26\x8b\x24\xe6\x6c\xd1\xac\x1b\x1a\x29\xae\xf9\x1e\x2f\xb1\x05\x4c\x87\x73\x17\x36\x4e\xd5\xbe\xa6\x27\xa2\x60\x15\x76\xda\x37\x78\xd0\xcc\xe0\x1d\x6e\xd3\xe0\xe1\xa1\xbf\xb6\x99\xcd\x09\xe7\x19\xb1\x76\xf1\x5b\xeb\x16\xfd\x8a\x16\xa7\x6b\xa6\x5d\x49\xb5\x68\x35\xd8\xab\xfd\x81\x61\x7d\x70\x70\x19\x59\xcd\xe0\xe6\xba\x77\x6a\x9d\x8a\xf8\x57\x07\x11\x7f\x73\x92\xb8\x22\x02\x39\xb8\xbf\xa1\x2e\x09\x6f\x9f\x9b\x68\xe9\x05\xdf\xe1\xa2\xd0\x9e\xd1\x9d\x6a\xdd\x59\x3f\x9d\x83\x5c\xa3\x5a\x72\xb9\x99\x01\xa9\x8d\x9c\x43\x49\xee\xbb\x7c\x77\x33\x9d\xf6\xf5\xb6\xa5\x20\xc9\x38\xba\xd3\x45\xe1\x2f\x35\x6a\xa3\xbb\xb3\xc4\x4f\xb9\xbf\xf6\x48\xa1\x28\x34\xd2\x03\x6b\x58\x89\xd6\xb4\x8e\xaa\xe7\xfa\xce\x98\x27\x75\x5f\x4a\xd9\xa5\x90\xbe\x1a\x0d\xeb\x5e\xb6\x0b\x16\x89\x51\x7b\xba\xb3\xc4\xd0\x8f\x4a\x01\xca\x96\x78\x8f\x65\x00\x7f\xa2\xd9\xbd\x57\x88\xca\xd7\x17\x16\xb2\xe0\x5e\x93\xd8\xd0\x4f\x90\x6c\x41\x98\x11\x8d\x1f\x22\xde\x65\xfa\xbd\x78\xf7\xfa\xa9\xf2\x0b\x24\xca\x64\x48\xcc\x87\x28\xb0\xac\x05\xed\xed\xdf\x9d\x9d\x9f\xaa\x40\x2d\xd8\x1a\x95\x66\x66\xfb\xa1\x1a\x20\xdd\xab\xe0\xdf\x87\x2a\x24\xb1\x51\x4f\x63\xad\xff\xf2\x99\x82\xfb\xb7\x4a\x92\x9b\xc5\xbf\xcb\x0d\x50\x89\x1a\x4c\xc1\x34\xd8\xe4\xfa\x65\x12\x17\x37\x1d\x49\xb5\x78\x67\x27\x9c\x51\x61\xe9\x4a\x0b\x60\x1a\x54\x2d\x5c\xe6\x95\x02\x4c\x81\xc3\x72\xa4\x49\xd2\x11\xbc\x93\xb6\xa4\x5b\xa3\x30\x50\x12\xce\x72\x26\x6b\x0d\x24\x37\x52\x69\x58\x2a\x59\x02\xde\x17\xa4\xd6\xc6\x32\xb2\xc7\x07\x59\x13\xc6\x5d\x2c\x39\x97\x82\x54\x40\xf2\xbc\x2e\x6b\x5b\x92\x8a\x15\xa0\x90\xf5\xaa\x68\x74\x31\x12\x7c\x62\xe2\x52\xac\x3a\x7d\x74\x45\x4a\x20\xc6\x90\xfc\x4e\x5f\x42\x7b\x2a\x00\x51\x08\x86\x21\xb5\xab\x72\x59\x96\x52\xc0\x8d\xa2\x50\x11\x65\xb6\xa0\x87\xb5\x05\xc9\x73\x97\xe5\x22\x78\x2d\xb6\x52\x20\x14\x64\xed\x34\x84\x77\xfe\x3a\x71\x09\x5f\x4b\xb9\xe2\x78\x61\x15\xfc\x1b\xc9\x31\x93\xb2\x5b\x06\x25\xd9\xb6\x72\x9b\x6d\x6c\x98\x29\x98\xb7\x53\x85\xaa\xb4\x3c\x28\x70\x56\x32\xa3\xa3\x24\xae\xf6\x47\xeb\x3e\x49\xf3\xb0\x90\x8a\xfd\x6a\x2b\x1c\xde\x3f\x47\xcd\xc1\x29\xd3\x1e\x92\xce\xfd\x1c\x97\x66\x06\x2f\xfc\x21\x79\x08\xe8\xe6\x2a\x74\x0a\xcd\x2d\x4f\x77\xc5\xb4\x99\x67\x06\x37\xbe\xae\xf5\x15\x05\x35\x3d\x0d\xe8\x01\xe6\xbc\xd0\xdb\xdb\xea\xbe\xd3\xa3\x2b\x8e\xa7\x1d\x13\x0b\x85\xa1\x51\xd6\xac\x67\xcf\x92\xdc\x21\x10\x48\xc8\xc1\x55\xb9\x51\xda\x5d\xb4\x98\x6b\x14\xc4\x66\x83\x68\xbe\xb4\x31\x9c\x7e\xef\x19\x32\xb1\x3a\xbf\x9e\x7a\x68\xda\x07\xcb\xfe\xfc\x7a\xca\x84\x91\xe7\xd7\xd3\xe9\xfd\xf4\x03\xff\x9d\x5f\x4f\xa5\x38\xbf\x9e\x9a\x02\xcf\xaf\xa7\xe7\xd7\x37\x7d\x50\xfb\x91\xb6\xc4\xb4\x54\xa8\xad\xb4\x16\xeb\x01\x18\xa2\x56\x68\xd2\xe0\x7f\x49\x26\x6b\x33\xcb\x38\x11\x77\xc1\xc2\xa9\x6b\xcb\x0e\x87\x82\xd3\x85\x2a\x54\x44\x5b\x48\x58\x8d\x1d\x4a\x9a\xa6\x88\x86\xb1\xae\x95\x92\xb5\xb0\xe9\x11\xec\x9e\x5d\xa8\x8a\x91\x45\x99\x35\xcc\x24\x4a\x32\x15\x2f\xde\xc8\x6a\x1b\x3a\x26\x6e\xf9\x91\x19\x75\x5d\x55\x52\x99\xa8\x6f\x4e\x62\x2f\x44\x1c\x75\x7c\x3b\x7d\x79\xfb\xea\x49\xf5\xb5\x2d\xb7\xdd\x1e\x3a\x0d\x49\x26\xd7\x08\xbe\xb8\xcf\xe4\x3d\x10\x41\x61\xc9\x14\x02\xd9\x90\xed\x17\x49\x4c\xdd\x55\xec\xd3\x51\xbb\x72\x81\x16\x56\xbc\xd6\xb6\x16\x61\x36\x50\xff\x50\x10\xf6\x27\x01\x7c\xc7\x6b\x7d\x09\x55\x9d\x71\xa6\x0b\x20\x20\x70\x03\x89\x36\x4a\x8a\xd5\xc2\x8d\xe6\xf6\xaa\xea\x5e\xa1\x92\xda\x3c\x85\x06\x2c\x33\xa4\xf4\x04\x1e\x7e\x27\x1c\xac\x3c\xe7\xc2\x7f\xbe\xfb\x96\xcd\xe1\xf8\x87\x72\x59\x7b\x62\xff\x51\xfd\x75\x14\xbe\x9b\xcd\x26\x6a\x2d\xe9\x62\xb7\x40\x5e\xc5\x36\x8d\xd5\x82\x99\x6d\xec\x4f\x41\x29\xe2\x2f\x19\x4d\xaf\x6f\xaf\x5f\xbd\xba\x7e\xf1\x6f\xb7\x2f\x5f\x5e\xdf\xbe\x78\xf9\x58\x60\x77\xa0\xf8\xfd\x71\xed\xaf\x43\xdf\xc8\xd7\xb5\x29\xba\xbb\x90\xc7\x4b\x5b\x83\xdb\x4a\x8b\xda\xbb\xa4\x0a\x7e\x37\x86\x6a\x61\x0b\xca\x90\xf0\x93\xb5\xe0\x47\xa0\xc8\xc1\xe8\x09\xcd\x3e\x11\x5a\x2d\x7c\x2c\x52\x64\x6d\xec\x0e\xdb\xa6\x0c\x93\xa2\x83\xd3\x25\x68\x56\x56\x7c\x0b\xf9\xde\xeb\xa7\x71\xf5\xa8\x53\x7e\x13\x56\x43\xb7\x79\x90\xb9\x2a\xae\x94\x14\x6d\xf5\xa6\x6b\x9d\x63\xe5\xba\xf5\xb6\x22\xfa\xcb\xf6\x57\x22\x0c\x13\xd8\x56\x4e\x11\x7c\x2b\xf8\x16\x6a\x8d\xb0\x94\x0a\x28\x66\xf5\x6a\xe5\xca\x3d\x05\x95\x62\x6b\x62\xb0\x2d\x97\x74\x83\x8a\x0e\x14\xbd\x1b\xaa\x2d\x5d\x79\xaf\x92\xfc\xbb\xac\x21\x27\x02\x8c\x22\xf9\x9d\x8f\x94\x5a\x29\x1b\x29\x15\xfa\xdd\x74\x05\x5b\x86\x5c\x6e\x1c\x89\xdf\xf7\x92\x21\x77\xd5\x9b\x46\x84\x42\x6e\xa0\xac\x73\x17\x90\xb6\x3a\x73\x9b\xd8\x10\x66\xa0\x16\x86\x71\x6f\x4f\x53\x2b\x61\x6b\x3d\x1c\x14\x59\x47\x77\xf8\x04\xcb\xc5\xbb\x02\x4f\x94\xb6\xdd\xed\x1b\x14\xbe\xf1\xe4\x50\x29\x69\x30\xb7\x0e\x05\xb2\x22\x4c\x68\xeb\x11\x57\xc6\x61\xf9\x01\xb7\xf3\xee\xa9\x79\xd8\x77\x9a\xdd\x74\x1c\xc3\xd7\x5c\x66\x84\xc3\xda\x22\x3d\xe3\xb6\x2c\x97\x50\x48\xbb\xf5\x9e\xb5\xb4\x21\xa6\xd6\x20\x97\x6e\xd4\x6b\x6e\xd7\xaf\x89\xb2\x1e\xc4\xb2\x32\x90\x36\x7d\x52\x3b\xa6\x51\xad\x9b\xee\xaf\x7d\x35\x0c\xd5\x60\xbe\xb3\x7a\x0a\x3f\xfd\x3c\x7f\xd6\xa8\xf2\x57\x5c\x3a\x48\x58\x7c\xfb\x2d\x9b\x82\x18\xc8\x15\x12\x83\x1a\x72\x2e\x75\xad\xbc\x86\x54\xc9\x0a\xac\x96\x2d\xa7\x96\xb3\x9d\xa8\x9c\xb4\x96\xc9\xb8\x20\xba\x98\x34\x6d\x5e\x85\xce\x4b\xdd\x5c\x3b\x7e\x66\x51\x37\xb6\x0c\x58\x3a\x9d\x03\x4b\x5a\xbe\x11\x47\xb1\x32\xc5\x1c\xd8\xc5\x45\x47\x7c\xc6\x96\x30\x6e\x29\x7e\x62\x3f\x47\xe6\x3e\xb2\x52\x20\x4d\xa1\x2f\xcd\x09\x6c\xf8\xe8\x8a\xb3\x1c\xc7\xec\x12\xae\x26\xf3\x76\x36\x53\x48\xee\xda\xb7\xc6\x8f\xfe\x3f\xf7\x77\x37\x1f\x5a\xc6\x19\x7f\x60\x1b\xdf\xc3\xd1\x40\x60\xc5\xb4\x81\x5a\x71\x68\x62\xd8\xbb\xa0\x73\x88\xa3\xeb\x5b\xe5\x08\x97\xcd\x43\x83\xa9\x76\x0b\x9e\x4d\xa4\x51\xd0\xf1\x7f\xfc\xf0\xed\x37\x91\x36\x8a\x89\x15\x5b\x6e\xc7\x0f\xb5\xe2\x33\x78\x3e\x0e\xfe\xa5\x56\x3c\x98\xfc\x34\xfd\x39\x5a\x13\x5e\xe3\xa5\xf3\xf7\xcc\xfd\x3d\x92\x72\x09\xcd\xe3\x0c\x86\x02\x77\x93\xc9\xfc\x74\xbf\xab\xd7\x9e\x53\xa8\xd1\x8c\x2d\x61\x07\xfc\x43\x1b\x11\x28\xd1\x14\xd2\x85\xae\xc2\x5c\x0a\x81\xb9\x81\xba\x92\xa2\x31\x09\x70\xa9\xf5\x1e\x88\x2d\x45\x7a\x0c\x8a\x86\x3e\x75\xc9\xfa\xbf\x31\xfb\x41\xe6\x77\x68\xc6\xe3\xf1\x86\x09\x2a\x37\x11\x97\xfe\xa8\x8d\x6c\x90\xca\x5c\x72\x48\xd3\x14\x9a\x2c\x1a\x4c\xe0\x4b\x08\x36\xda\xe6\xd3\x00\x66\xf6\xd1\x3e\x4d\xe0\x02\x0e\x97\x17\x36\xdf\x5f\x40\x10\x93\x8a\x05\x13\x1f\x0e\xad\xe1\xa5\x28\x51\x6b\xb2\xc2\xbe\x82\xee\x86\xdb\x81\xcc\xee\xa3\xd4\x2b\x48\xc1\x39\xa8\x22\x4a\xa3\x27\x89\x28\x31\xa4\x45\x9b\xc5\xac\x23\x4b\x53\x10\x35\xe7\x7b\x90\xfa\xa0\x98\xb7\xf0\x1b\x90\x47\x3e\xd7\x7c\x91\xa6\x50\x0b\xea\x4c\x4c\xf7\x2b\xad\xf3\x7d\x33\x64\x12\xd9\xbc\xb0\x5f\x31\x99\xf7\xd1\x3c\xe0\x86\xf4\xb7\xd8\x21\x3d\xe4\x87\xf4\x11\x86\xae\xf7\xf4\x14\x3f\xdf\xab\xea\xb1\x73\x03\x8f\x70\x13\x75\x99\xa1\x7a\x8a\x9d\xef\x3d\x35\xec\x9c\xa9\xdf\x0a\xd3\x5b\x7b\x09\x57\xaf\x26\x8f\x70\x47\xa5\xe4\xa3\xcc\x85\x34\xdb\xf1\x03\x27\x5b\x5b\x33\xc1\xc8\xc8\xea\x8d\x6b\x15\x8d\x2e\x5d\xc6\x9d\x41\xc7\xe1\xd2\x7d\x04\x98\xc1\xc8\xbd\xd9\x79\x56\xa2\x5b\xf5\x72\x3a\x9d\x5e\x42\xfb\xf5\xec\x2f\xc4\x06\xa1\xaa\x71\xf7\x88\x3e\xba\xce\x73\x9b\xf7\x3f\x45\xa3\x86\x47\xa7\x53\xf3\xfe\x09\x5a\x75\xb9\x61\xa0\x16\xfc\xe9\x4f\x70\x34\x3b\x84\x71\x1c\xc3\x7f\x11\x75\xe7\x1a\x3b\x95\xc2\xb5\x6b\xfe\x74\xf4\x25\xd3\xda\x35\x55\x34\x50\x29\xb0\x59\xf3\x71\xc7\xfe\x91\x8e\x0d\x19\x2c\x60\x7a\xa8\xa0\x3d\x0e\x7b\x69\xe1\x44\xb6\xe8\xf1\x1d\x26\x82\xb3\x5d\x5f\xde\x60\x25\x2b\x11\xbe\x48\x21\x08\xfa\x8b\x8f\x28\x2c\x41\xc7\xec\x4c\xa3\x79\xe7\x7d\x31\x6e\xb2\xe3\xa9\xdc\x35\xb9\x84\x9b\xe9\x74\x3a\x39\x52\x62\xb7\x37\xef\xeb\xca\x96\x4d\x40\xc4\xd6\x1d\x89\x9d\x6d\x5d\xe1\x68\x4b\x20\x7b\xa4\x71\xc8\x25\xe7\xbe\x66\x69\x96\x5a\x03\x37\x4d\xb0\x14\xc2\xab\xf9\x89\x2c\xda\xb3\x64\x6f\x6b\x87\xee\x39\x61\xfb\x43\x17\x0d\x6d\x76\x40\x1c\x5e\x0d\x9c\x32\xf0\xd7\x69\xc7\x9c\x75\x7a\xb3\xbd\x45\x0f\xdc\xb5\xf7\xd7\xa1\xcd\x7a\xfa\x7b\x3e\x17\x57\x1f\xb8\x8d\x6e\xba\xaa\x75\x31\x3e\x50\x74\x32\x3f\xf6\xcd\x5b\x83\xca\x56\xc9\xd2\xa6\x2c\xeb\x0b\x7b\x15\x50\x78\xe4\x12\x57\xaa\x2b\x0c\x15\x0a\x8a\xaa\x2d\x29\x7c\x65\x6f\x0b\xc0\x81\xcb\xfc\xad\xb2\x0f\xa7\x8f\x0c\x18\x57\x92\x49\x81\x00\x00\x07\x41\xe0\x80\x3a\x40\xaa\x25\x46\x4e\x2a\x8d\x14\x52\xf0\x3f\x66\x18\x4f\xa2\x5a\xb0\xfb\xf1\x24\x6c\xde\x0f\x79\xb4\xf3\xf3\xee\x9a\xd8\xaa\x7d\x91\x42\x90\x18\x05\x8c\xa6\xa3\x00\x2e\x4e\x85\xa0\xcd\xba\xa3\xc5\x5e\x83\xfe\x52\x80\xc4\xd0\x85\xeb\x67\xfb\xfb\xda\x3f\x82\x8c\xe4\x77\x2b\x77\x11\x9a\xd9\x52\x6b\x7c\xc4\x96\xac\x89\x21\xca\x71\x9d\xcc\x61\x4f\xde\x5c\x14\x73\xeb\x9c\x39\xf8\x1b\xa9\x6b\x9b\x43\xf7\xa9\xc9\xbd\x65\x52\x51\x54\xa1\x22\x94\xd5\x7a\x06\x2f\xaa\xfb\xf9\x3f\xda\x4f\x71\xae\xb9\xff\xa4\xaa\x95\xc2\xc5\x91\x46\x4d\x93\xf8\x02\x82\x24\xb6\x04\xbf\xc5\xa6\xdb\x6c\xff\x47\x14\x70\xe2\x13\x06\x74\x3f\x71\x68\xc6\x4b\x46\x29\x47\xab\xf0\x9e\xbd\x0d\x46\xeb\xff\x7e\x48\x0d\x45\x42\xf3\xed\x62\xbf\x66\x07\xc8\x35\x3e\xb1\xa0\xfb\x0c\x32\xb2\x00\x08\xed\x96\x99\xb3\x79\x73\xd9\x76\xc3\x6a\xe4\x6c\xd1\xfc\x24\x86\xd6\xca\xd5\x5a\xe3\xb0\x01\xd8\x25\x8c\xb4\xad\xfd\xa8\x1e\x4d\xa2\xa2\x2e\x89\x60\xbf\xe2\xd8\xe6\xa5\x89\xb7\x95\xfb\xae\x12\x1c\x1f\xc9\x47\xca\xec\x3f\x78\x8c\xda\x1c\x37\x6a\x8c\x38\x6a\xbd\xfb\x62\x7f\xb7\x9f\xc1\x74\x3e\xfa\x48\x0b\x9d\x96\x12\x66\x44\x41\xff\x25\x6c\x93\x2f\x28\x69\xa5\xb7\x73\x19\x51\x23\xdf\xc9\x70\xf5\xb9\x90\x9b\x74\x74\x33\xed\x94\xf4\x8e\x76\x7e\x1e\x35\x58\x3b\x72\x86\xd5\xb2\x0d\xcd\x05\xdc\x4c\x3f\x87\xb6\xbe\x1b\x72\xb0\x03\xa3\x58\x85\x14\x48\x6e\xd8\x1a\xff\x1f\x36\xf2\x19\x8c\xfc\xd1\x2a\x5a\x1c\xb6\xc6\x73\x30\x1d\xe8\x6b\x67\x3b\xdb\xfe\xab\x8d\x37\x88\x9d\x85\x2f\x20\x38\xb9\x91\x47\x91\x78\x40\x78\x10\xda\x8f\xc7\xbd\xfb\x50\x18\x1c\xe6\x14\x5b\xed\x76\x1f\xb9\x27\x51\x61\x4a\x3e\x0e\x12\xe3\x7e\xec\x64\x75\xee\x38\x38\x06\x7e\x78\x58\xd2\xed\x86\x17\x19\x7b\x7f\xc7\x83\x7b\x16\xf4\x8a\x93\xee\x2e\xd6\x56\x22\xb0\xdb\xff\x26\x2c\x8e\xe1\x07\x43\x94\x01\x02\x3f\xbe\x85\xba\xa2\xc4\xf8\x4f\x72\x36\x3f\xfa\xae\x73\xfb\xa3\xb1\x8c\x28\x0d\x4b\xa9\x36\x44\xd1\xa6\x3f\x63\x0a\xdc\xba\x4f\x72\x6d\xe9\xa7\xd1\xbc\xb5\xa7\xd8\x9a\xf0\xf1\xd1\xbd\xef\xf9\x78\x14\xf5\x5d\x3e\x9a\x44\x48\xf2\xe2\x98\xd0\x65\xac\x4e\x6e\x0a\xdf\xb8\x2b\xc0\xf8\xf9\xd8\x14\x4c\x4f\x22\x62\x8c\x1a\x8f\x06\x60\x18\x4d\xac\x5f\xaf\x7a\x57\xb2\x6e\x79\x32\x08\xab\xa7\x78\xec\x8b\xe9\xae\x10\x68\xc9\x73\xad\xc7\x1e\x57\xa3\xcb\x1e\xef\x21\xac\x46\xe7\xa3\xce\x51\xfb\xf0\xde\xef\x23\x3d\xa9\xc9\x80\xf5\xc8\x46\xd9\xe8\x48\x3c\xa1\xf4\x8d\x8d\x9f\x71\x70\x22\xd2\x0f\xd1\x31\xe9\x8c\xed\xcf\xeb\x27\xad\xec\x7f\x5e\xf3\x88\x89\x19\x1d\x4d\x22\x5d\x67\xbe\x37\x31\x7e\xd9\x5d\xc0\x5a\x32\x07\xde\xc3\x54\x70\x54\x50\x58\x11\xc3\xa2\x22\x3c\x28\x42\x9e\xc8\x1a\x8d\x48\xbf\xab\xdd\xa5\x35\xf8\x74\xd2\xb5\xb6\xbe\xd2\xb6\xb8\xf2\xad\xff\x0d\x66\xda\x75\x12\xa0\xc1\xbb\xeb\xe6\xf8\xae\xcd\xeb\xef\xde\xf6\x3a\x37\x5d\x44\x8c\x1d\xf7\xee\xf7\x9c\xa7\xfa\x24\x27\x7f\x40\xba\xd9\x6c\x22\xff\x45\xcb\xb5\xf1\xbb\x46\x4a\x4c\x2a\x16\xbd\xd7\x01\x10\xbd\x15\x39\x50\x5c\xa2\x5a\xf4\xd8\x37\xdd\x95\x24\xf6\x3f\x6d\x4c\x62\xff\xeb\xed\xff\x0b\x00\x00\xff\xff\x56\xf8\xb5\xef\xce\x2d\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9b, 0x9e, 0x78, 0xa7, 0x98, 0x53, 0xe5, 0xcc, 0x18, 0xfc, 0x7e, 0x9a, 0xb2, 0xa6, 0xa5, 0x82, 0xbc, 0x1e, 0xb5, 0x94, 0xa5, 0x5d, 0xfe, 0x42, 0x78, 0x75, 0x6d, 0x98, 0x6d, 0x26, 0x43, 0x8d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"faucet.html": faucetHtml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
