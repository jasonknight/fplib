package fplib

func EveryInt(cb func (e int) bool, ints []int) bool {
    for i := range ints {
        if !cb(ints[i]) {
            return false
        }
    }
    return true
}


func EveryUint(cb func (e uint) bool, uints []uint) bool {
    for i := range uints {
        if !cb(uints[i]) {
            return false
        }
    }
    return true
}


func EveryByte(cb func (e byte) bool, bytes []byte) bool {
    for i := range bytes {
        if !cb(bytes[i]) {
            return false
        }
    }
    return true
}


func EveryRune(cb func (e rune) bool, runes []rune) bool {
    for i := range runes {
        if !cb(runes[i]) {
            return false
        }
    }
    return true
}


func EveryString(cb func (e string) bool, strings []string) bool {
    for i := range strings {
        if !cb(strings[i]) {
            return false
        }
    }
    return true
}


func EveryFloat32(cb func (e float32) bool, float32s []float32) bool {
    for i := range float32s {
        if !cb(float32s[i]) {
            return false
        }
    }
    return true
}


func EveryFloat64(cb func (e float64) bool, float64s []float64) bool {
    for i := range float64s {
        if !cb(float64s[i]) {
            return false
        }
    }
    return true
}


func EveryBool(cb func (e bool) bool, bools []bool) bool {
    for i := range bools {
        if !cb(bools[i]) {
            return false
        }
    }
    return true
}
