# crnt-data-manager :man_office_worker:

This service is an entrypoint to CurrentTask task-manager, which means that all the actions performing with the
task-manager are being processed through this service

### Running :runner:

In the root directory there is a Makefile. All the commands you need are stored there.

1) Install brew
2) Install Go, preferably the latest version
3) Install Make. If you are using Mac, you may install XCode
4) In the root directory type:
    1) `make buf-install` (only for first time)
    2) `make generate` - generates proto and swagger files
    3) Ask me for .env file :blush:
    4) `make run` - running application

:thinking: Actually, only 'make run' is necessary. The shortest way to run the app:
1) Install Go
2) Ask me for .env file :blush:
3) Type in the root directory command which is hidden under 'make run' in Makefile

:atom_symbol: By default, app starts at port 9090 for GRPC and 8080 for HTTP.
If you have any troubles with these ports, you may change them in the .env file
in the root directory.

:computer: Swagger UI is also provided, by default on localhost:8080/swagger/

:arrows_clockwise: You may change services in the menu in the top right corner named 'Select a definition'.

:exclamation: For now, only HTTP schema is provided.