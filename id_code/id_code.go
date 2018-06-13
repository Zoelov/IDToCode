package id_code

import (
	"container/list"
	"errors"
)

var baseCodeStr = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"

// GenCodeBase34 -- 将id转换成6位长度的code
func GenCodeBase34(id uint64) []byte {
	num := id
	mod := uint64(0)
	l := list.New()

	baseCodeByte := []byte(baseCodeStr)

	for num != 0 {
		mod = num % 34
		num = num / 34

		l.PushFront(baseCodeByte[int(mod)])
	}

	listLen := l.Len()

	var res []byte
	if listLen >= 6 {
		res = make([]byte, 0, listLen)
		for i := l.Front(); i != nil; i = i.Next() {
			res = append(res, i.Value.(byte))
		}
	} else {
		res = make([]byte, 0, 6)
		for i := 0; i < 6; i++ {
			if i < 6-listLen {
				res = append(res, baseCodeByte[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}
		}
	}

	return res
}

// CodeToIDBase34 -- 将code逆向转换成原始id
func CodeToIDBase34(idByte []byte) (uint64, error) {
	baseCodeByte := []byte(baseCodeStr)
	baseMap := make(map[byte]int)
	for i, v := range baseCodeByte {
		baseMap[v] = i
	}

	if idByte == nil || len(idByte) == 0 {
		return 0, errors.New("param id nil or empyt")
	}

	var res uint64
	var r uint64

	for i := len(idByte) - 1; i >= 0; i-- {
		v, ok := baseMap[idByte[i]]
		if !ok {
			return 0, errors.New("param contain illegle character")
		}

		var b uint64 = 1
		for j := uint64(0); j < r; j++ {
			b *= 34
		}

		res += b * uint64(v)
		r++
	}

	return res, nil
}
