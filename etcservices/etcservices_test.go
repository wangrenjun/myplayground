package etcservices

import (
    "testing"
)

func TestEtcServices(t *testing.T) {
    services, err := loadEtcServices()
    if err != nil || len(services) == 0 {
        t.Errorf("loadEtcServices failed")
    }
    for k, v := range services {
        if v.ServiceName == "" || v.PortNumber == "" || v.Protocol == "" {
            t.Errorf("error")
        }
        t.Logf("%v: %+v", k, v)
    }
}
