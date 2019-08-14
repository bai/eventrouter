/*
 * REST API
 *
 * Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package rockset
import (
    "bytes"
    "encoding/json"
    "fmt"
    
)

type QueryResponseStats struct {
	// query time in milliseconds
	ElapsedTimeMs int64 `json:"elapsed_time_ms,omitempty"`
	// rows scanned as part of query execution
	RowsScanned int64 `json:"rows_scanned,omitempty"`
}
func (m QueryResponseStats) PrintResponse() {
    r, err := json.Marshal(m)
    var out bytes.Buffer
    err = json.Indent(&out, []byte(string(r)), "", "    ")
    if err != nil {
        fmt.Println("error parsing string")
        return
    }

    fmt.Println(out.String())
}

