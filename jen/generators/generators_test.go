package generators

import (
    "testing"
)

func TestRandomString(t *testing.T){
    var s = RandomString(10)

    if len(s) != 10{
        t.Error("Problem! this should be 10")
    }

    s = RandomString(-1)
    if len(s) != 0{
        t.Error( "Needs to be 0 length but geting " , len(s))
    }
}

func TestGenerateString(t *testing.T) {
    var v map[string]string
    var g = NewGenerator()
    v = g.GenerateString( "somekey", 10)
    if len(v["somekey"]) != 10 {
        t.Error( "Expected 10, got ", len(v["somekey"]))
    }
}
