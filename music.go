package main

import (
	"fmt"
)

type (
	Note struct {
		value string
		next  *Note
	}
)

func OctaveCircle() *Note {
	notes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	var pFirst *Note
	var pLast *Note

	for _, noteValue := range notes {
		if pFirst == nil {
			pFirst = &Note{value: noteValue}
			pLast = pFirst
			continue
		}

		pNew := &Note{value: noteValue}
		(*pLast).next = pNew
		pLast = pNew
	}

	(*pLast).next = pFirst

	return pFirst
}

func findRootNote(rootNote string, octaveCircle *Note) *Note {
	var cursor *Note = octaveCircle
	for {
		if cursor.value == rootNote {
			break
		}
		cursor = cursor.next
	}
	return cursor
}

func printLinkedList(head *Note, limit int) {
	var cursor *Note = head
	fmt.Println()
	for i := 0; i <= limit || head == nil; i++ {
		if i != 0 {
			fmt.Print(" - ")
		}
		fmt.Print(cursor.value)
		cursor = cursor.next
	}
}

func printScale(octaveCircle *Note, root string, scaleIntervals []int, description string) {
	var cursor *Note = findRootNote(root, octaveCircle)
	fmt.Printf("\n%v: %v", description, root)
	for _, interval := range scaleIntervals {
		switch interval {
		case 3:
			cursor = cursor.next.next.next
		case 2:
			cursor = cursor.next.next
		case 1:
			cursor = cursor.next
		}
		fmt.Printf(" - %v", cursor.value)
	}
}

// W, W, H, W, W, W, H
var DIATONIC_MAJOR_NATURAL = []int{2, 2, 1, 2, 2, 2, 1}

// "W", "W", "H", "W", "H", "WH", "H"
var DIATONIC_MAJOR_HARMONIC = []int{2, 2, 1, 2, 1, 3, 1}

// "W", "H", "W", "W", "H", "W", "W"
var DIATONIC_MINOR_NATURAL = []int{2, 1, 2, 2, 1, 2, 2}

// "W", "H", "W", "W", "H", "WH", "H"
var DIATONIC_MINOR_HARMONIC = []int{2, 1, 2, 2, 1, 3, 1}

func main() {
	octaveCircle := OctaveCircle()
	printLinkedList(octaveCircle, 20)
	printScale(octaveCircle, "C", DIATONIC_MAJOR_NATURAL, "C Major Natural")
	printScale(octaveCircle, "C", DIATONIC_MAJOR_HARMONIC, "C Major Harmonic")
	printScale(octaveCircle, "A", DIATONIC_MINOR_NATURAL, "A Minor Natural")
	printScale(octaveCircle, "A", DIATONIC_MINOR_HARMONIC, "A Minor Harmonic")
}
