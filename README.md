# FORUM
## Things to make:

### Backend

#### Users system
##### [ ] Repo
AddUser(username string, email string, password hashedPassword) error \
GetUser(userId) (models.User, error) (models.User, error) \
CheckToken(token string) error


##### [ ] Models
type User struct \
Id string \
Username string \
Email string \
Password hashedPassword \
Token token

##### [ ] Service
RegisterUser(username string, email string, password string)(error) // receives password in plaintext from handler \
Authorize(username string, email string, password string) (token, error) \
LogOut(token string) error \
CheckToken(token string) error 

##### [ ] Handler
AddUserHandler(w, r) \
AuthorizeHandler(w, r) \
LogOutHandler(w, r)

###### [ ] Middleware
IsAuthorized(token string) error

#### Workflow:
##### Add user 
1. Middleware checks if user if already authorized, by using service.CheckToken()(service.CheckToken() calls repo.CheckToken() and check if token is valid), then handler receives and handles user registration request, which contains username, email, password, then calls RegisterUser() with request, and receives nil, if user was successfully registred 
2. RegisterUser() validates input, hashes password, then calls repo.GetUser() to check if user exists, then calls repo.AddUser() to add user to database, if no error occured, returns nil 
3. repo.AddUser() performs sql command to add user to database and returns successful status

##### Authorize
1. Middleware checks if user is already authorized, then handler receives authorization POST request, which contains username and password, then calls service.Authorize() to authorize user, receives token and sends it back to user as cookie
3. Authorize user validates input, hashes password, then calls repo.GetUser() to see if user exists, then creates token for user, adds it to db using repo.AddToken() 
4. repo.GetUser() returns models.User. repo.AddUserToken() receives token and models.User, then adds the token to the associated user in db

##### Log Out 
1. Middleware checks if user is authorized, then handler receives POST request with userId and userToken and calls service.Logout() with this data, receives the bool which indicates if user has been deauthorized or not, and acts accordingly
2. service.Logout() calls repo.GetUser() to get models.User, then calls repo.DeleteUserToken() to delete/expire user session token and returns the status to handler
3. repo.DeleteUserToken() receives the user, deletes user's token from database, then returns status to the service

#### Posts system

##### [ ] Repo
GetPost(id string) \
AddPost (Title string, Text string, Author model.User) (Id string, error) \
ChangePost(Id string, NewTitle string, NewText string) (error) \
DeletePost(Id string) error 


##### [ ] Model
Post struct \
Id string \
Title string \
Text string \
Author model.User \
Comments []models.Comment \
Likes []models.Like

##### [ ] Service
CreatePost(Title, Text, Author string) (models.Post, error) \
GetAllPosts()([]models.Post, error) \
GetPostById(id string) (models.Post, error) \
EditPost(id, Title, Text string) error

##### [ ] Handler
CreatePostHandler \
MainPageHandler \
PostPageHandler \
EditPostHandler \
DeletePostHandler

#### Workflow:
##### Create Post
1. Middleware checks if user is authorized, then handler receives POST request with post title, post text, post category/tags and calls service.CreatePost()
2. service.CreatePost() checks if post already exists using repo.CheckPost(), then calls repo.AddPost() to add post
3. repo.AddPost uses sql command to add post into the database 

##### Get Post
1. Handler receives the GET request for POST and calls service.GetPost() with post id
2. service.GetPost() calls repo.GetPost() to retrieve post and returns it to handler
3. repo.GetPost() executes sql command to get post from db by id and returns it to service as models.Post

##### Edit Post
1. Middleware checks if user is authorized and is author of the post, then handler receives POST request, which contains postId, editedText and userToken and calls service.EditPost(), receives postId and error, if err == nil, redirects the request to the editedPost
2. service.EditPost() calls repo.ChangePostText() with postId and edited text, receives 

#### Comment system
##### [ ] Repo
AddComment()
GetComment()
GetComments()
ChangeComment()
DeleteComment()

##### [ ] Model
Comment
TODO: add comment model here
Id int
PostId int 
Text string
Author models.User

##### [ ] Service
CreateComment()
GetComments()
EditComment()
DeleteComment()

##### [ ] Handler
AddComment()
GetComments()
EditComment()
DeleteComment()

##### [ ] Middleware 
IsAuthorized()
IsCommentAuthor()

#### Workflow 
##### Add comment
1. Middleware checks if user is authorized, then handler receives POST request with postId, commentAuthor, commentText and calls service.CreateComment()
2. service.CreateComment() validates request, then calls repo.GetPost() to get models.Post for comment, then calls repo.GetUser() to get comment author and sends them to repo.AddComment() to add comment to post
3. repo.AddComment() adds comment to the db, then returns the OK status to the service


##### Get Comments
1. Handler receives the GET request with postId and calls service.GetComments() with postId
2. service.GetComments() receives the request, validates it, then calls repo.GetPost() to get models.Post from it, then calls repo.GetComments() to get []models.Comment from it, then returns it to handler
3. repo.GetComments() uses sql commands to get all the comments associated with post, then returns them to service

##### Edit Comment
1. Middleware checks if user is author of the comment, then handler receives the POST request with postId, commentId, commentUpdatedText and calls service.EditComment() with them
2. service.EditComment() calls repo.GetPost() to get models.Post and repo.GetComment() to get models.Comment for the post and sends them to repo.ChangeComment() with updated text
3. repo.ChangeComment() changes the comment text in db, and returns OK 

##### Delete Comment
1. Middleware checks if user is author of the comment, then handler receives the POST request with postId, commentId and calls service.DeleteComment()
2. service.DeleteComment() calls repo.GetComment(), receives models.Comment form it, and sends it to repo.DeleteComment()
3. repo.DeleteComment() deletes the comment from db, then returns OK 

#### Likes/Dislikes
##### [ ] Repo
##### [ ] Model
##### [ ] Service
##### [ ] Handler

#### Tags/Categories
##### [ ] Repo
##### [ ] Model
##### [ ] Service
##### [ ] Handler

#### Filter system
##### [ ] Repo
##### [ ] Model
##### [ ] Service
##### [ ] Handler

### Frontend 
#### [ ] Main Page
#### [ ] Authorization page
#### [ ] Post Page
#### [ ] Search Page
