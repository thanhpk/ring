package ring

// Ring is a circular array
// Ring starts from a small circle, when needed, can expand up to its full size,
// and then when appropriated it shrink smaller to save memory.
// Note: this type not threadsafe
type Ring struct {
	// holds all elements of the ring
	elements []interface{}

	// current number of used element in me.elements
	size int

	// index of the head in me.elements slice
	head int
}

// NewRing creates a new Ring object
func NewRing() Ring { return Ring{} }

// rotate shift the ring to the right n elements
// shift the right to the left -n elements if n is negative
// this function modify the given arr
// current implementation has complexity of O(n)
func rotate(arr []interface{}, n int) {
	lenArr := len(arr)
	if lenArr == 0 {
		return
	}

	if n == 0 {
		return
	}

	n = (n%lenArr + lenArr) % lenArr // make sure n alway greater than 0

	swapped := 0 // number of swap has taken
	for i := 0; i < n; i++ {
		if swapped == lenArr {
			break // quick done
		}
		place := arr[i]
		for j := i; (j-n+lenArr)%lenArr != i; j -= n {
			swapped++
			j = (j + 2*lenArr) % lenArr
			arr[j] = arr[(j-n+lenArr)%lenArr]
		}
		arr[(i+n)%lenArr] = place
		swapped++
	}
}

// Size returns number of elements in the ring
func (me Ring) Size() int { return me.size }

// Cap returns capacity of the slice holding elements
func (me *Ring) Cap() int { return cap(me.elements) }

// Append adds a new element to the end of the ring
func (me *Ring) Append(value interface{}) {
	if me.size == len(me.elements) { // need to expand
		rotate(me.elements, -me.head)
		me.head = 0

		// double me.elements capacity
		me.elements = append(me.elements, nil)
		me.elements = me.elements[:cap(me.elements)]
	}
	me.elements[(me.head+me.size)%len(me.elements)] = value
	me.size++
}

// Drop removes the first element
// do nothing if there is nothing to remove
func (me *Ring) Drop() {
	if me.size == 0 {
		me.head = 0
		return
	}
	me.head = (me.head + 1) % len(me.elements)
	me.size--

	// shrink the slice if appropriate
	halflength := len(me.elements) / 2
	if me.size < halflength {
		rotate(me.elements, -me.head)
		me.head = 0
		me.elements = append([]interface{}{}, me.elements[:halflength]...)
	}
}

// Each calls function f on every element of the ring, in forward order.
func (me *Ring) Each(f func(interface{})) {
	if me == nil {
		return
	}
	for i := 0; i < me.size; i++ {
		f(me.elements[(me.head+i)%len(me.elements)])
	}
}
