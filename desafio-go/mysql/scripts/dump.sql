CREATE TABLE routes (
  id int auto_increment primary key, 
  name varchar(255), 
  source json, 
  destination json
)