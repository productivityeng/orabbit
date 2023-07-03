[![Test Inventory Service](https://github.com/productivityeng/orabbit/actions/workflows/inventory_test.yaml/badge.svg)](https://github.com/productivityeng/orabbit/actions/workflows/inventory_test.yaml)
[![codecov](https://codecov.io/gh/productivityeng/orabbit/branch/main/graph/badge.svg?token=DDBUGyuGSt)](https://codecov.io/gh/productivityeng/orabbit)
# ORabbit
Organization Rabbit. The objective of this project is promote a platform for managing multiple clusters of 
rabbitmq in an organizational/enterprise environment. Our focus in auditability,pattern consistency, low administrative overhead
and ease of use.
If in your environment some time one of the following question arise, ORabbit will
aggregate value to your operation.

```
1. When and by who this queue was created ?
2. When and by who this exchange was created ?
3. What application and bussiness unit is responsible for consuming this queue?
4. How we can create a pattern for the name of queues and exchanges ?
5. How we can approve the changes in our cluster before production ?
```

# What we need to install ORabbit ?

```
1. A docker environment like a AWS EC2 with docker instaled or Kubernetes
2. A persistence store provider, actually MySQL.
3. Populate appropriate environment variables.
```
## What are the depencies of this project ?
```
1. MySQL Database for persistence storage
2. Keycloack for user and permission management
```


## Quality Gates
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=bugs)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=coverage)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=productivityeng_orabbit&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=productivityeng_orabbit)