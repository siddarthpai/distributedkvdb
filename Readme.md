1. Installed boltdb as its also written in go (https://github.com/etcd-io/bbolt)

2. install it using `go get go.etcd.io/bbolt@latest`

3. We create a parseFlags function to validate what flags we passed and we call this function and hence when we run go run main.go, we get `Provide a valid DB location`

4. We use this command to create the database `go run main.go -db-location=$PWD/my.db`

---

We will now create an HTTP server with two endpoints: /get and /set.

- http.ResponseWriter is an interface : This is the response writer interface. You use it to send the response back to the client.
- \*http.Request is a pointer to an http.Request struct : This is a pointer to the request struct. It provides access to all the details of the incoming HTTP request, such as method, URL, headers, and body. The use of a pointer allows efficient access and modification of the request data.
- func(w http.ResponseWriter, r \*http.Request) is an anonymous function

We now add in user defined httpaddr and then we can input the addr if required or by default it will use 8080

now, we can rerun our program using `go run main.go -db-location=$PWD/my.db` and it will start on 8080

using curl as follows, we can see it work

`curl 'http://localhost:8080`

---

We will now create a package to seperate out the database code from the http code

db package:

1. db \*bolt.DB: This field is a pointer to a bolt.DB instance. bolt.DB is a type from the BoltDB package, a fast key/value store embedded in Go. This field will hold the reference to the actual database instance that the Database struct will interact with.
   - type Database struct { ... }: This line is defining a new type called Database. - db \*bolt.DB: Inside this struct, there's a field named db. This field is a pointer to a bolt.DB type (from the BoltDB package). This means db will hold the memory address of a bolt.DB instance, not the actual database itself.
2. func NewDatabase(db *bolt.DB) *Database { ... }: This defines a function named NewDatabase that takes a pointer to a bolt.DB instance as an argument and returns a pointer to a Database struct.
   -db \*bolt.DB: This function takes one argument, db, which is a pointer to a bolt.DB instance.
   -This line creates a new Database struct, setting its db field to the provided db pointer. The & symbol returns a pointer to this new Database struct.

But this implementation of the function isnt apt and its better to use the dbPath than the pointer itself.

The issues with it ?

1. Assumes Initialization: Assumes that the bolt.DB instance is already initialized and passed in. It does not handle any errors that might occur during the initialization of the bolt.DB instance. But it is our duty to initialise it
2. No Close Function: The caller(the user) of this function is responsible for managing and closing the bolt.DB instance. There’s no built-in way to ensure the database is closed properly or handle errors from closing it.

So we will modify the function
