package transformation

import "strconv"

type StrTo string

// 转字符串
func (s StrTo) String() string {
	return string(s)
}

// Int string 转 int 返回 int 和 error
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt string 转 int 仅返回 int
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// Int64 string 转 int64 返回 int64 和 error
func (s StrTo) Int64() (int64,error) {
	v,err := strconv.ParseInt(s.String(),16,32)
	return v,err
}

// MustInt64 string 转 int64 仅返回 int64
func (s StrTo) MustInt64() int64 {
	v,_ := s.Int64()
	return v
}

// Uint32 string 转 uint32 返回 uint32 和 error
func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// MustUint32 string 转 uint32 仅返回 uint32
func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}

// Uint64 string 转 uint64 返回 uint64 和 error
func (s StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	return uint64(v), err
}

// MustUint64 string 转 uint64 仅返回 uint64
func (s StrTo) MustUint64() uint64 {
	v, _ := s.Uint64()
	return v
}
