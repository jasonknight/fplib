package fplib
import (
    "bufio"
    "io"
)
func WithOpenFile(cb func (f *os.File,err error), fname string ) {
    f,err := os.Open(fname)
    cb(f,err)
}