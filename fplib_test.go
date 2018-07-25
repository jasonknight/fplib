package fplib

import "testing"
import "fmt"
import "strconv"

func TestMapIntsToInts(t *testing.T) {
    ints := []int{1,2,3,4,5}
    new_ints := MapIntsToInts(func (e int) int {
        ei := e
        return ei * ei
    },ints)
    for i := range ints {
        oe := ints[i]
        ne := new_ints[i]
        if ne != (oe * oe) {
            t.Errorf("Expected %d * %d to equal %d\n",oe,oe,ne)
        }
    }
}
func TestMapIntsToStrings(t *testing.T) {
    ints := []int{1,2,3,4,5}
    tstrings := []string{"1","2","3","4","5"}
    new_ints := MapIntsToStrings(func (e int) string {
        return fmt.Sprint(e)
    },ints)
    for i := range ints {
        ne := new_ints[i]
        if ne != tstrings[i] {
            t.Errorf("Expected %s to equal %d\n",tstrings[i],ne)
        }
    }
}
func TestMapStringsToInts(t *testing.T) {
    tints := []int{1,2,3,4,5}
    strings := []string{"1","2","3","4","5"}
    new_ints := MapStringsToInts(func (e string) int {
        i,_ := strconv.Atoi(e)
        return i
    },strings)
    for i := range tints {
        ne := new_ints[i]
        if ne != tints[i] {
            t.Errorf("Expected %s to equal %d\n",tints[i],ne)
        }
    }
}
func TestEveryInt(t *testing.T) {
    good_ints := []int{1,1,1}
    bad_ints := []int{1,2,3}
    fun := func (e int) bool {
        if e == 1 {
            return true
        }
        return false
    }
    if EveryInt(fun,good_ints) != true {
        t.Errorf("expected good_ints to pass")
    }
    if EveryInt(fun,bad_ints) {
        t.Errorf("expected bad_ints to fail")
    }
}
func TestEachInt(t *testing.T) {
    good_ints := []int{1,1,1}
    tval := false
    fun := func (e int,ind int) {
        if e == 1 {
            tval = true
            return
        }
        tval = false
    }
    EachInt(fun,good_ints)
    if tval == false {
        t.Errorf("expected tval to be set to true!")
    }
}