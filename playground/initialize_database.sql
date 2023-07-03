CREATE SCHEMA `keycloak` ;
CREATE USER 'keycloak'@'%' IDENTIFIED BY 'keycloak';
GRANT ALL PRIVILEGES ON keycloak.* TO 'keycloak'@'%';

CREATE SCHEMA `orabbit_inventory`;
CREATE USER 'inventory'@'%' IDENTIFIED BY 'inventory';
GRANT ALL PRIVILEGES ON orabbit_inventory.* TO 'inventory'@'%';