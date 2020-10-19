CREATE TABLE `account` (
  `id`              INT(11) NOT NULL AUTO_INCREMENT,
  `document_number` VARCHAR(14) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `document_unique` UNIQUE (`document_number`)
);

CREATE TABLE `operation_type`
  ( 
     `id`           INT(11) NOT NULL,
     `description`  VARCHAR(40) NOT NULL, 
     PRIMARY KEY (`id`) 
  ); 

CREATE TABLE `transaction` (
  `id`                INT(11) NOT NULL AUTO_INCREMENT,
  `account_id`        INT(11) NOT NULL,
  `operation_type_id` INT(11) NOT NULL,
  `amount`            DECIMAL(10,2) NOT NULL,
  `event_date`        TIMESTAMP(6) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `account`(`id`),
  FOREIGN KEY (`operation_type_id`) REFERENCES `operation_type`(`id`)
);
