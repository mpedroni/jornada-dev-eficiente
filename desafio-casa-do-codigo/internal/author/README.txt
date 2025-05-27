Here in the author package, the approach used consists of separating the concerns into:
  - a domain layer that contains the business logic and entities (including a repository interface)
  - a handler layer that contains the http-related logic and it is also acting as a service layer (data flow and entities choreography management).

In the [internal/category] package a different approach was adopted for comparison purposes.