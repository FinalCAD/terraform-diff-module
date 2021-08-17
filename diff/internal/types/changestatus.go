package types

import (
    "github.com/fatih/color"
)

type StatusType int

type ChangeStatusType struct {
    SAME         StatusType
    REMOVE       StatusType
    ADD          StatusType
    CHANGE       StatusType
}

var ChangeStatus = ChangeStatusType {
    SAME:        0,
    REMOVE:      1,
    ADD:         2,
    CHANGE:      3,
}

func (c StatusType) String() string {
    switch c {
        case ChangeStatus.CHANGE:
            color.Set(color.FgYellow)
            return "~"
        case ChangeStatus.ADD:
            color.Set(color.FgGreen)
            return "+"
        case ChangeStatus.REMOVE:
            color.Set(color.FgRed)
            return "-"
        default:
            color.Unset()
            return ""
        }
}
