# Server

## Routes Documentation

### Users

#### Signup User

- **Endpoint:** `/user/signup`
- **Method:** `POST`
- **Description:** Signup a new user and store details in the database.

##### Request
Takes a [UserRequest](https://learn.reboot01.com/git/nradhi/forum/src/branch/master/internal/server/README.md#UserRequest)
##### Response
Probaby a UserSession
##### Response Codes
* [201 Created](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201): All Good.
* [400 Bad Request](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400): Your Payload is wrong.
* [406 Not Acceptable](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406): Either username or email exists, or password is bad, response will have that information.
```json
{"error": "Bad Password"}
```
* [413 Content Too Large](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413): Your payload is too big ;).

#### Login User

- **Endpoint:** `/user/login`
- **Method:** `POST`
- **Description:** Logs in a new user

##### Request
Takes a [UserRequest](https://learn.reboot01.com/git/nradhi/forum/src/branch/master/internal/server/README.md#UserRequest)
##### Response
Probaby a UserSession
##### Response Codes
* [201 Created](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201): All Good.
* [400 Bad Request](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400): Your Payload is wrong.
* [406 Not Acceptable](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406): Invalid password or username
```json
{"error": "Wrong Credentials"}
```
* [413 Content Too Large](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413): Your payload is too big ;).

#### Get User Profile

- **Endpoint:** `/user/{username}/profile`
- **Method:** `GET`
- **Description:** Gets the profile of a user

##### Parameters
- `{username}` (required): The username of the user
##### Response
Probaby a User struct
##### Response Codes
* [200 OK](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200): All Good.
* [404 Bad Request](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404): user not found
* [413 Content Too Large](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413): Your payload is too big ;).