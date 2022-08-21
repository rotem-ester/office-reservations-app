# office-reservations-app
get monthly information about office reservations for a specific month

# Service
to run the service, cd to ./service and run `make build` to build the binary.  
the service default port is 8080, you can provide a different port via a command line argument:  
`./or_service.out 3000` for example.  
when initiating, the service loads data from the csv file `./rent_data.txt`. you can provide a different file path via the second command line argument: `./or_service.out 8080 ./some/path/file.txt`. the provided file has to be in csv format.  
to initiate the service with the default values: `./or_service.out`

# CLI
cd into ./cli and run `make build` to build the binary.  
to get information about office reservations for a given month,  
run `./ofre monthly <year in format YYY> <month in format MM>`.  
if you run your service on a different host/port than the default (`http://localhost:8080`), provide the full service url via the flag `--service-url`.

## Other options:
`./ofre revenue <year> <month>` will retrieve info about revenue only  
`./ofre capacity <year> <month>` will retrieve info about capacity only

## Example results:

2013-01: expected revenue: $8100, expected total capacity of the unreserved offices: 254  
  
2013-06: expected revenue: $15150, expected total capacity of the unreserved offices: 241  
  
2014-03: expected revenue: $37214, expected total capacity of the unreserved offices: 203  
  
2014-09: expected revenue: $86700, expected total capacity of the unreserved offices: 120  
  
2015-07: expected revenue: $76225, expected total capacity of the unreserved offices: 135  
