# Forum

This project consists in creating a web forum that allows :

- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.



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


#### Audit list

<a href="https://github.com/01-edu/public/tree/master/subjects/forum/audit" target="_blank">forum audit</a>