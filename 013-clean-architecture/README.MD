# Clean Architecture Implementation

## NOTE:
This is just a copy of 012-layered-architecture-di, but we move the interfaces inward, so dependencies have to point inward.
We will notice that the business layer no longer depends on any infra layer, and that the repo and db no longer will be creatable from the web or business layer