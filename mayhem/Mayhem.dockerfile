FROM golang:1.19 as builder

COPY . /krelay
WORKDIR /krelay/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_krelay_ports && go-fuzz-build -libfuzzer -o fuzz_krelay_ports.a && \
    clang -fsanitize=fuzzer fuzz_krelay_ports.a -o fuzz_krelay_ports.libfuzzer
    
RUN cd fuzz_krelay_xnet && go-fuzz-build -libfuzzer -o fuzz_krelay_xnet.a && \
    clang -fsanitize=fuzzer fuzz_krelay_xnet.a -o fuzz_krelay_xnet.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /krelay/mayhem/fuzz_krelay_ports/fuzz_krelay_ports.libfuzzer /
COPY --from=builder /krelay/mayhem/fuzz_krelay_xnet/fuzz_krelay_xnet.libfuzzer /