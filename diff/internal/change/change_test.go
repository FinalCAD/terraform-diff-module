package change

import (
    "testing"
)

func TestTransformToString(t *testing.T){
    got := TransformToString("true")
    want := "true"
    if got != want {
        t.Errorf("got %s, wanted %s", got , want)
    }

    got = TransformToString(true)
    want = "true"
    if got != want {
        t.Errorf("got %s, wanted %s", got , want)
    }

    got = TransformToString(10)
    want = "10"
    if got != want {
        t.Errorf("got %s, wanted %s", got , want)
    }
}

func TestProtectString(t *testing.T){
    got := ProtectString("true")
    want := "true"

    if got != want {
        t.Errorf("got %s, wanted %s", got , want)
    }
    got = ProtectString("t\"rue")
    want = "t\\\"rue"

    if got != want {
        t.Errorf("got %s, wanted %s", got , want)
    }
}

func TestCompare(t *testing.T){

}
