				Welcome to the JumpCloud Project
					By Peter Birk

INSTRUCTIONS

To start the server, run the script server-startup.sh located in the scripts folder.  This will call all of the necessary
Go files needed for the server to start.

jc_project>  ./scripts/server-startup.sh

The server runs on port 8080.  I left port 8080 open for ease of use, but wouldn't for production.

To run all of the testcases, issue the following script.

jc_project>  ./scripts/test_client.sh

 
TODO's if I had more time to complete.  Just mentioning so you didn't think I forgot this stuff.  :-)
- Authentication (JWT or OIDC) and Roles (Admin + Developer)
- TLS only enabled
- More OWASP Top 10 Validation
- Use of Rabbit MQ for queueing hash requests
- Use of the Swagger-generated /clients folder (see note below)
- Use of a full blown testing framework to implement all testcases.


Note: I created a Swagger OpenAPI definition for the project.  Check jc_project/internal/api/swagger.yaml.
I was going to use this to generate the server code skaffolding, but it had a few non-stdlib imports so 
I chose to skip it.  I do have API client code included in jc_project/client generated from Swagger but didn't 
use it for anything given the time.  You can review the API documentation by loading jc_project/doc/index.html 
in your browser.


