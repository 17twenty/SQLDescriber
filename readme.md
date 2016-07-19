## SQL Describer ##

Simple tool for use with MySQL to dump out the structure of tables contained in a given database.
So dumb but it is what it is.

```
$ gb vendor restore
 Getting github.com/go-sql-driver/mysql
 Getting golang.org/x/crypto/ssh/terminal

$ gb build
 golang.org/x/crypto/ssh/terminal
 github.com/go-sql-driver/mysql
 cmd/SQLDescriber
$ bin/SQLDescriber --help
 Usage of bin/SQLDescriber:
   -db string
     	Database to use
   -host string
     	Host to connect to (default "localhost")
   -p	Specify that you wish to provide a password (You will be prompted)
   -port string
     	Port to connect to (default "3306")
   -u string
     	Username to connect with (default "root")

```