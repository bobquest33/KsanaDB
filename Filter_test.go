package KsanaDB 
import (
    "testing"
)

func Test_filter(t *testing.T) {
    Hitfilter := []int64{5,2,1}
    Missfilter := []int64{0,7}
    tags := []string{"0", "1", "2", "3", "4", "5"}
    smallTags := []string{"0"}

    retT := filter(Hitfilter, tags)

    if retT == false {
        t.Error(Hitfilter) 
    }    

    retF := filter(Missfilter, tags)
    if retF == true {
        t.Error(Missfilter) 
    }    
    retF = filter(Missfilter, smallTags)
    if retF == true {
        t.Error(Missfilter) 
    }    

}