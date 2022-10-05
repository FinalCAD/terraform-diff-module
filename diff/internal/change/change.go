package change

import (
    "fmt"
    "encoding/json"
    "github.com/FinalCAD/terraform-diff-module/diff/internal/types"
    "strconv"
    "strings"
)

type Secrets interface{}

func ProtectString(data string) (res string) {
    return strings.Replace(data, "\"", "\\\"", -1)
}

func TransformToString(data interface{}) (res string) {
    switch v := data.(type) {
    case float64:
        res = strconv.FormatFloat(data.(float64), 'g', 6, 64)
    case float32:
        res = strconv.FormatFloat(float64(data.(float32)), 'g', 6, 32)
    case int:
        res = strconv.FormatInt(int64(data.(int)), 10)
    case int64:
        res = strconv.FormatInt(data.(int64), 10)
    case uint:
        res = strconv.FormatUint(uint64(data.(uint)), 10)
    case uint64:
        res = strconv.FormatUint(data.(uint64), 10)
    case uint32:
        res = strconv.FormatUint(uint64(data.(uint32)), 10)
    case json.Number:
        res = data.(json.Number).String()
    case string:
        res = data.(string)
    case bool:
        res = strconv.FormatBool(data.(bool))
    case []byte:
        res = string(v)
    default:
        res = ""
    }
    return
}

func Compare(jsonOldData Secrets, jsonNewData Secrets) types.Results {
    results := make(types.Results)

    for key, _ := range jsonOldData.(map[string]interface{}) {
        results[key] = types.NewResult()
    }

    for key, _ := range jsonNewData.(map[string]interface{}) {
        results[key] = types.NewResult()
    }

    for key, _ := range results {
        oldvalue, old := jsonOldData.(map[string]interface{})[key]
        newvalue, new := jsonNewData.(map[string]interface{})[key]

        strValue := TransformToString(oldvalue)
        strValue2 := TransformToString(newvalue)

        if (!old) {
            results[key].UpdateResult(types.ChangeStatus.ADD, ProtectString(strValue2), "")
        } else if (!new) {
            results[key].UpdateResult(types.ChangeStatus.REMOVE, ProtectString(strValue), "")
        } else if (strValue != strValue2) {
            results[key].UpdateResult(types.ChangeStatus.CHANGE, ProtectString(strValue), ProtectString(strValue2))
        }
    }

    return results
}

func UnMarshal(data string) Secrets{
	var jsonData interface{}
    err := json.Unmarshal([]byte(data), &jsonData)
    if err != nil {
        fmt.Println(err)
    }
    return jsonData
}

func Change(original string, updated string) types.Results {
    return Compare(UnMarshal(original), UnMarshal(updated))
}
