// Package savestates enables emulator state manipulation.
package nanoarch

import (
	"io/ioutil"
)

type state []byte

// Saves the current state to the filesystem.
// Deadlock warning: locks the emulator.
func (na *naEmulator) Save() error {
	na.GetLock()
	defer na.ReleaseLock()

	if state, err := getState(); err == nil {
		return state.toFile(na.GetHashPath())
	} else {
		return err
	}
}

// Load the state from the filesystem.
// Deadlock warning: locks the emulator.
func (na *naEmulator) Load() error {
	na.GetLock()
	defer na.ReleaseLock()

	path := na.GetHashPath()
	if state, err := fromFile(path); err == nil {
		return restoreState(state)
	} else {
		return err
	}
}

// Returns the current emulator state.
func getState() (state, error) {
	if dat, err := serialize(serializeSize()); err == nil {
		return dat, nil
	} else {
		return state{}, err
	}
}

// Restores an emulator state.
func restoreState(dat state) error {
	return unserialize(dat, serializeSize())
}

// Writes the state to a file with the path.
func (st state) toFile(path string) error {
	return ioutil.WriteFile(path, st, 0644)
}

// Reads the state from a file with the path.
func fromFile(path string) (state, error) {
	if bytes, err := ioutil.ReadFile(path); err == nil {
		return bytes, nil
	} else {
		return state{}, err
	}
}
