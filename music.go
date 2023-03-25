package main

import (
	"github.com/k-evstigneev/go-music/octave"
	"github.com/k-evstigneev/go-music/scales"
)

func main() {
	octaveCircle := octave.CreateOctaveLoop()

	// Print all semitones
	octave.Print(octaveCircle, "All semitones")

	//octave.Print(octave.CreateScaleLoop("C", scales.DiatonicMajorNatural), "C major")
	//octave.Print(octave.CreateScaleLoop("A", scales.DiatonicMinorNatural), "A minor")
	octave.Print(octave.CreateScaleLoop("C", scales.DiatonicMinorNatural), "C minor")
	//octave.Print(octave.CreateScaleLoop("C", scales.DiatonicMajorHarmonic), "C major harmonic")
	//octave.Print(octave.CreateScaleLoop("C", scales.DiatonicMinorHarmonic), "C minor harmonic")
}
