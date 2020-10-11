package prototype

type ClonePrototyper interface {
	/*Print(string)*/
	Clone() ClonePrototyper
}
