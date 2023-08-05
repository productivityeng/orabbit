[![Unit Testing](https://github.com/productivityeng/orabbit/actions/workflows/unit_testing.yaml/badge.svg)](https://github.com/productivityeng/orabbit/actions/workflows/unit_testing.yaml)
[![codecov](https://codecov.io/gh/productivityeng/orabbit/branch/main/graph/badge.svg?token=DDBUGyuGSt)](https://codecov.io/gh/productivityeng/orabbit)

# ORabbit
Organization Rabbit. The objective of this project is to promote a platform for managing multiple clusters of RabbitMQ in an organizational/enterprise environment. Our focus is on auditability, pattern consistency, low administrative overhead, and ease of use. If, at any time in your environment, any of the following questions arise, ORabbit will add value to your operation.

1. When and by whom was this queue created?
2. When and by whom was this exchange created?
3. What application and business unit is responsible for consuming this queue?
4. How can we create a pattern for the names of queues and exchanges?
5. How can we approve changes in our cluster before production?

## What we need to install ORabbit?

1. A docker environment like AWS EC2 with Docker installed or Kubernetes.
2. A persistence store provider, currently MySQL.
3. Populate appropriate environment variables.

## What are the dependencies of this project?

1. MySQL Database for persistence storage
2. Keycloak for user and permission management

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
