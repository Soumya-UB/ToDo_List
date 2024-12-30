1. Create a simple golang app that allows users to create, get, list, update and delete tasks.
     - App handles CRUDL operations for Task objects.
     - Uses https://pkg.go.dev/github.com/gorilla/mux#section-readme as request router and dispatcher.
3. Each task has the following properties.
     a. ID
     b. Title
     c. Description
     d. StartTime
3. Dockerize the application
     - docker compose is used to spin up multi-container application.
4. Manage secrets
7. Add an environment specific config.
8. Configure logging
9. Setup a CI/CD pipeline to deploy to OCI Compute instance
10. Configure authn/authz
     - Implemented authn using mTLS
