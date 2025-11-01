package main

import "unsafe"

type COWBuffer struct {
	data []byte
	refs *int
}

// создать буффер с определенными данными
func NewCOWBuffer(data []byte) COWBuffer {
	refsCount := 1
	return COWBuffer{
		data: data,
		refs: &refsCount,
	}
}

// создать новую копию буфера
func (b *COWBuffer) Clone() COWBuffer {
	if b.refs != nil {
		*b.refs++
	}
	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

// перестать использовать копию буффера
func (b *COWBuffer) Close() {
	if b.refs == nil {
		return
	}

	*b.refs--
	if *b.refs == 0 {
		b.data = nil
		b.refs = nil
	}
}

// изменить определенный байт в буффере
func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}

	if b.refs != nil && *b.refs > 1 {
		// need to create a new copy of the data
		*b.refs--
		newData := append([]byte(nil), b.data...)
		b.data = newData
		refCount := 1
		b.refs = &refCount
	}

	b.data[index] = value
	return true
}

// сконвертировать буффер в строку
func (b *COWBuffer) String() string {
	return *(*string)(unsafe.Pointer(&b.data))
}
