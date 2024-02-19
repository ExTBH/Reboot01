-- Drop existing tables (if necessary)
DROP TABLE IF EXISTS PostCategory;
DROP TABLE IF EXISTS ReactionType;
DROP TABLE IF EXISTS PostReaction;
DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS UserRole;
DROP TABLE IF EXISTS Report;
DROP TABLE IF EXISTS PromoteRequest;
DROP TABLE IF EXISTS Post;
DROP TABLE IF EXISTS Category;
DROP TABLE IF EXISTS UserNotification;
DROP TABLE IF EXISTS UserSession;
DROP TABLE IF EXISTS UploadedImage;

-- Create the Image table
CREATE TABLE UploadedImage (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    data BLOB
);

-- Create the User table
CREATE TABLE User (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type_id INTEGER NOT NULL,
    username VARCHAR(10) NOT NULL,
    first_name VARCHAR(16) NOT NULL,
    last_name VARCHAR(16) NOT NULL,
    date_of_birth DATE,
    email VARCHAR(30) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    image_id INTEGER NOT NULL,
    banned_until DATE,
    github_name VARCHAR(20),
    linkedin_name VARCHAR(20),
    twitter_name VARCHAR(20),
    FOREIGN KEY(image_id) REFERENCES UploadedImage(id),
    FOREIGN KEY (type_id) REFERENCES UserRole(id)
);

-- Create the User Role table
CREATE TABLE UserRole (
    id INTEGER PRIMARY KEY,
    role_name VARCHAR(10),
    description VARCHAR(250),
    can_post BOOLEAN,
    can_react BOOLEAN,
    can_manage_category BOOLEAN,
    can_delete BOOLEAN,
    can_report BOOLEAN,
    can_promote BOOLEAN
);

-- Create the Report table
CREATE TABLE Report (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    reporter_user_id INTEGER,
    reported_user_id INTEGER,
    report_message VARCHAR(250),
    reported_post_id INTEGER,
    time DATE,
    is_post_report boolean,
    is_pending boolean,
    report_response VARCHAR(250),
    FOREIGN KEY (reported_user_id) REFERENCES User(id),
    FOREIGN KEY (reported_post_id) REFERENCES Post(id),
    FOREIGN KEY (reporter_user_id) REFERENCES User(id)
);

-- Create the PromoteRequest table
CREATE TABLE PromoteRequest (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    description TEXT,
    time DATE,
    is_pending boolean,
    FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Create the Post table
CREATE TABLE Post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    parent_id INTEGER,
    title VARCHAR(64),
    message TEXT,
    image_id INTEGER,
    time DATE,
    like_count    INTEGER,
	dislike_count INTEGER,
	love_count    INTEGER,
	haha_count    INTEGER,
	skull_count   INTEGER,
    FOREIGN KEY (parent_id) REFERENCES Post(id),
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (image_id) REFERENCES UploadedImage(id)
);

-- Create the Notification table
CREATE TABLE UserNotification (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    comment_id VARCHAR(10),
    PostReaction_id VARCHAR(150),    
    FOREIGN KEY (comment_id) REFERENCES Post(id),
    FOREIGN KEY (PostReaction_id) REFERENCES PostReaction(id)
);

-- Create the Category table
CREATE TABLE Category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(10),
    description VARCHAR(150),
    color VARCHAR(7)
);

-- Create the PostCategory table
CREATE TABLE PostCategory (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER,
    category_id INTEGER,
    FOREIGN KEY (category_id) REFERENCES Category(id),
    FOREIGN KEY (post_id) REFERENCES Post(id)
);

-- Create the ReactionType table
CREATE TABLE ReactionType (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type VARCHAR(20)
);

-- Create the PostReaction table
CREATE TABLE PostReaction (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER,
    user_id INTEGER,
    reaction_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (post_id) REFERENCES Post(id),
    FOREIGN KEY (reaction_id) REFERENCES ReactionType(id)
);

-- Create the Session table
CREATE TABLE UserSession (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    token CHAR(64),
    user_id INTEGER,
    creation_time INTEGER,
    FOREIGN KEY (user_id) REFERENCES User(id)
);


-- Insert the default user roles
INSERT INTO UserRole 
    (id, role_name, description, can_post, can_react, can_manage_category, can_delete, can_report, can_promote)
    VALUES 
    (0, 'guest', '', 0, 0, 0, 0, 0, 0),
    (1, 'user', '', 1, 1, 0, 0, 1, 0),
    (2, 'moderator', '', 1, 1, 1, 1, 1, 0),
    (3, 'admin', '', 1, 1, 1, 1, 1, 1);


-- Insert default categories
INSERT INTO Category 
    (id, name, description, color)
    VALUES 
    (1, 'General', 'General discussion', '#000000'),
    (2, 'Announcement', 'Announcements', '#000000'),
    (3, 'Question', 'Questions', '#000000');

-- Insert default reactions
INSERT INTO ReactionType
    (id, type)
    VALUES 
    (1, 'like'),
    (2, 'dislike'),
    (3, 'haha'),
    (4, 'skull');
    

-- Insert default image 
INSERT INTO UploadedImage 
    (id, data)
    VALUES 
    (1, null);

-- Insert default users
INSERT INTO User 
    (type_id, username,
        first_name, last_name, date_of_birth,
        email, hashed_password, image_id,
        banned_until, github_name,
        linkedin_name, twitter_name)
    VALUES 
    (1, 'natheer', 'Natheer', 'Radhi', '01-01-2004',
        'random@example.com',
        '$2a$10$Jh3uNYj1l50.7eVFbmx6d.CqHZf/9TIL/z9D94qIIdgeAWY81lNEu', -- password is password123
        1, null, null, null, null);


-- Insert default posts for user 1
INSERT INTO Post 
    (user_id, parent_id, title, message, image_id, time, like_count, dislike_count, love_count, haha_count, skull_count)
    VALUES 
    (1, null, 'Welcome to the forum', 'Welcome to the forum! Feel free to post anything you want.', 1, '2021-01-01 00:00:00', 0, 0, 0, 0, 0),
    (1, null, 'First post', 'This is the first post on the forum!', 1, '2021-01-01 12:22:44', 0, 0, 0, 0, 0);