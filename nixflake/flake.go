// Copyright (c) Will Drengwitz <ghthor@gmail.com>
// SPDX-License-Identifier: MIT

package nixflake

type Lock struct {
	Nodes   map[string]LockNode `json:"nodes"`
	Root    string              `json:"root"`
	Version int                 `json:"version"`
}

type LockNode struct {
	Inputs   map[string]any   `json:"inputs"`
	Locked   LockNodeLocked   `json:"locked"`
	Original LockNodeOriginal `json:"original"`
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
	Repo  string `json:"repo,omitempty"`
	Ref   string `json:"ref,omitempty"`
	Type  string `json:"type,omitempty"`
	Url   string `json:"url,omitempty"`
}
