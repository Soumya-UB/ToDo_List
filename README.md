1. Create a simple golang app that allows users to create, get, list, update and delete Files' metadata.
     - App handles CRUDL operations for File objects.
     - Uses https://pkg.go.dev/github.com/gorilla/mux#section-readme as request router and dispatcher.
3. Each task has the following properties.
     a. ID
     b. Name
     c. Size
     d. CreatedTime
     e. LastUpdatedTime
     f. IsDir
3. Dockerize the application
     - docker compose is used to spin up multi-container application.
4. Manage db passwords as docker secrets.
7. Add an environment specific config.
8. Configure logging
9. Configure authn/authz
     - Implemented authn using mTLS
