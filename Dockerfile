FROM heroiclabs/nakama-pluginbuilder:3.20.0 AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOPRIVATE "quana.com/karata"
WORKDIR /backend

COPY . .


RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM registry.heroiclabs.com/heroiclabs/nakama:3.20.0

COPY --from=go-builder /backend/backend.so /nakama/data/modules/
COPY --from=go-builder /backend/local.yml /nakama/data/