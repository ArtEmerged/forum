![01-edu-system-blue](https://raw.githubusercontent.com/GoArtyom/study/1a66b22c5b511ccce94b582481a45dfd7f001d3a/alem.svg)

# Forum

<details>
<summary>Intro</summary>

<br>

## To build and run the project follow the steps below:

<br>

### Clone repository

```bash
    git clone git@github.com:ArtEmerged/forum.git
```

### Move to the direcroty

```bash
    cd forum
```

### Run Locally

- with makefile

```bash
    make build
    make run
```

- without docker

```bash
    go run cmd/main.go
```

- with docker

```bash
    docker build -t forum .
    docker run -p 8080:8080 forum
```

- server will run on the next route

```
    http://localhost:8080
```

</details>

<details>
<summary>Technical Requirements</summary>

## Forum Technical Requirements

### Objectives

This project consists of creating a *web forum* that allows:

- communication between users;
- associating categories to posts;
- liking and disliking posts and comments;
- filtering posts.

#### SQLite

In order to store the data your forum (like users, posts, comments, etc.) you will use the database library SQLite.

SQLite is a popular choice as an embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

To structure your database and to achieve better performance, we highly advise you to take a look at the **entity relationship diagram** and build one based on your own database.

- You must use at least one *SELECT*, one *CREATE* and one *INSERT* queries.

To know more about SQLite, you can check the [SQLite page](https://www.sqlite.org/index.html).

#### Authentication

In this segment the client must be able to `register` as a new user on the forum, by inputting their credentials. You also have to create a `login session` to access the forum and be able to add posts and comments.

You should use *cookies to allow each user to have only one opened session*. Each of these sessions must contain an *expiration date*. It is up to you to decide how long the cookie stays "alive". The use of *UUID* is a Bonus.

##### Instruction for user registration:

- Must ask for *email*
	- When the email is already taken, return an error response.
- Must ask for *username*
- Must ask for *password*
	- The password must be encrypted when stored (Bonus). 

The forum must be able to check if the email provided is present in the database and that all credentials are correct. It has to check whether the password provided is the same as the obe stored in the database. If the passwords do not match, it has to return an error response.

#### Communication

In order for users to communicate between each other, they will have to be able to create posts and comments.

- Only registered users will be able to create posts and comments;
- When registered users are creating a post, they can associate one or more categories to it;
	- The implementation and choice of categories is up to you.
- The posts and comments should be visible to all users (registered or not);
- Non-registered users will only be able to see posts and comments.

#### Likes and dislikes

Only registered users will be able to like or dislike posts and comments.

The number of likes and dislikes should be visible by all users (registered or not).

#### Filters

You need to implement a filter mechanism, that will allow users to filter the displayed posts by:

- categories;
- created posts;
- liked posts.

You can look at filtering by categories as subforums. A subforum is a section of an online forum dedicated to a specific topic.

Note that the last two are only available for registered users and must refer to the logged in user.

#### Docker

For the forum project you must use *Docker*.

### Instructions

- You must use **SQLite**;
- You must handle website errors, HTTP status;
- You must handle all sort of technical errors;
- The code must respect the **good practices**;
- It is recommended to have **test files** for *unit testing*.

### Allowed packages

- All standard Go packages are allowed
- *sqlite3*
- *bcrypt*
- *UUID*

> You must not use any fronted libraries or frameworks like React, Angular, Vue etc.

This project will help you learn about:

- The basics of web:
	- HTML
	- HTTP
	- Sessions and cookies

- Using and setting up Docker
	- Containerizing an application
	- Compatibility/Dependency
	- Creating images

- SQL language
	- Manipulation of databases

- The basics of encryption

</details>

<details>
<summary>Architecture & Design</summary>

<br>

### ERD 

<br>

![ERD](https://github.com/ArtEmerged/study/blob/main/db_forum.png?raw=true)

<br>
<br>

### Routing Requests

<br>

| HTTP Method | URL Pattern                  | Handler                   | Action                                      |
|-------------|------------------------------|---------------------------|---------------------------------------------|
| Any         | /static/                     | fileServer                | Serves static files                         |
| GET         | /                            | h.indexGET                | Display the index page                      |
| GET         | /signin                      | h.signinGET               | Display the signin page                     |
| POST        | /auth/signin                 | h.signinPOST              | Process signin form submission              |
| GET         | /signup                      | h.signupGET               | Display the signup page                     |
| POST        | /auth/signup                 | h.signupPOST              | Process signup form submission              |
| POST        | /auth/signout                | h.signoutPOST             | Process signout (authenticated)             |
|-------------|------------------------------|---------------------------|---------------------------------------------|
| GET         | /auth/google/signin          | h.signinGoogle            | Initiate Google signin process              |
| GET         | /google/callback             | h.callbackGoogle          | Handle Google signin callback               |
| GET         | /auth/github/signin          | h.signinGithub            | Initiate GitHub signin process              |
| GET         | /github/callback             | h.callbackGithub          | Handle GitHub signin callback               |
|-------------|------------------------------|---------------------------|---------------------------------------------|
| GET         | /post                        | h.onePostGET              | Display a single post                       |
| GET         | /post/create                 | h.createPostGET_POST      | Display the create post page                |
| DELETE      | /post/delete                 | h.deletePostDELETE        | Delete a post (authenticated)               |
| POST        | /post/update                 | h.updatePostGET_POST      | Update a post (authenticated)               |
| POST        | /post/vote/create            | h.createPostVotePOST      | Create a vote for a post (authenticated)    |
|-------------|------------------------------|---------------------------|---------------------------------------------|
| POST        | /comment/create              | h.createCommentPOST       | Create a comment (authenticated)            |
| DELETE      | /comment/delete              | h.deleteCommentDELETE     | Delete a comment (authenticated)            |
| POST        | /comment/update              | h.updateCommentGET_POST   | Update a comment (authenticated)            |
| POST        | /comment/vote/create         | h.createCommentVotePOST   | Create a vote for a comment (authenticated) |
|-------------|------------------------------|---------------------------|---------------------------------------------|
| GET         | /filterposts                 | h.filterPostsGET          | Display filtered posts                      |
| GET         | /myactivity                  | h.myActivityGET           | Display user's activity (authenticated)     |
| GET         | /mynotifications             | h.myNotificationsGET      | Display user's notifications (authenticated)|
|-------------|------------------------------|---------------------------|---------------------------------------------|
| PATCH       | /moderator/request           | h.moderatorRequestPATCH   | Process moderator request (authenticated)   |
| POST        | /post/reporting              | h.reportingPostPOST       | Process post reporting (authenticated moderator)|
| GET         | /admin                       | h.adminGET                | Display admin panel (authenticated admin)   |
| DELETE      | /admin/report                | h.adminReportDELETE       | Delete admin report (authenticated admin)   |
| DELETE      | /admin/categories/delete     | h.adminCategoriesDELETE   | Delete admin categories (authenticated admin)|
| POST        | /admin/categories/create     | h.adminCategoriesCREATE   | Create admin categories (authenticated admin)|
| PATCH       | /admin/moderator-request     | h.adminModeratorRequestPATCH| Process admin moderator request (authenticated admin)|
</details>
