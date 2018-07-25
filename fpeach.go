package fplib

func EachInt(cb func (value int, ind int), ints []int) {
    for i := range ints {
        cb(ints[i],i)
    }
}


func EachUint(cb func (value uint, ind int), uints []uint) {
    for i := range uints {
        cb(uints[i],i)
    }
}


func EachByte(cb func (value byte, ind int), bytes []byte) {
    for i := range bytes {
        cb(bytes[i],i)
    }
}


func EachRune(cb func (value rune, ind int), runes []rune) {
    for i := range runes {
        cb(runes[i],i)
    }
}


func EachString(cb func (value string, ind int), strings []string) {
    for i := range strings {
        cb(strings[i],i)
    }
}


func EachFloat32(cb func (value float32, ind int), float32s []float32) {
    for i := range float32s {
        cb(float32s[i],i)
    }
}


func EachFloat64(cb func (value float64, ind int), float64s []float64) {
    for i := range float64s {
        cb(float64s[i],i)
    }
}


func EachBool(cb func (value bool, ind int), bools []bool) {
    for i := range bools {
        cb(bools[i],i)
    }
}
