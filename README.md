#  BigQuery SQL Driver for Golang
This is an implementation of the BigQuery Client as a database/sql/driver for easy integration and usage.

# Authentication

As this is using the Google Cloud Go SDK, you will need to have your credentials available
via the GOOGLE_APPLICATION_CREDENTIALS environment variable point to your credential JSON file.

Alternatively, you can specify `apiKey` connection string parameter with API key value,
or `CredentialsFile` parameter with Credentials File(json format) Local Path
or `credentials` parameter with base-64 encoded service account or refresh token JSON credentials as the value.  
Connection string examples:  
```js
"bigquery://projectid/location/dataset?CredentialsFile=/path/filename.json"
"bigquery://projectid/location/dataset?apiKey=AIzaSyB6XK8IO5AzKZXoioQOVNTFYzbDBjY5hy4"
"bigquery://projectid/location/dataset?credentials=eyJ0eXBlIjoiYXV0..."
```

## Installation
Simple install the package to your $GOPATH with the go tool from shell: 

```go
go get -u github.com/iodeal/go-sql-bigquery
```

## Vanilla *sql.DB usage

Just like any other database/sql driver you'll need to import it 

```go
package main

import (
    "database/sql"
    _ "github.com/iodeal/go-sql-bigquery"
    "log"
)

func main() {
    config := "bigquery://{project_id}/{location}/{dataset_id}?CredentialsFile={Local Credentials File Path}"
    db, err := sql.Open("bigquery", config)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    var rs string
	err = db.QueryRow("select 'query row' msg").Scan(&rs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)
	rows, err := db.Query("select 1 union all select 2 union all select 3")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var cnt int
	for rows.Next() {
		err = rows.Scan(&cnt)
		fmt.Println(cnt)
	}
}
```


# Contribution
