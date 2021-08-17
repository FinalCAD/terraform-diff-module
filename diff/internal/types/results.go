package types

import (
    "fmt"
    "sort"
    "github.com/fatih/color"
)

type Result struct {
    Change StatusType
    Diff string
}

func NewResult() *Result {
    return &Result{Change: ChangeStatus.SAME, Diff: ""}
}

func (r *Result) UpdateResult(change StatusType, diff string) {
    r.Change = change
    r.Diff = diff
}

type Results map[string]*Result

func (r Results) Print(jsonFlag bool) {
    keys := make([]string, 0, len(r))
    for k := range r {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    if jsonFlag {
        PrintJson(keys, r)
    } else {
        PrintFmt(keys, r)
    }
}

func PrintFmt(keys []string, r Results) {
    color.Unset()
    for _, key := range keys {
        if (r[key].Change > 0) {
            fmt.Printf("%s %s: %s", r[key].Change, key, r[key].Diff)
        }
    }
    color.Unset()
}

func PrintJson(keys []string, r Results) {
    i := 0
    color.NoColor = true
    fmt.Print("{")
    for _, key := range keys {
        if (r[key].Change > 0) {
            if i != 0 {
                fmt.Print(", ")
            }
            fmt.Printf("\"change %d\": \"%s %s: %s\"", i, r[key].Change, key, r[key].Diff)
            i += 1
        }
    }
    fmt.Print("}")
}
