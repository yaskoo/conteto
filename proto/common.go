package proto

const (
	Nil FrameType = iota
	Begin
	End
	Ok
	Error
	BeginMeta
	EndMeta
	Meta
	BeginData
	Data
	EndData
)

type FrameType int64

type Frame struct {
	Type   FrameType
	Length int64
	Data   []byte
}

func (f FrameType) String() string {
	switch f {
	case Begin:
		return "Begin"
	case End:
		return "End"
	case Ok:
		return "Ok"
	case Error:
		return "Error"
	case BeginMeta:
		return "BeginMeta"
	case EndMeta:
		return "EndMeta"
	case Meta:
		return "Meta"
	case BeginData:
		return "BeginData"
	case EndData:
		return "EndData"
	case Data:
		return "Data"
	case Nil:
		return "Nil"
	}
	return "some shit frame"
}
