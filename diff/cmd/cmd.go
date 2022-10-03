package cmd

import (
    "flag"
    "github.com/FinalCAD/terraform-diff-module/diff/internal/change"
)

func Execute() {
    var initial string
    var updated string

    flag.StringVar(&initial, "i", "{}", "Specify initial json")
    flag.StringVar(&updated, "u", "{}", "Specify updated json")
    printJson := flag.Bool("json", false, "Output to json format")
    flag.Parse()
    var change = change.Change(initial, updated)
    change.Print(*printJson)
}
