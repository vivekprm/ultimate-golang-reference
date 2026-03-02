Interfaces should only be used when their added value is clear. We see too many packages that declare interfaces unnecessarily, sometimes just for the sake of using interfaces. The use of interfaces when they are not necessary is called interface pollution.

# Code Example
Let’s look at a code example that contains questionable design choices that raise flags for interface pollution.

```go
package tcp

// Server defines a contract for tcp servers.
type Server interface {
 Start() error
 Stop() error
 Wait() error
}

// server is our Server implementation.
type server struct {
    /* impl */
}

// NewServer returns an interface value of type Server
// with an xServer implementation.
func NewServer(host string) Server {
 return &server{host}
}

// Start allows the server to begin to accept requests.
func (s *server) Start() error {
    /* impl */
}

// Stop shuts the server down.
func (s *server) Stop() error {
    /* impl */
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {
    /* impl */
}
```

Here is the interface pollution smell list for the code above:

- The package declares an interface that matches the entire API of its own concrete type.
- The factory function returns the interface value with the unexported concrete type value inside.
- The interface can be removed and nothing changes for the user of the API.
- The interface is not decoupling the API from change.

Let’s break down the code:

We see the declaration of the exported interface type **Server**. This interface declares an exact duplication of the API declared by the unexported concrete type **server**. These two lines of code check the box for items 1 in the smell list.

Then we see the factory function **NewServer**. This function creates a value of the unexported concrete type server and returns it to the user inside an exported interface value of type **Server**. This checks the box for item 2 in the smell list.

The next code listing shows how removing the interface changes nothing for the user:

```go
// Remove the interface and change the concrete type to be exported.

type Server struct {
    /* impl */
}

// Have the NewServer function return a pointer of the concrete type instead
// of the interface type.

func NewServer(host string) *Server {
 return &Server{host}
}
```

Having the user work with the concrete type directly doesn’t change anything for the user or the API. **This change has actually improved things because the extra level of indirection to call the methods through the interface value has been removed**. 
This checks the box for item 3 in the smell list.

Finally, if we ask what can change in the code, it is never going to be a new implementation of the Server. **Having an interface to decouple the server struct type from the API is not helping the API decouple itself from change**. This checks the final box for item 4 in the smell list.



