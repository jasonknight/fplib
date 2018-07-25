package fplib

func MapIntsToInts(cb func (e int) int,arr []int) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapIntsToUints(cb func (e int) uint,arr []int) []uint {
    ret := make([]uint,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapIntsToBytes(cb func (e int) byte,arr []int) []byte {
    ret := make([]byte,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapIntsToRunes(cb func (e int) rune,arr []int) []rune {
    ret := make([]rune,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapIntsToStrings(cb func (e int) string,arr []int) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapIntsToFloat32s(cb func (e int) float32,arr []int) []float32 {
    ret := make([]float32,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapIntsToFloat64s(cb func (e int) float64,arr []int) []float64 {
    ret := make([]float64,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapIntsToBools(cb func (e int) bool,arr []int) []bool {
    ret := make([]bool,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapUintsToInts(cb func (e uint) int,arr []uint) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapUintsToUints(cb func (e uint) uint,arr []uint) []uint {
    ret := make([]uint,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapUintsToBytes(cb func (e uint) byte,arr []uint) []byte {
    ret := make([]byte,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapUintsToRunes(cb func (e uint) rune,arr []uint) []rune {
    ret := make([]rune,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapUintsToStrings(cb func (e uint) string,arr []uint) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapUintsToFloat32s(cb func (e uint) float32,arr []uint) []float32 {
    ret := make([]float32,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapUintsToFloat64s(cb func (e uint) float64,arr []uint) []float64 {
    ret := make([]float64,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBytesToInts(cb func (e byte) int,arr []byte) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBytesToUints(cb func (e byte) uint,arr []byte) []uint {
    ret := make([]uint,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBytesToBytes(cb func (e byte) byte,arr []byte) []byte {
    ret := make([]byte,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBytesToRunes(cb func (e byte) rune,arr []byte) []rune {
    ret := make([]rune,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBytesToStrings(cb func (e byte) string,arr []byte) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBytesToFloat32s(cb func (e byte) float32,arr []byte) []float32 {
    ret := make([]float32,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBytesToFloat64s(cb func (e byte) float64,arr []byte) []float64 {
    ret := make([]float64,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapRunesToInts(cb func (e rune) int,arr []rune) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapRunesToUints(cb func (e rune) uint,arr []rune) []uint {
    ret := make([]uint,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapRunesToBytes(cb func (e rune) byte,arr []rune) []byte {
    ret := make([]byte,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapRunesToRunes(cb func (e rune) rune,arr []rune) []rune {
    ret := make([]rune,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapRunesToStrings(cb func (e rune) string,arr []rune) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapRunesToFloat32s(cb func (e rune) float32,arr []rune) []float32 {
    ret := make([]float32,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapRunesToFloat64s(cb func (e rune) float64,arr []rune) []float64 {
    ret := make([]float64,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToInts(cb func (e string) int,arr []string) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToUints(cb func (e string) uint,arr []string) []uint {
    ret := make([]uint,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToBytes(cb func (e string) byte,arr []string) []byte {
    ret := make([]byte,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToRunes(cb func (e string) rune,arr []string) []rune {
    ret := make([]rune,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToStrings(cb func (e string) string,arr []string) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToFloat32s(cb func (e string) float32,arr []string) []float32 {
    ret := make([]float32,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToFloat64s(cb func (e string) float64,arr []string) []float64 {
    ret := make([]float64,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapStringsToBools(cb func (e string) bool,arr []string) []bool {
    ret := make([]bool,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat32sToInts(cb func (e float32) int,arr []float32) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat32sToUints(cb func (e float32) uint,arr []float32) []uint {
    ret := make([]uint,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat32sToBytes(cb func (e float32) byte,arr []float32) []byte {
    ret := make([]byte,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat32sToRunes(cb func (e float32) rune,arr []float32) []rune {
    ret := make([]rune,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat32sToStrings(cb func (e float32) string,arr []float32) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat32sToFloat32s(cb func (e float32) float32,arr []float32) []float32 {
    ret := make([]float32,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat32sToFloat64s(cb func (e float32) float64,arr []float32) []float64 {
    ret := make([]float64,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat64sToInts(cb func (e float64) int,arr []float64) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat64sToUints(cb func (e float64) uint,arr []float64) []uint {
    ret := make([]uint,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat64sToBytes(cb func (e float64) byte,arr []float64) []byte {
    ret := make([]byte,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat64sToRunes(cb func (e float64) rune,arr []float64) []rune {
    ret := make([]rune,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat64sToStrings(cb func (e float64) string,arr []float64) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat64sToFloat32s(cb func (e float64) float32,arr []float64) []float32 {
    ret := make([]float32,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapFloat64sToFloat64s(cb func (e float64) float64,arr []float64) []float64 {
    ret := make([]float64,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBoolsToInts(cb func (e bool) int,arr []bool) []int {
    ret := make([]int,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBoolsToStrings(cb func (e bool) string,arr []bool) []string {
    ret := make([]string,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}


func MapBoolsToBools(cb func (e bool) bool,arr []bool) []bool {
    ret := make([]bool,len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}
