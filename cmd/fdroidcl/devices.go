// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"fmt"
	"log"

	"github.com/mvdan/fdroidcl/adb"
)

var cmdDevices = &Command{
	UsageLine: "devices",
	Short:     "List connected devices",
}

func init() {
	cmdDevices.Run = runDevices
}

func runDevices(args []string) {
	startAdbIfNeeded()
	devices, err := adb.Devices()
	if err != nil {
		log.Fatalf("Could not get devices: %v", err)
	}
	for _, device := range devices {
		fmt.Printf("%s - %s (%s)\n", device.ID, device.Model, device.Product)
	}
}

func startAdbIfNeeded() {
	if adb.IsServerRunning() {
		return
	}
	log.Printf("Starting ADB server...")
	if err := adb.StartServer(); err != nil {
		log.Fatalf("Could not start ADB server: %v", err)
	}
}

func mustOneDevice() *adb.Device {
	startAdbIfNeeded()
	devices, err := adb.Devices()
	if err != nil {
		log.Fatalf("Could not get devices: %v", err)
	}
	if len(devices) != 1 {
		log.Fatalf("Exactly one connected device is needed")
	}
	return devices[0]
}
