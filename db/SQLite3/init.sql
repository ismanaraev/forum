CREATE TABLE IF NOT EXISTS users (
	ID TEXT UNIQUE NOT NULL,
	name CHAR(50) NOT NULL,
	username VARCHAR(50) NOT NULL UNIQUE,
	email VARCHAR(50) NOT NULL UNIQUE, 
	password VARCHAR(50) NOT NULL,
	token TEXT,
	expiretime
);
CREATE TABLE IF NOT EXISTS post (
	ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	author TEXT NOT NULL, 
	createdat INTEGER NOT NULL,
	like INTEGER DEFAULT 0,
	dislike INTEGER DEFAULT 0,
	FOREIGN KEY (author) REFERENCES users(ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS comments (
	ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	postID INTEGER,
	content TEXT NOT NULL,
	author STRING NOT NULL, 
	like INTEGER DEFAULT 0,
	dislike INTEGER DEFAULT 0,
	createdat INTEGER NOT NULL,
	FOREIGN KEY (postID) REFERENCES post(ID) ON DELETE CASCADE,
	FOREIGN KEY (author) REFERENCES users(ID) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS likePost (
	ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	userID TEXT,
	postID INTEGER DEFAULT 0,
	status INTEGER DEFAULT 0,
	FOREIGN KEY (userID) REFERENCES users(ID) ON DELETE CASCADE,
	FOREIGN KEY (postID) REFERENCES post(ID) ON DELETE CASCADE
	);
CREATE TABLE IF NOT EXISTS likeComments(
	ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	userID TEXT,
	commentsID INTEGER DEFAULT 0,
	status INTEGER DEFAULT 0,
	FOREIGN KEY (userID) REFERENCES users(ID) ON DELETE CASCADE,
	FOREIGN KEY (commentsID) REFERENCES comments(ID) ON DELETE CASCADE
	);
CREATE TABLE IF NOT EXISTS categories(
	ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(50) UNIQUE NOT NULL
	);
CREATE TABLE IF NOT EXISTS categoriesPost(
	categoryID INTEGER NOT NULL,
	postID INTEGER NOT NULL,
	FOREIGN KEY (categoryID) REFERENCES categories(ID) ON DELETE CASCADE,
	FOREIGN KEY (postID) REFERENCES post(ID) ON DELETE CASCADE
	);
CREATE TABLE IF NOT EXISTS picture(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	value TEXT NOT NULL,
	type TEXT NOT NULL,
	size INTEGER NOT NULL,
	postID INTEGER NOT NULL,
	FOREIGN KEY (postID) REFERENCES post(ID) ON DELETE CASCADE
	);