package types

import (
    "fmt"
    "sort"
    "strings"
    "github.com/fatih/color"
)

var ProtectedKeys = []string{
    "accesskey",
    "secretkey",
    "projectkey",
    "database",
    "_key",
    "password",
    "apikey",
    "clientid",
    "roleid",
    "token",
    "connectionstring",
}

type Result struct {
    Change StatusType
    Diff string
    NewValue string
}

type Results map[string]*Result

func NewResult() *Result {
    return &Result{Change: ChangeStatus.SAME, Diff: "", NewValue: ""}
}

func (r *Result) UpdateResult(change StatusType, diff string, newvalue string) {
    r.Change = change
    r.Diff = diff
    r.NewValue = newvalue
}

func isSensible(key string) bool {
    for _, substring := range ProtectedKeys {
        if strings.Contains(strings.ToLower(key), substring) {
            return true
        }
    }
    return false
}

func CensorValue(key string, r *Result) string {
    showdiff := r.Diff
    shownewvalue := r.NewValue
    if isSensible(key) {
        showdiff = "*****"
        shownewvalue = "*****"
        if len(r.Diff) > 4 {
            showdiff = fmt.Sprintf("%s*****%s", r.Diff[0:2], r.Diff[len(r.Diff)-2:len(r.Diff)])
        }
        if len(r.NewValue) > 4 {
            shownewvalue = fmt.Sprintf("%s*****%s", r.NewValue[0:2], r.NewValue[len(r.NewValue)-2:len(r.NewValue)])
        }
    }
    if r.Change == ChangeStatus.CHANGE {
        showdiff = fmt.Sprintf("%s ~> %s", showdiff, shownewvalue)
    }
    return fmt.Sprintf("%s : %s", key, showdiff)
}

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
            fmt.Printf("%s %s", r[key].Change, CensorValue(key, r[key]))
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
            fmt.Printf("\"change %d\": \"%s %s\"", i, r[key].Change, CensorValue(key, r[key]))
            i += 1
        }
    }
    fmt.Print("}")
}
