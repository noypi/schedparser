package gen

type OnSchedParseCb func(*SchedParseRes)

type SchedLex struct {
	S            string
	pos          int
	OnSchedParse OnSchedParseCb
	Itemchan     chan *SchedParseRes
	buf          []byte
}
