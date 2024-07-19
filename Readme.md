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
