# desafio-casa-do-codigo

This project aims to implement the requirements of one of the challenges from the _Jornada Dev Eficiente_. The challenge consists of creating a REST API with some features inspired by _Casa do Código_, an online bookstore.

Additionally, I took the opportunity to explore different ways of organizing the application's packages in order to identify potential advantages and disadvantages of each approach. Each module under the [`internal`](./internal) directory contains a README explaining the approach used for organizing the components within that module.

## Insights & lessons learned

### Business logic right on the HTTP handler (a.k.a controller) code

In frameworks like Spring on Java's ecosystem, there is a non-intrusive approach, which means the HTTP handler in such applications has basically user code. In such cases, may be reasonable keeping the business rules right on the handler level specially when dealing with a simple application, while creating an additional layer (i.e a service layer) just to deal with the business logic may not pay out due the increase of indirections on the code.

That is different on Go's handler/controller because in Go we usually need a significant amount of HTTP-related code on the handler code (request/response decoding, request parsing, error handling, etc.), which adds more complexity to the code. Also, testing a business rule encased by a handler directly may make unit testing such rule harder, as we will need to provide an `http.ResponseWriter` and `*http.Request` to satisfy the handler's func signature.

### SQL queries right on service layer

In general, I prefer to keep the SQL code in a separated package in order to separate the business logic from the SQL code. However, in this case, I decided to keep things together in the service layer to see how it would look like and I kinda liked it. Due to the simplicity of the application, I don't think an additional layer would pay out (specially because I would test everything with the database integration anyway). Also, it seems that would be relatively easy to extract the SQL code into a separated package once things start getting more complex.

### Each module managing its setup

Each module has a `Setup` function that is responsible for setting up the module's dependencies and registering the HTTP handlers, letting to the main package creating common dependencies such as database connections and configuring the router/mux. This approach allows each module to be self-contained and makes it easier to manage dependencies between modules. It also allows for better separation of concerns, as each module can focus on its own functionality without worrying about the rest of the application. Besides, it allows for better testing, as each module can be tested independently.
