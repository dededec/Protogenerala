package main

type Serializer interface {
	Serialize() []byte
	Deserialize(b []byte) Serializer
}
