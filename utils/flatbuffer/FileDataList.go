// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuffer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FileDataList struct {
	_tab flatbuffers.Table
}

func GetRootAsFileDataList(buf []byte, offset flatbuffers.UOffsetT) *FileDataList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FileDataList{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FileDataList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FileDataList) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FileDataList) Files(obj *FileData, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *FileDataList) FilesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func FileDataListStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func FileDataListAddFiles(builder *flatbuffers.Builder, files flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(files), 0)
}
func FileDataListStartFilesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FileDataListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
