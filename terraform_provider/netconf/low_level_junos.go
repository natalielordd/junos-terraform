// Copyright (c) 2018, Juniper Networks, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netconf

import (
	"os/exec"
)

// TransportJunos maintains the information necessary to communicate with Junos
// via local shell NETCONF interface.
type TransportJunos struct {
	TransportBasicIO
	cmd *exec.Cmd
}

// Close closes an existing local NETCONF session.
func (t *TransportJunos) Close() error {
	if t.cmd != nil {
		t.ReadWriteCloser.Close()
	}
	return nil
}

// Open creates a new local NETCONF session.
func (t *TransportJunos) Open() error {
	var err error

	t.cmd = exec.Command("xml-mode", "netconf", "need-trailer")

	w, err := t.cmd.StdinPipe()
	if err != nil {
		return err
	}

	r, err := t.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	t.ReadWriteCloser = NewReadWriteCloser(r, w)
	return t.cmd.Start()
}

// Dial creates a new NETCONF session via Junos local shell
// NETCONF interface (xml-mode netconf need-trailer).
func lowlevelDial() (*Session, error) {
	var t TransportJunos
	err := t.Open()
	if err != nil {
		return nil, err
	}

	session, err := NewSession(&t)
	if err != nil {
		return nil, err
	}

	return session, nil
}
