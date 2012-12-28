/*
 *	File I/O Wrapper
 *	----------------
 * 	Copyright (c) 2013, Scott Cagno, All rights reserved. 
 *	Use of this source code is governed by a BSD-style
 *	license that can be found in the LICENSE file.
 */

package filekit

import (
	"io/ioutil"
	"strings"
	"fmt"
	"os"
)

// filestore & methods
type FileStore struct {
	Pth string
	Ext string
}

// return a new file store instance
func NewFileStore(Pth, Ext string) *FileStore {
	if !strings.HasSuffix(Pth, "/") {
		Pth = Pth + "/"
	}
	if !strings.HasPrefix(Ext, ".") {
		Ext = "." + Ext
	}
	return &FileStore{Pth, Ext}
}

// write or append to a file, if file doesn't exist create new one
func (fs *FileStore) Write(filename, data string) {
	data = data + "\n"
	file, err := os.OpenFile(fs.Pth+filename+fs.Ext, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		if err = ioutil.WriteFile(fs.Pth+filename+fs.Ext, []byte(data), 0644); err != nil {
			fmt.Println(err)
		}
	} else {
		defer file.Close()
		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// write raw data to a new file
func (fs *FileStore) WriteRaw(filename string, data []byte) {
	if err := ioutil.WriteFile(fs.Pth+filename+fs.Ext, data, 0644); err != nil {
		fmt.Println(err)
	}
}

// read from a file, return slice of data newline seperated
func (fs *FileStore) Read(filename string) []string {
	var data []string
	bd, err := ioutil.ReadFile(fs.Pth + filename + fs.Ext)
	if err != nil {
		fmt.Printf("%s: not found!\n", fs.Pth+filename+fs.Ext)
		return data
	}
	data = strings.Split(string(bd), "\n")
	return data[:len(data)-1]
}

// read raw data from a file
func (fs *FileStore) ReadRaw(filename string) []byte {
	bd, err := ioutil.ReadFile(fs.Pth + filename + fs.Ext)
	if err != nil {
		fmt.Printf("%s: not found!\n", fs.Pth+filename+fs.Ext)
		return bd
	}
	return bd
}

// delete file by name
func (fs *FileStore) Delete(filename string) {
	if err := os.Remove(fs.Pth + filename + fs.Ext); err != nil {
		fmt.Printf("Could not zap file!\n%s: not found!\n", fs.Pth+filename+fs.Ext)
	}
}
