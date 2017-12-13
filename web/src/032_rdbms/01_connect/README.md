# Create an SQL instance on GCP
1. Create an SQL instance
  - [Follow this tutorial] (https://cloud.google.com/go/getting-started/using-cloud-sql)
2. Add your local ip address on network authorization settings
3. Create a new database
4. Create a new user
5. Download & install the Cloud SQL Proxy to connect to you Cloud SQL instance
  - [Download] (https://cloud.google.com/go/getting-started/using-cloud-sql#install_the_sql_proxy)
6. Make the proxy executable
  - "chmod +x cloud_sql_proxy"
7. Start the Cloud SQL Proxy using the connectionName
  - ./cloud_sql_proxy -instances="[YOUR_INSTANCE_CONNECTION_NAME]"=tcp:3306


# Using MySQL

1. Install MySQL
  - [Download MySQL Community Server](http://dev.mysql.com/downloads/)
1. We will need a MySQL driver
  - go get github.com/go-sql-driver/mysql
  - [read the documentation](https://github.com/go-sql-driver/mysql#installation)
  - [see all SQL drivers](https://github.com/golang/go/wiki/SQLDrivers)
  - [Astaxie's book](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html)
1. Include the driver in your imports
  - _ "github.com/go-sql-driver/mysql"
  - [Read the documentation](https://github.com/go-sql-driver/mysql#usage)
1. Determine the Data Source Name
  - user:password@tcp(localhost:5555)/dbname?charset=utf8
  - [Read the documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
1. Open a connection
  - db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")

[package sql](https://godoc.org/database/sql)