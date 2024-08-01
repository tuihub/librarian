package libtype

import "encoding/json"

func DeepCopyStruct(src, dst any) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}
