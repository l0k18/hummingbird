// Package alloc - simple implementation of a bit array for tracking used and unused elements in an array
package alloc

import "errors"

// Alloc - data structure for tracking how many nodes exist in the Dtree and enabling them to be deleted for decomposition
type Alloc struct {
	Count uint32
	List  []byte
}

// NewAlloc - creates a new Alloc structure for tracking Dtree node allocations
func NewAlloc() Alloc {
	var list Alloc
	list.List = make([]byte, 1)
	list.New()
	list.New()
	return list
}

// New - Finds the first free slot in the allocation list and marks it. Grows list if it has no free slots
func (a *Alloc) New() (uint32, bool) {
	i, j := uint32(0), uint32(0)
	length := uint32(len(a.List))
	allocated := a.List[i]&(1<<j) != 0
	grew := false
	for i < length && allocated {
		if ^a.List[i] != 0 {
			j++
			allocated = a.List[i]&(1<<j) != 0
			if j > 7 {
				j = 0
				i++
			}
		} else {
			i++
		}
	}
	if j > 7 || i >= uint32(len(a.List)) {
		a.List = append(a.List, 1)
		grew = true
	} else {
		a.List[i] = a.List[i] ^ (1 << j)
	}
	a.Count++
	return i*8 + j, grew
}

// Free - Removes an allocation for a slot in the bitlist of tree node array
func (a *Alloc) Free(slot uint32) error {
	length := uint32(len(a.List))
	remainder := slot % 8
	bytepos := slot / 8
	if bytepos > length {
		return errors.New("Requested slot exceeds allocation list length, invalid request")
	}
	if a.List[bytepos]&1<<remainder == 0 {
		return errors.New("Requested slot is already free")
	}
	a.List[bytepos] ^= 1 << remainder
	a.Count--
	return nil
}
