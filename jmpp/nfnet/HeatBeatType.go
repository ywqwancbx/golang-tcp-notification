package nfnet

type BaseHB struct {
	ST int //milliseconds 
}

type HS01 struct {
	H string
	BaseHB
}

type HC01 struct {
	H string
	BaseHB
}
