package fuzz_krelay_xnet

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
                var testAddr xnet.Addr
                fuzzConsumer.GenerateStruct(&testAddr)

                testAddr.Marshal()
                return 0

            case 1:
                var testAddr xnet.Addr
                fuzzConsumer.GenerateStruct(&testAddr)

                testAddr.String()
                return 0

            case 2:
                var testAddr xnet.Addr
                fuzzConsumer.GenerateStruct(&testAddr)

                testAddr.IsZero()
                return 0

            case 3:
                addrType, _ := fuzzConsumer.GetByte()
                data, _ := fuzzConsumer.GetBytes()

                xnet.AddrFromBytes(addrType, data)
                return 0

            case 4:
                ipStr, _ := fuzzConsumer.GetString()

                xnet.AddrFromIP(ipStr)
                return 0

            case 5:
                host, _ := fuzzConsumer.GetString()

                xnet.AddrFromHost(host)
                return 0

            case 6:
                var udp xnet.UDPConn
                fuzzConsumer.GenerateStruct(&udp)
                buf, _ := fuzzConsumer.GetBytes()

                udp.ReadFrom(buf)
                return 0

            case 7:
                var udp xnet.UDPConn
                fuzzConsumer.GenerateStruct(&udp)
                buf, _ := fuzzConsumer.GetBytes()

                udp.Read(buf)
                return 0

            case 8:
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