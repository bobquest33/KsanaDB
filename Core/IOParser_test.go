package KsanaDB
import(
        "testing" 
        "fmt"
)

func Test_ParseQueryJson(t *testing.T) {  
    query := `{"startabsolute":1389024000000,"endabsolute":1389456010500,"metric":{"groupby":["host","speed"],"aggregator":{"name":"sum","sampling":{"unit":"d","value":1}},"name":"wyatt_test"}}`
    _, err := ParseQueryJson(query)
    if err != nil {
        t.Error(err) 
    }  
}

func Test_ParseQueryJsonFail(t *testing.T) {  
    query := `xxxxxxxx{"startabsolute":1389024000000,"endabsolute":1389456010500,"metric":{"groupby":["host","speed"],"aggregator":{"name":"sum","sampling":{"unit":"d","value":1}},"name":"wyatt_test"}}`
    _, err := ParseQueryJson(query)
    if err == nil {
        t.Error(err) 
    }  
}

func Test_ParseDataJson(t *testing.T) {  
    datapoint := `[{"name":"wyatt_new","tags":{"host":"server1","speed":"10","type":"tp2"},"datapoints":[[1458528731000,0],[1458528731001,1],[1458528731002,2]]},{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"datapoints":[[1458528731003,0],[1458528731103,1],[1458528731203,2]]}]`
    _, err := ParseDataJson(datapoint) 
    if err != nil {
        t.Error(err) 
    }    
}

func Test_ParseDataJsonFail(t *testing.T) {  
    datapoint := `xxxxxxx[{"name":"wyatt_new","tags":{"host":"server1","speed":"10","type":"tp2"},"datapoints":[[1458528731000,0],[1458528731001,1],[1458528731002,2]]},{"name":"wyatt_new","tags":{"host":"server11","speed":"11","type":"tp1"},"datapoints":[[1458528731003,0],[1458528731103,1],[1458528731203,2]]}]`
    _, err := ParseDataJson(datapoint) 
    if err == nil {
        t.Error(err) 
    }    
}

func Test_ParseJsonHash(t *testing.T) {  
    data := `{"tags":["16","17","18"],"timestamp":"1389162974606","value":"6.000000"}`
    timestamp, value, tags, err := ParseJsonHash(data)
    if err != nil {
        t.Error(err) 
    }  
    
    if timestamp != 1389162974606 {
        t.Error("timestamp fail") 
        fmt.Println(timestamp)
    } 

    if value != 6 {
        t.Error("value fail") 
        fmt.Println(value)
    } 

    if tags[0] != "16" || tags[1] != "17" || tags[2] != "18" {
        t.Error("tags fail") 
        fmt.Println(tags)
    }
}

func Test_ParseJsonHashFail(t *testing.T) {  
    data := `xxxxxx{"tags":["16","17","18"],"timestamp":"1389162974606","value":"6.000000"}`
    _, _, _, err := ParseJsonHash(data)
    if err == nil {
        t.Error(err) 
    }  
}

func Test_ParseJsonHashTimeFail(t *testing.T) {  
    data := `{"tags":["16","17","18"],"timestamp":"xxxxxxx","value":"6.000000"}`
    _, _, _, err := ParseJsonHash(data)
    if err == nil {
        t.Error(err) 
    }  
}

func Test_ParseJsonHashValueFail(t *testing.T) {  
    data := `{"tags":["16","17","18"],"timestamp":"1389162974606","value":"xxxxxxxxxx"}`
    _, _, _, err := ParseJsonHash(data)
    if err == nil {
        t.Error(err) 
    }  
}

func Test_ParseJsonHashTagFail(t *testing.T) {  
    data := `{"tags":[16,17,18],"timestamp":"1389162974606","value":"6.000000"}`
    _, _, _, err := ParseJsonHash(data)
    if err == nil {
        t.Error(err) 
    }  
}
