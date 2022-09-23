# Questions and notes

### Realisation notes

Application is made as a stateless microservice with dynamic configuration for Tax Rules.

Communication with application performed via HTTP:
 send POST request to endpoint /api/v1/calculate with city, vehicle_type and datetime records to get total tax

 ### Future improvements

 - Authentication: could be simple auth, JWT, etc.

 - More detailed response with ability to get tax per day/week/month/year
 
 - I could assume that tax rules in DB are filled from another microservice application / website. To keep microservice architecture clean and not have multiple services connected to the same DB calculator service should get such data from different microservice instead of DB.

- Logging and monitoring should be implemented

- Currently max daily tax supports only integer values that could not work if different cities have decimal values.

- Improve test coverage: currently some tests for GetTax function exist

- Fetch information about tax free transport types from another service/DB
