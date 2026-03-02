Packages exists to help provide support for specific problems that are commonly found in the different applications we are building. A package API should be intuitive and simple to use so application developers can focus on their concerns and hopefully develop their applications faster. Tests are an artifact of development and exist to make sure the code we are writing has integrity before it is published. Tests are not a part of the application. They do not get built with the application and none of that code runs when the application is running.

When designing a package API that will be used by an application developer, we have been taught to **focus on writing testable API’s first**. With a focus on not only what the API needs in terms of testing, but also what the application developer needs in terms of testing. This idea of writing API’s with a focus on what the application developer needs for their tests is something that I don’t agree with for Go. I believe **package developers should focus on how the application developer needs to use the API in their applications, not their tests**. Application tests are solely a concern for and the responsibility of the application developer, not the package developer.

**Note**: Who am I to say as a package developer, what your application needs in terms of testing? This is not the purpose of my API. I can’t focus on what you may or may not need to do with your tests. I want to focus the API on providing you the simplest and most intuitive way to solve the problems for your application. Focusing on your need to test, assuming what you need in terms of testing, is beyond the scope of the package API I’m writing and in my opinion a slippery slope.

Let’s look at an example that shows how focusing on what the application developer needs for their applications, and not their tests, can help to keep package API’s simpler, minimized and more intuitive.

# Code Example
One day I am asked to write a new package that provides an API to the internal queuing system that all applications are required to use. The API has to provide basic publish and subscribe support. Since this is the only queuing system that is allowed to be used, this new pubsub package can be implemented using only concrete types. Nothing can change since we will not be required to support another queuing system, therefore interfaces are not required.

With this in mind, I write the following package:

## Listing 1
```go
// Package pubsub simulates a package that provides
// publication/subscription type services.
package pubsub

// PubSub provides access to a queue system.
type PubSub struct {
    /* impl */
}

// New creates a pubsub value for use.
func New(/* impl */) *PubSub {
    ps := PubSub{
       /* impl */
    }

    /* impl */
    return &ps
}

// Publish sends the data for the specified key.
func (ps *PubSub) Publish(key string, v interface{}) error {
    /* impl */
}

// Subscribe requests the data for the specified key.
func (ps *PubSub) Subscribe(key string) error {
    /* impl */
}
```

If we look at the code for the pubsub package in listing 1 we see a concrete type named **PubSub** declared with two methods. One method is named **Publish** and the other method is named **Subscribe**. We have a factory function named **New** that creates and initializes a PubSub value for use.

There is no need for an interface because the application developer who will use the package does not need to provide any implementation details. Further, we do not require supporting different queuing systems internally inside the package. I write tests that hit the actual system so I know 100% that the package is working. Then I publish the package and tell the team it is ready for use.

Not too long after the team starts using the package, a team member approaches me with a problem. They are trying to write tests for their application and they don’t have access to the internal queuing system when their tests run. They ask me to provide an interface so they can mock access to the internal queuing system for their tests. I very quickly tell them NO and simply state that the pubsub package does not need an interface so I won’t be providing one. However, **if they need an interface for their testing, nothing is stopping them from declaring one for themselves**

## Listing 2:
```go
// publisher is an interface to allow this package to mock the
// pubsub package support.
type publisher interface {
    Publish(key string, v interface{}) error
    Subscribe(key string) error
}
```

Listing 2 shows the declaration of the **publisher** interface I tell the application developer to declare. The interface declares the set of methods associated with my pubsub package. This means that **the concrete type PubSub in my package implements this new publisher interface**.

With this interface declared by the application, the developer can now write their application and tests to use this interface. My PubSub value can now be decoupled from their app and they have the ability to mock this behavior in their tests.

## Listing 3:
```go
// mock is a concrete type to help support the mocking of the
// pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {
    /* impl */
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {
    /* impl */
}
```

Listing 3 shows a mocking implementation of the publisher interface with the declaration of the concrete type mock. This concrete type can be used for testing since it provides its own implementation that does not need to talk with the physical system.

# Conclusion
I disagree with designing a package API that focuses on the application developer’s need to write tests. Since the compiler will identify interface compliance through convention and not configuration, the application developer has the ability to decouple the things they need decoupled and to apply this throughout their applications and tests. I would like to see package developers take an application focused approach to API design as a first priority. I believe this will keep package API’s simpler, minimized and more intuitive for application developers.
