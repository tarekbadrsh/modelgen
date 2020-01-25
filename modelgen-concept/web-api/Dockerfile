
##################################
# STEP 1 build executable binary #
##################################
FROM golang AS builder

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# copy the content 
COPY . .

# install dependencies
RUN export GO111MODULES=on
RUN CGO_ENABLED=0 GOOS=linux go build -o main

##############################
# STEP 2 build a small image #
##############################
FROM scratch

# Copy our static executable.
COPY --from=builder /app/main /app/main

# Run the binary.
CMD ["/app/main"]

