package convert

import (
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

/********************************************************************
created:    2021-01-06
author:     lixianmin

this file is extracted from go-redis/v8/internal/util.go
*********************************************************************/

func AppendArg(b []byte, v interface{}) []byte {
	switch v := v.(type) {
	case nil:
		return append(b, "<nil>"...)
	case string:
		return appendUTF8String(b, v)
	case []byte:
		return appendUTF8String(b, String(v))
	case int:
		return strconv.AppendInt(b, int64(v), 10)
	case int8:
		return strconv.AppendInt(b, int64(v), 10)
	case int16:
		return strconv.AppendInt(b, int64(v), 10)
	case int32:
		return strconv.AppendInt(b, int64(v), 10)
	case int64:
		return strconv.AppendInt(b, v, 10)
	case uint:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint8:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint16:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint32:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint64:
		return strconv.AppendUint(b, v, 10)
	case float32:
		return strconv.AppendFloat(b, float64(v), 'f', -1, 64)
	case float64:
		return strconv.AppendFloat(b, v, 'f', -1, 64)
	case bool:
		if v {
			return append(b, "true"...)
		}
		return append(b, "false"...)
	case time.Time:
		return v.AppendFormat(b, time.RFC3339Nano)
	default:
		return append(b, fmt.Sprint(v)...)
	}
}

func appendUTF8String(b []byte, s string) []byte {
	for _, r := range s {
		b = appendRune(b, r)
	}
	return b
}

func appendRune(b []byte, r rune) []byte {
	if r < utf8.RuneSelf {
		switch c := byte(r); c {
		case '\n':
			return append(b, "\\n"...)
		case '\r':
			return append(b, "\\r"...)
		default:
			return append(b, c)
		}
	}

	l := len(b)
	b = append(b, make([]byte, utf8.UTFMax)...)
	n := utf8.EncodeRune(b[l:l+utf8.UTFMax], r)
	b = b[:l+n]

	return b
}
