// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main_test

import (
	"bytes"
	"net"
	"os/exec"
	"testing"
)

func TestEcho(t *testing.T) {
	t.SkipNow()
	setup(t)

	t.Run("should be connectable via tcp4", func(t *testing.T) {
		conn, err := net.Dial("tcp4", ":6767")
		if err != nil {
			t.Error(err);
		} else {
			defer conn.Close()
		}
	})

	t.Run("should accept data", func(t *testing.T) {
		conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
			Port: 6767,
		})
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		data := []byte("sample data")
		_, err = conn.Write(data)

		if err != nil {
			t.Errorf("could not write data")
		}
	})

	t.Run("should write back", func(t *testing.T) {
		conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
			Port: 6767,
		})
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		data := []byte("sample data")
		conn.Write(data)

		got := make([]byte, 0, len(data))
		conn.Read(got)
		want := data

		if !bytes.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func setup(tb testing.TB) {
	tb.Helper()
	compile(tb)
	runServer(tb)
}

func compile(tb testing.TB) {
	tb.Helper()
	err := exec.Command("go", "build", "-o", "echod").Run()
	if err != nil {
		tb.Fatal("failed to build echod")
	}
	tb.Cleanup(func() {
		_ = exec.Command("go", "clean").Run()
	})
}

func runServer(tb testing.TB) {
	tb.Helper()
	cmd := exec.Command("./echod", "127.0.0.1:2929")
	if err := cmd.Start(); err != nil {
		tb.Fatal("cannot start echod:", err)
	}
	tb.Cleanup(func() {
		if err := cmd.Process.Kill(); err != nil {
			tb.Logf("cannot kill echod: %s", err)
		}
	})
}
