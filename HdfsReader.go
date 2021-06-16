// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.
package main

import (
	"errors"
	"github.com/colinmarc/hdfs/v2"
)

// Allows to open an HDFS file as a seekable read-only stream
// Concurrency: not thread safe: at most on request at a time
type HdfsReader struct {
	BackendReader *hdfs.FileReader
}

var _ ReadSeekCloser = (*HdfsReader)(nil) // ensure HdfsReader implements ReadSeekCloser

// Creates new instance of HdfsReader
func NewHdfsReader(backendReader *hdfs.FileReader) ReadSeekCloser {
	return &HdfsReader{BackendReader: backendReader}
}

// Read a chunk of data
func (this *HdfsReader) Read(buffer []byte) (int, error) {
	return this.BackendReader.Read(buffer)
}

// Seeks to a given position
func (this *HdfsReader) Seek(pos int64) error {
	actualPos, err := this.BackendReader.Seek(pos, 0)
	if err != nil {
		return err
	}
	if pos != actualPos {
		return errors.New("Can't seek to requested position")
	}
	return nil
}

// Returns current position
func (this *HdfsReader) Position() (int64, error) {
	actualPos, err := this.BackendReader.Seek(0, 1)
	if err != nil {
		return 0, err
	}
	return actualPos, nil
}

// Closes the stream
func (this *HdfsReader) Close() error {
	return this.BackendReader.Close()
}
