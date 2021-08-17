package cmd

import (
    "flag"
    "github.com/FinalCAD/terraform-diff-module/diff/internal/diff"
)

func Execute() {
    var initial string
    var updated string

    flag.StringVar(&initial, "i", "{}", "Specify initial json")
    flag.StringVar(&updated, "u", "{}", "Specify updated json")
    printJson := flag.Bool("json", false, "Output to json format")
    flag.Parse()
    var r = diff.Diff(initial, updated)
    r.Print(*printJson)
}
