package etcservices

import (
    "runtime"
    "os"
    "bufio"
    "unicode"
    "strings"
    _ "fmt"
)

type Service struct {
    ServiceName string
    PortNumber  string
    Protocol    string
    Alias       string
    Comment     string
}

func loadEtcServices() (map[string]Service, error) {
    etcservicesPath := "/etc/services"
    if runtime.GOOS == "windows" {
        etcservicesPath = "C:\\WINDOWS\\system32\\drivers\\etc\\services"
    }
    file, err := os.Open(etcservicesPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    rder := bufio.NewReader(file)
    serviceMap := make(map[string]Service)
    for {
        l, _, err:= rder.ReadLine()
        if err != nil {
            break
        }
        line := strings.TrimSpace(string(l))
        if line == "" || line[0] == '#' {
            continue
        }
        split := strings.SplitN(line, "#", 2)
        f := func(c rune) bool {
            return unicode.IsSpace(c) || c == '/'
        }
        fields := strings.FieldsFunc(split[0], f)
        f3 := strings.Join(fields[3:], " ")
        f4 := ""
        if len(split) > 1 {
            f4 = strings.TrimSpace(split[1])
        }
        s := Service {
            ServiceName : fields[0],
            PortNumber  : fields[1],
            Protocol    : fields[2],
            Alias       : f3,
            Comment     : f4,
        }
        serviceMap[s.PortNumber + "/" + s.Protocol] = s
    }
    return serviceMap, nil
}
