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

1. A docker environment like a AWS EC2 with docker instaled or Kubernetes
2. A persistence store provider, actually we support AWS S3,MySQL and PostgreSQL
3. Populate this required environment variables.
   
   a. PERSISTENCE_PROVIDER='S3' or 'PostgreSQL' or 'MySQL'
   b. BUCKET_NAME='Your bucket name', only if S3 persistence provider is choosen.
   c. CONNECTION_STRING='Your connection string', only if PostgreSQL or MySQL is provided




