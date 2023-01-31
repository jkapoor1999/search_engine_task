FROM golang:1.19

LABEL Jai Kapoor and Vishal M Shekhar

WORKDIR /projects/data/

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY ./ ./

RUN go build -o se.exe ./cmd/

EXPOSE 4000

CMD [ "./se.exe" ]