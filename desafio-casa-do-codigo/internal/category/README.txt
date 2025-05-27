Here in category package we are using a different approach compared to the [internal/author] package.
Instead of using the handler as the service layer, we are actually using a separate service layer
to handle the business logic, but the service is dealing with the database directly (SQL directly).

Although I don't like that much having business logic mixed up with database concerns, it seems to be acceptable
in this context since the application is small and the business logic is simple.
Besides, it would't make much sense to test the service layer without the database anyway.

This approach allows us to keep the handler layer thin, focusing on request handling and response formatting. In addition of that,
we can always refactor the service layer later if we need to add more complex business logic or if we want to introduce a repository pattern without impacting the handler layer.

Check the [internal/author] for a different approach on package organization.