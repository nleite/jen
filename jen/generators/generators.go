package generators

import (
    "math/rand"
//    "fmt"
    "bytes"
    "time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func NewGenerator() *Generator{
    rand.Seed( time.Now().UTC().UnixNano())
    //max integer value by default will 1024 cause I'm not very creative today
    return &Generator{alphabet, 1024}
}

type Generator struct{
    alphabet string
    intsize int
}

func RandomString(size int) string{
    var buf bytes.Buffer
    var tmp string
    for i :=0; i < size; i++{
        tmp = string( alphabet[ rand.Intn(len(alphabet))] )
        buf.WriteString(tmp)
    }
    return buf.String()
}


func RandomInt(n int) int{
    return rand.Intn(n)
}

func (g Generator) GenerateInteger(k string) map[string]int {
    m := make(map[string]int)
    m[k] = RandomInt(g.intsize)
    return m
}

func (g Generator) GenerateString(k string, size int) map[string] string{
    var v = RandomString(size)
    m := make(map[string]string)
    m[k] = v
    return m
}
