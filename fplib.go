package fplib
import (
    "bufio"
    "io"
    "io/ioutil"
    "strings"
)
func WithOpenFile(cb func (f *os.File,err error), fname string ) {
    f,err := os.Open(fname)
    cb(f,err)
}
func WithFileAsString(cb func (contents string, err error), fname string) {
    con,err := ioutil.ReadFile(fname)
    cb(con,err)
}
func WithFileAsSlice(cb fun (lines []string,err error),fname string, sep string) {
    fun := func (contents string,err error) {
        lines := strings.Split(contents,sep)
        cb(lines,err)
    }
    WithFileAsString(fun,fname)
}