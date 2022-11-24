package models

type Token struct {
	Type  Type
	value string
}

type Type int

const (
	sepBeg Type = iota

	CurvedO Type = iota - 1
	CurvedC
	RollO
	RollC
	Assign

	sepEnd
)

const (
	notSepBeg = iota + sepEnd

	NumFloat = iota + sepEnd - 1
	NumInt
	String
	BeginData
	Bool
	Field

	notSepEnd

	Unknown
)

func (t *Token) Clone() *Token {
	newToken := *t
	return &newToken
}
