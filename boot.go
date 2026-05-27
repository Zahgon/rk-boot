// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

// Package rkboot is bootstrapper for rk style application
package rkboot

import (
	"context"
	"embed"

	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

type hookFuncM map[string]map[string]func(ctx context.Context)

func newHookFuncM() hookFuncM { _ = "STUB: not implemented"; return *new(hookFuncM) }

func (m hookFuncM) addFunc(entryType, entryName string, f func(ctx context.Context)) {
	_ = "STUB: not implemented"
	return
}

func (m hookFuncM) getFunc(entryType, entryName string) func(ctx context.Context) {
	_ = "STUB: not implemented"
	return nil
}

// Boot is a structure for bootstrapping rk style application
type Boot struct {
	bootConfigPath string    `yaml:"-" json:"-"`
	embedFS        *embed.FS `yaml:"-" json:"-"`
	bootConfigRaw  []byte    `yaml:"-" json:"-"`
	beforeHookF    hookFuncM `yaml:"-" json:"-"`
	afterHookF     hookFuncM `yaml:"-" json:"-"`
	EventId        string    `yaml:"-" json:"-"`
	pluginEntries  map[string]map[string]rkentry.Entry
	userEntries    map[string]map[string]rkentry.Entry
	webEntries     map[string]map[string]rkentry.Entry
}

// BootOption is used as options while bootstrapping from code
type BootOption func(*Boot)

// WithBootConfigPath provide boot config yaml file.
func WithBootConfigPath(filePath string, fs *embed.FS) BootOption {
	_ = "STUB: not implemented"
	return *new(BootOption)
}

// WithBootConfigRaw provide boot config as string.
func WithBootConfigRaw(raw []byte) BootOption { _ = "STUB: not implemented"; return *new(BootOption) }

// NewBoot create a bootstrapper.
func NewBoot(opts ...BootOption) *Boot { _ = "STUB: not implemented"; return nil }

// Register entries need to pre-build.

// AddHookFuncBeforeBootstrap run functions before certain entry Bootstrap()
func (boot *Boot) AddHookFuncBeforeBootstrap(entryType, entryName string, f func(ctx context.Context)) {
	_ = "STUB: not implemented"
	return
}

// AddHookFuncAfterBootstrap run functions before certain entry Bootstrap()
func (boot *Boot) AddHookFuncAfterBootstrap(entryType, entryName string, f func(ctx context.Context)) {
	_ = "STUB: not implemented"
	return
}

// Bootstrap entries as sequence of plugin, user defined and web framework
func (boot *Boot) Bootstrap(ctx context.Context) { _ = "STUB: not implemented"; return }

// WaitForShutdownSig wait for shutdown signal.
// 1: Call shutdown hook function added by user.
// 2: Call interrupt function of entries in rkentry.GlobalAppCtx.
func (boot *Boot) WaitForShutdownSig(ctx context.Context) { _ = "STUB: not implemented"; return }

// Shutdown shutdown boot. for non-web application
// 1: Call shutdown hook function added by user.
// 2: Call interrupt function of entries in rkentry.GlobalAppCtx.
func (boot *Boot) Shutdown(ctx context.Context) {
	_ = "STUB: not implemented"
	// Call shutdown hook function
	return
}

// Call interrupt

// AddShutdownHookFunc add shutdown hook function
func (boot *Boot) AddShutdownHookFunc(name string, f rkentry.ShutdownHook) {
	_ = "STUB: not implemented"
	return
}

// Interrupt entries as sequence of plugin, user defined and web framework
func (boot *Boot) interrupt(ctx context.Context) { _ = "STUB: not implemented"; return }

// Interrupt external entries

// readYAML read YAML file
func (boot *Boot) readYAML() []byte {
	_ = "STUB: not implemented"
	// case 1: if user provide raw then, continue
	return nil
}

// case 2: if embed.FS is not nil, then try to read from it

// case 3: try to read from local, if bootConfigPath is empty, then try to read from default boot.yaml

// sync logs
func syncLog(eventId string) { _ = "STUB: not implemented"; return }
