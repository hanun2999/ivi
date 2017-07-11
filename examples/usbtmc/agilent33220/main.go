// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"log"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen/agilent33220"
	"github.com/gotmc/usbtmc"
	_ "github.com/gotmc/usbtmc/driver/truveris"
)

func main() {

	// Create a USBTMC context and set the debug level
	ctx, err := usbtmc.NewContext()
	if err != nil {
		log.Fatalf("Error creating new USB context: %s", err)
	}
	defer ctx.Close()
	ctx.SetDebugLevel(1)

	// Create a new USBTMC device
	dev, err := ctx.NewDevice("USB0::2391::1031::MY44035849::INSTR")
	if err != nil {
		log.Fatalf("NewDevice error: %s", err)
	}
	defer dev.Close()

	// Create a new IVI instance of the Agilent 33220 function generator
	fgen, err := agilent33220.New(dev)
	if err != nil {
		log.Fatalf("IVI instrument error: %s", err)
	}

	// Channel specific methods can be accessed directly from the instrument
	// using 0-based index to select the desirec channel.
	fgen.Channels[0].DisableOutput()
	fgen.Channels[0].SetAmplitude(0.4)

	// Alternatively, the channel can be assigned to a variable.
	ch := fgen.Channels[0]
	ch.SetStandardWaveform(ivi.Sine)
	ch.SetDCOffset(0.1)
	ch.SetFrequency(2340)

	// Instead of configuring attributes of a standard waveform individually, the
	// standard waveform can be configured using a single method.
	ch.ConfigureStandardWaveform(ivi.RampUp, 0.4, 0.1, 2340, 0)
	ch.EnableOutput()

	// Query the FGen
	freq, err := ch.Frequency()
	if err != nil {
		log.Printf("error querying frequency: %s", err)
	}
	log.Printf("Frequency = %.0f Hz", freq)
	amp, err := ch.Amplitude()
	if err != nil {
		log.Printf("error querying amplitude: %s", err)
	}
	log.Printf("Amplitude = %.3f Vpp", amp)
	wave, err := ch.StandardWaveform()
	if err != nil {
		log.Printf("error querying standard waveform: %s", err)
	}
	log.Printf("Standard waveform = %s", wave)
}
