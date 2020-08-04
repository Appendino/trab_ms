CREATE DATABASE `moviedb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

use `moviedb`;

CREATE TABLE `moviedb`.`movies` (
    `movieId` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `genre` VARCHAR(80) NOT NULL,
    `description` VARCHAR(1000) NOT NULL,
    `views` INT,
    PRIMARY KEY (`movieId`));

INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Zack and Miri Make a Porno", "Romance", "The description of Zack and Miri Make a Porno is a Romance movie", 64);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Youth in Revolt", "Comedy", "The description of Youth in Revolt is a Comedy movie", 68);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("You Will Meet a Tall Dark Stranger", "Comedy", "The description of You Will Meet a Tall Dark Stranger is a Comedy movie", 43);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("When in Rome", "Comedy", "The description of When in Rome is a Comedy movie", 15);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("What Happens in Vegas", "Comedy", "The description of What Happens in Vegas is a Comedy movie", 28);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Water For Elephants", "Drama", "The description of Water For Elephants is a Drama movie", 60);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("WALL-E", "Animation", "The description of WALL-E is a Animation movie", 96);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Waitress", "Romance", "The description of Waitress is a Romance movie", 89);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Waiting For Forever", "Romance", "The description of Waiting For Forever is a Romance movie", 6);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Valentine's Day", "Comedy", "The description of Valentine's Day is a Comedy movie", 17);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Tyler Perry's Why Did I get Married", "Romance", "The description of Tyler Perry's Why Did I get Married is a Romance movie", 46);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Twilight: Breaking Dawn", "Romance", "The description of Twilight: Breaking Dawn is a Romance movie", 26);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Twilight", "Romance", "The description of Twilight is a Romance movie", 49);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Ugly Truth", "Comedy", "The description of The Ugly Truth is a Comedy movie", 14);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Twilight Saga: New Moon", "Drama", "The description of The Twilight Saga: New Moon is a Drama movie", 27);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Time Traveler's Wife", "Drama", "The description of The Time Traveler's Wife is a Drama movie", 38);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Proposal", "Comedy", "The description of The Proposal is a Comedy movie", 43);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Invention of Lying", "Comedy", "The description of The Invention of Lying is a Comedy movie", 56);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Heartbreak Kid", "Comedy", "The description of The Heartbreak Kid is a Comedy movie", 30);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Duchess", "Drama", "The description of The Duchess is a Drama movie", 60);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Curious Case of Benjamin Button", "Fantasy", "The description of The Curious Case of Benjamin Button is a Fantasy movie", 73);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("The Back-up Plan", "Comedy", "The description of The Back-up Plan is a Comedy movie", 20);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Tangled", "Animation", "The description of Tangled is a Animation movie", 89);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Something Borrowed", "Romance", "The description of Something Borrowed is a Romance movie", 15);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("She's Out of My League", "Comedy", "The description of She's Out of My League is a Comedy movie", 57);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Sex and the City Two", "Comedy", "The description of Sex and the City Two is a Comedy movie", 15);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Sex and the City 2", "Comedy", "The description of Sex and the City 2 is a Comedy movie", 15);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Sex and the City", "Comedy", "The description of Sex and the City is a Comedy movie", 49);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Remember Me", "Drama", "The description of Remember Me is a Drama movie", 28);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Rachel Getting Married", "Drama", "The description of Rachel Getting Married is a Drama movie", 85);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Penelope", "Comedy", "The description of Penelope is a Comedy movie", 52);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("P.S. I Love You", "Romance", "The description of P.S. I Love You is a Romance movie", 21);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Over Her Dead Body", "Comedy", "The description of Over Her Dead Body is a Comedy movie", 15);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Our Family Wedding", "Comedy", "The description of Our Family Wedding is a Comedy movie", 14);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("One Day", "Romance", "The description of One Day is a Romance movie", 37);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Not Easily Broken", "Drama", "The description of Not Easily Broken is a Drama movie", 34);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("No Reservations", "Comedy", "The description of No Reservations is a Comedy movie", 39);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Nick and Norah's Infinite Playlist", "Comedy", "The description of Nick and Norah's Infinite Playlist is a Comedy movie", 73);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("New Year's Eve", "Romance", "The description of New Year's Eve is a Romance movie", 8);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("My Week with Marilyn", "Drama", "The description of My Week with Marilyn is a Drama movie", 83);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Music and Lyrics", "Romance", "The description of Music and Lyrics is a Romance movie", 63);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Monte Carlo", "Romance", "The description of Monte Carlo is a Romance movie", 38);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Miss Pettigrew Lives for a Day", "Comedy", "The description of Miss Pettigrew Lives for a Day is a Comedy movie", 78);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Midnight in Paris", "Romance", "The description of Midnight in Paris is a Romance movie", 93);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Marley and Me", "Comedy", "The description of Marley and Me is a Comedy movie", 63);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Mamma Mia!", "Comedy", "The description of Mamma Mia! is a Comedy movie", 53);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Made of Honor", "Comedy", "The description of Made of Honor is a Comedy movie", 13);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Love Happens", "Drama", "The description of Love Happens is a Drama movie", 18);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Love & Other Drugs", "Comedy", "The description of Love & Other Drugs is a Comedy movie", 48);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Life as We Know It", "Comedy", "The description of Life as We Know It is a Comedy movie", 28);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("License to Wed", "Comedy", "The description of License to Wed is a Comedy movie", 8);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Letters to Juliet", "Comedy", "The description of Letters to Juliet is a Comedy movie", 40);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Leap Year", "Comedy", "The description of Leap Year is a Comedy movie", 21);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Knocked Up", "Comedy", "The description of Knocked Up is a Comedy movie", 91);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Killers", "Action", "The description of Killers is a Action movie", 11);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Just Wright", "Comedy", "The description of Just Wright is a Comedy movie", 45);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Jane Eyre", "Romance", "The description of Jane Eyre is a Romance movie", 85);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("It's Complicated", "Comedy", "The description of It's Complicated is a Comedy movie", 56);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("I Love You Phillip Morris", "Comedy", "The description of I Love You Phillip Morris is a Comedy movie", 71);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("High School Musical 3: Senior Year", "Comedy", "The description of High School Musical 3: Senior Year is a Comedy movie", 65);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("He's Just Not That Into You", "Comedy", "The description of He's Just Not That Into You is a Comedy movie", 42);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Good Luck Chuck", "Comedy", "The description of Good Luck Chuck is a Comedy movie", 3);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Going the Distance", "Comedy", "The description of Going the Distance is a Comedy movie", 53);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Gnomeo and Juliet", "Animation", "The description of Gnomeo and Juliet is a Animation movie", 56);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Gnomeo and Juliet", "Animation", "The description of Gnomeo and Juliet is a Animation movie", 56);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Ghosts of Girlfriends Past", "Comedy", "The description of Ghosts of Girlfriends Past is a Comedy movie", 27);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Four Christmases", "Comedy", "The description of Four Christmases is a Comedy movie", 26);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Fireproof", "Drama", "The description of Fireproof is a Drama movie", 40);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Enchanted", "Comedy", "The description of Enchanted is a Comedy movie", 93);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Dear John", "Drama", "The description of Dear John is a Drama movie", 29);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Beginners", "Comedy", "The description of Beginners is a Comedy movie", 84);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("Across the Universe", "Romance", "The description of Across the Universe is a Romance movie", 54);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("A Serious Man", "Drama", "The description of A Serious Man is a Drama movie", 89);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("A Dangerous Method", "Drama", "The description of A Dangerous Method is a Drama movie", 79);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("27 Dresses", "Comedy", "The description of 27 Dresses is a Comedy movie", 40);
INSERT INTO `moviedb`.`movies` (`name`, `genre`, `description`, `views`) VALUES ("(500) Days of Summer", "Comedy", "The description of (500) Days of Summer is a Comedy movie", 87);

create user 'moviews' identified by '123456'; 
grant all on moviedb.* to 'moviews';


CREATE DATABASE userdb /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

use userdb;

create table users (
  userid INT NOT NULL AUTO_INCREMENT,
  username varchar(80) NOT NULL UNIQUE,
  PRIMARY KEY(userid));

create table user_list (
  id INT NOT NULL AUTO_INCREMENT,
  userid INT NOT NULL,
  moveid INT NOT NULL,
  PRIMARY KEY(id));

alter table user_list add constraint watched_uq unique(userid, moveid);

create table watched_movies (
  id INT NOT NULL AUTO_INCREMENT,
  userid INT NOT NULL,
  moveid INT NOT NULL,
  PRIMARY KEY(id));

alter table watched_movies add constraint watched_uq unique(userid, moveid);

create table voted_movies (
  id INT NOT NULL AUTO_INCREMENT,
  userid INT NOT NULL,
  moveid INT NOT NULL,
  PRIMARY KEY(id));

alter table voted_movies add constraint watched_uq unique(userid, moveid);

insert into users(username) values ('axel');
insert into users(username) values ('henrique');
insert into users(username) values ('pedro');

create user 'userws' identified by '123456'; 
grant all on userdb.* to 'userws';
