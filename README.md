# Go MBV Go 
This is the template I use for kick-starting all my Golang projects.

It is meant to provide some initial structure that is flexible but also lets you get building fast.

## Dependencies
Golang migrate: `brew install golang-migrate`

PostgreSQL: `brew install postgres`

## Makefile commands
This template uses Makefile commands to run the service, work with migrations and building binaries (coming). You have the following commands available:
- apply-migrations
- apply-x-migrations
- down-migrations
- down-x-migrations
- drop-database
- create-migration
- run 
  
#### Description 
The template uses `golang-migrate` to handle migrations. This library will generate two types for each migration you create, an `up` and a `down` version. When applying migrations the library will apply the `up` version to the database. When removing migrations, it will run the `down` versoin (which often is just a sql command to drop the table). 

**apply-migrations**: simply applies all the `up` migrations in the migrations' folder.

**apply-x-migrations**: apply x migrations from the earliest one and up. This means that if you have 5 migations, and you pass it 3, it will run the 3 first migrations. To use this command run: `make apply-x-migrations x=number-of-migrations` 

**down-migrations**: simply applies all the `down` migrations. 

**down-x-migrations**: same as `apply-x-migations` but for the down versions, also starting from the earliest migration and counting up.

**drop-database**: will use the `golang-migrate` to drop the entire database. Be vary.

**create-migration**: will generate a pair of `up` and `down` version for your migration. To name it, simply run `make create-migration name=name-of-migration`. If this
is not applied it will only have a number, i.e. `000001.sql`.

**run**: this will call the `go.sh` file and start the project. The `go.sh` script will `vet` and `fmt` the project as well as sourcing any
environmental variables you have in `.env`.

#### Migrations
A small side note on migrations. When you eventually are going to get this into staging/production, you will of course want to have your migrations applied to the DB. To do this, you need to call `ApplyMigrations()` after initializing the DB in `main.go`. This is because, at this point, the migrations should have been tested locally so we are sure that the migrations won
t break anything.

This is different from your local development environment, where I suggest you use the Makefile commands associated with migrations.

## Permissions
To use the make command: `make run` you need to give the `go.sh` permission to run. This is done by running this command in your terminal: `chmod +x go.sh` from the root of this directory.

You also need to create a `.env` file that holds all of your environmental variables. This can be done by changing the name of the `.example.env` file to `.env` and add any variables you need here.

