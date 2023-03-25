// Package octave supposed to represent the sound system basic structures:
// octave divided into 12 semitones, note structure and few utility methods
// to navigate around octave
package octave

import "fmt"

type Note struct {
	Value string
	Next  *Note
}

// LoopedSequence can be replaced with ring
type LoopedSequence struct {
	root *Note
}

// var allNoteLabels = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var allNoteLabels = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var allNotesLoopedSequence = CreateOctaveLoop()

func CreateOctaveLoop() *LoopedSequence {
	return createLoopedSequenceFromNoteLabels(allNoteLabels)
}

func CreateScaleLoop(root string, scaleIntervals []int) *LoopedSequence {
	return createScale(root, scaleIntervals)
}

func Print(loopedSequence *LoopedSequence, description string) {
	var cursor = loopedSequence.root
	for i := 0; i == 0 || cursor != loopedSequence.root; i++ {
		if i == 0 {
			fmt.Printf("\n%v: %v", description, cursor.Value)
		} else {
			fmt.Printf(" - %v", cursor.Value)
		}
		cursor = cursor.Next
	}
	//println()
}

func createLoopedSequenceFromNoteLabels(noteLabels []string) *LoopedSequence {
	var pFirst *Note
	var pLast *Note

	for _, noteValue := range noteLabels {
		if pFirst == nil {
			pFirst = &Note{Value: noteValue}
			pLast = pFirst
			continue
		}

		pNew := &Note{Value: noteValue}
		(*pLast).Next = pNew
		pLast = pNew
	}

	(*pLast).Next = pFirst

	return &LoopedSequence{root: pFirst}
}

func createScale(root string, scaleIntervals []int) *LoopedSequence {
	var cursor = findNote(root, allNotesLoopedSequence)
	var scaleNoteLabels = make([]string, len(scaleIntervals)+1)

	scaleNoteLabels[0] = cursor.Value
	var currentIndex = 1
	for _, interval := range scaleIntervals {
		cursor = iterate(cursor, interval)
		scaleNoteLabels[currentIndex] = cursor.Value
		currentIndex++
	}
	return createLoopedSequenceFromNoteLabels(scaleNoteLabels)
}

func iterate(cursor *Note, times int) *Note {
	for i := 0; i < times; i++ {
		cursor = cursor.Next
	}
	return cursor
}

func findNote(note string, octaveCircle *LoopedSequence) *Note {
	var cursor = octaveCircle.root
	for {
		if cursor.Value == note {
			break
		}
		cursor = cursor.Next
	}
	return cursor
}
