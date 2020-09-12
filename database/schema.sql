
CREATE TABLE account(
	login VARCHAR(30) PRIMARY KEY,
 	createdAt TIMESTAMP NOT NULL,
  	email VARCHAR(100) NOT NULL,
  	userPass VARCHAR(30) NOT NULL
);

CREATE TABLE posts(
    id SERIAL8 PRIMARY KEY,
	title VARCHAR(150) NOT NULL,
  	createdAt TIMESTAMP NOT NULL,
  	userLogin VARCHAR(30) NOT NULL,
  	body VARCHAR(1000),
    FOREIGN KEY (userLogin) REFERENCES account(login)
);

CREATE TABLE comments(
	createdAt TIMESTAMP NOT NULL,
  	id SERIAL NOT NULL,
  	body VARCHAR(1000) NOT NULL,
  	userLogin VARCHAR(30) NOT NULL,
  	post INT8 NOT NULL,	
  
  	FOREIGN KEY (userLogin) REFERENCES account(login),
  	FOREIGN KEY (post) REFERENCES posts(id),
  	PRIMARY KEY (post,id)
  
);