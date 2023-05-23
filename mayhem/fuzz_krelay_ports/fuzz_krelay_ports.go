package fuzz_krelay_ports

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/knight42/krelay/pkg/ports"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                var testParse ports.Parser
                fuzzConsumer.GenerateStruct(&testParse)

                testParse.Parse()
                return 0

            case 1:
                var testArgs []string
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    temp, _ := fuzzConsumer.GetString()
                    testArgs = append(testArgs, temp)
                }

                ports.NewParser(testArgs)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}