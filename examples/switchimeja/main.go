package main

import (
	"fmt"
	"os"

	"github.com/hnakamur/moderniejapanizer"
	"github.com/hnakamur/w32version"
	"github.com/mattn/go-ole"
)

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	version, err := w32version.GetVersion()
	if err != nil {
		panic(err)
	}

	err = moderniejapanizer.SwitchInputMethodJa(version)
	if err != nil {
		fmt.Printf("Error while setting display language: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Switched language to Japanese.")
}
