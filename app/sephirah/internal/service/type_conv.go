package service

import librarian "github.com/tuihub/protos/pkg/librarian/v1"

func toBizInternalID(i *librarian.InternalID) int64 {
	return i.GetId()
}

func toBizInternalIDList(idl []*librarian.InternalID) []int64 {
	res := make([]int64, len(idl))
	for i, id := range idl {
		res[i] = toBizInternalID(id)
	}
	return res
}
