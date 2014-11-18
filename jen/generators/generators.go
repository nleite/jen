package generators

import (
    "math/rand"
    "fmt"
    "bytes"
    "time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func NewGenerator() *Generator{
    rand.Seed( time.Now().UTC().UnixNano())
    return &Generator{alphabet}
}

type Generator struct{
    alphabet string
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

func (g Generator) GenerateString(k string, size int) map[string] string{
    var v = RandomString(size)
    fmt.Println( v )
    m := make(map[string]string)
    m[k] = v
    return m
}
