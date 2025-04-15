// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main_test

import (
	"os/exec"
	"testing"
)

func TestEcho(t *testing.T) {
	setup(t)
	t.Fatal("NOT COMPLETE YET")
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
