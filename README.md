# crnt-data-manager
This service is an entrypoint to CurrentTask task-manager, which means that all the actions performing with the task-manager are being processed through this service

### Running
In the root directory there is a Makefile. All the commands you need are stored there. 

1) Install Go, preferably the latest version
2) Install Make. If you are using Mac, you may install XCode
3) In the root directory type:
    1) make generate
    2) make run

By default, app starts at port 9090 for GRPC and 8080 for HTTP.
If you have any troubles with these ports, you may change them in the .env file
in the root directory.

Swagger UI is also provided, by default on localhost:8080/swagger/
You may change services in the menu in the top right corner named 'Select a definition'.

! For now, only HTTP schema is provided.