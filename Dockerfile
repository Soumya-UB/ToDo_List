FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
# RUN MKDIR /app/config
# COPY config/default.yml ./config/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /ToDo_List

EXPOSE 8080

CMD ["/ToDo_List", "&"]