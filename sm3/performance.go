package sm3

import (
	"hash"
	"sync"
)

// Pool is a sync.Pool for *SM3 instances to reduce allocations
var pool = sync.Pool{
	New: func() interface{} {
		return New()
	},
}

// Get returns a *SM3 from the pool
func Get() hash.Hash {
	return pool.Get().(hash.Hash)
}

// Put returns a *SM3 to the pool
func Put(h hash.Hash) {
	h.Reset()
	pool.Put(h)
}

// Sum returns the SM3 checksum of the data without allocating a new hasher
func Sum(data []byte) [32]byte {
	h := Get()
	defer Put(h)
	
	h.Write(data)
	var result [32]byte
	copy(result[:], h.Sum(nil))
	return result
}

// NewWriter returns a writer that computes the SM3 checksum of written data
func NewWriter() *Writer {
	return &Writer{h: Get()}
}

// Writer is a writer that computes the SM3 checksum of written data
// and can be reset for reuse
//
// Example:
//
//  w := sm3.NewWriter()
//  defer w.Close()
//  io.Copy(w, data)
//  sum := w.Sum(nil)
//  w.Reset()
//  io.Copy(w, moreData)
//  sum2 := w.Sum(nil)
type Writer struct {
	h hash.Hash
}

func (w *Writer) Write(p []byte) (int, error) {
	return w.h.Write(p)
}

func (w *Writer) Sum(b []byte) []byte {
	return w.h.Sum(b)
}

func (w *Writer) Reset() {
	w.h.Reset()
}

func (w *Writer) Close() {
	if w.h != nil {
		Put(w.h)
		w.h = nil
	}
}