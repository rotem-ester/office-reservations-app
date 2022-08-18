BINARY_NAME=or_service.out
 
build:
	go build -o ${BINARY_NAME} main.go ./office_reservation/office_reservation.go ./pkg/file_util/file_util.go ./pkg/store/store.go
 
run:
	go build -o ${BINARY_NAME} main.go ./office_reservation/office_reservation.go ./pkg/file_util/file_util.go ./pkg/store/store.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}