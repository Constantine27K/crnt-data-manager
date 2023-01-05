# crnt-data-manager :man_office_worker:
This service is an entrypoint to CurrentTask task-manager, which means that all the actions performing with the task-manager are being processed through this service

### Running :runner:
In the root directory there is a Makefile. All the commands you need are stored there. 

1) Install Go, preferably the latest version
2) Install Make. If you are using Mac, you may install XCode
3) In the root directory type:
    1) make generate - installs dependencies and generates proto and swagger files
    2) make run - running application

:atom_symbol: By default, app starts at port 9090 for GRPC and 8080 for HTTP.
If you have any troubles with these ports, you may change them in the .env file
in the root directory.

:computer: Swagger UI is also provided, by default on localhost:8080/swagger/

:arrows_clockwise: You may change services in the menu in the top right corner named 'Select a definition'.

:exclamation: For now, only HTTP schema is provided.