CREATE TABLE Types(
    ID INT NOT NULL PRIMARY KEY,
    Type VARCHAR(50)
);

CREATE TABLE Users(
    ID UUID PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    Email TEXT [] NOT NULL,
    Bio VARCHAR(50) NOT NULL,
    PhoneNumber TEXT [] NOT NULL,
    TypeID INT NOT NULL,
    Status BOOLEAN NOT NULL,
    FOREIGN KEY (TypeID)
    REFERENCES Types(ID)
    ON DELETE CASCADE
);
CREATE TABLE addresses(
    ID uuid NOT NULL PRIMARY KEY,
    UserID uuid,
    Country VARCHAR(50) NOT NULL,
    City VARCHAR(50) NOT NULL,
    District VARCHAR(50) NOT NULL ,
    PostalCode BIGINT NOT NULL,
    FOREIGN KEY(UserID)
      REFERENCES Users(ID)
      ON DELETE CASCADE
);
SELECT u.*, t.Type FROM users u join types t on  t.id = u.TypeID;
SELECT * FROM addresses WHERE UserID = '81878595-a7d1-4575-94d6-60f5cd0f6a24'