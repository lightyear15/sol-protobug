package main

import (
    "math/rand"
    "encoding/hex"
    "fmt"

    "github.com/golang/protobuf/proto"
)

func main() {
    sz := 5
    internal1 := make([]*InternalMsg3, sz)
    for idx := range internal1 {
        internal1[idx] = new(InternalMsg3)
        f1 := new(InternalOneOfMessage1)
        f1.Field1 = make([]byte, sz)
        rand.Read(f1.Field1)
        f1.Field2 = make([]byte, sz)
        rand.Read(f1.Field2)
        f1.Field4 = []int32{ rand.Int31(), rand.Int31(), rand.Int31()} // 3 is a magic number
        internal1[idx].FieldOneOf = &InternalMsg3_Field1{Field1: f1}
    }
    internal2 := make([]*InternalMsg2, 20)
    internal2[0] = new(InternalMsg2)
    internal2[0].Field1 = Enum1_ENUM1_OPTION2
    internal2[0].Field2 = make([]byte, 30)
    rand.Read(internal2[0].Field2)
    internal2[0].Field3 = []byte{3, 4, 5}
    oneof := OneOfMessage3{
        Field1: internal1,
        Field2: internal2,
    }
    msg3 := MainMessage_Msg3{Msg3: &oneof}
    msg := MainMessage{Internal: &msg3}
    bb, _ := proto.Marshal(&msg)
    fmt.Println(hex.EncodeToString(bb))
}
