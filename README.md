# This is template to create golang project using gin, casbin and gorm

## Authentication

By default it is set to have some noAuth endpoins and the others that require authentication<br>
How the authentication is done:<br>

in the request you client is asked to send 2 headers:
1. k - api key provided by you from users table
2. s - the MD5 signature for the string "[payload]+[API secret]"

Note: Payload for POST is its body, for GET - full URL without domain

## Database, migrations and policy updates
PosgreSQL is used as a default database system<br>
You can change to any other database from [this list](https://gorm.io/docs/connecting_to_the_database.html)

for the convience there is a flag -act:
 - app (default) - runs the program
 - migrate - migrates db
 - updateP - updates default policies

## Further documentation and help
 - [Gin](https://gin-gonic.com/docs/)
 - [Gorm](https://gorm.io/docs/index.html)
 - [Casbin](https://casbin.org/docs/en/overview)
 - [Viper](https://github.com/spf13/viper)

## P.S

Don't hesitate to request PULL and give constructive feedbacks

