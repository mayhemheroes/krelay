package fuzz_krelay_udp

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/knight42/krelay/pkg/xnet"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {

            case 0:
                var udp xnet.UDPConn
                fuzzConsumer.GenerateStruct(&udp)
                buf, _ := fuzzConsumer.GetBytes()

                udp.ReadFrom(buf)
                return 0

            case 1:
                var udp xnet.UDPConn
                fuzzConsumer.GenerateStruct(&udp)
                buf, _ := fuzzConsumer.GetBytes()

                udp.Read(buf)
                return 0

            case 2:
                host, _ := fuzzConsumer.GetString()
                temp, _ := fuzzConsumer.GetInt()
                port := uint16(temp)

                xnet.JoinHostPort(host, port)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}