// Copyright (c) Will Drengwitz <ghthor@gmail.com>
// SPDX-License-Identifier: MIT

package nixflake

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Lock struct {
	Nodes   map[string]LockNode `json:"nodes"`
	Root    string              `json:"root"`
	Version int                 `json:"version"`
}

type LockNode struct {
	Inputs   map[string]any    `json:"inputs,omitempty"`
	Locked   *LockNodeLocked   `json:"locked,omitempty"`
	Original *LockNodeOriginal `json:"original,omitempty"`
}

type LockNodeLocked struct {
	LastModifiedUnix int64  `json:"lastModified"`
	NarHash          string `json:"narHash"`
	Dir              string `json:"dir,omitempty"`
	Owner            string `json:"owner,omitempty"`
	Repo             string `json:"repo,omitempty"`
	Ref              string `json:"ref,omitempty"`
	Rev              string `json:"rev,omitempty"`
	RevCount         int64  `json:"revCount,omitempty"`
	Type             string `json:"type,omitempty"`
	Url              string `json:"url,omitempty"`
}

type LockNodeOriginal struct {
	Dir   string `json:"dir,omitempty"`
	Id    string `json:"id,omitempty"`
	Owner string `json:"owner,omitempty"`
	Ref   string `json:"ref,omitempty"`
	Repo  string `json:"repo,omitempty"`
	Type  string `json:"type,omitempty"`
	Url   string `json:"url,omitempty"`
}

func Parse(from io.Reader) (*Lock, error) {
	var lock Lock
	if err := json.NewDecoder(from).Decode(&lock); err != nil {
		return nil, err
	}
	return &lock, nil
}

func ParseFile(path string) (*Lock, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Parse(file)
}

func (l *Lock) String() string {
	var (
		b   strings.Builder
		enc = json.NewEncoder(&b)
	)
	enc.SetIndent("", "  ")
	err := enc.Encode(l)
	if err != nil {
		return err.Error()
	}
	return b.String()
}

var _ io.WriterTo = (*Lock)(nil)

func (l *Lock) WriteTo(w io.Writer) (n int64, err error) {
	var (
		buf bytes.Buffer
		enc = json.NewEncoder(&buf)
	)

	enc.SetIndent("", "  ")
	err = enc.Encode(l)
	if err != nil {
		return 0, err
	}

	return buf.WriteTo(w)
}
