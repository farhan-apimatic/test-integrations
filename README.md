# Getting started

This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Swagger%20Petstore-GoLang&projectName=swaggerpetstore_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=swaggerpetstore_lib)

## How to Use

The following section explains how to use the SwaggerpetstoreLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=swaggerpetstore_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=swaggerpetstore_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=swaggerpetstore_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=swaggerpetstore_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=swaggerpetstore_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Swagger%20Petstore-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=swaggerpetstore_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=swaggerpetstore_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| oAuthAccessToken | OAuth 2.0 Access Token |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [pet_pkg](#pet_pkg)
* [store_pkg](#store_pkg)
* [user_pkg](#user_pkg)

## <a name="pet_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".pet_pkg") pet_pkg

### Get instance

Factory for the ``` PET ``` interface can be accessed from the package pet_pkg.

```go
pet := pet_pkg.NewPET()
```

### <a name="add_pet"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.AddPet") AddPet

> Add a new pet to the store


```go
func (me *PET_IMPL) AddPet(body *models_pkg.Pet)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Pet object that needs to be added to the store |


#### Example Usage

```go
var body *models_pkg.Pet

var result 
result,_ = pet.AddPet(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 405 | Invalid input |



### <a name="update_pet"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.UpdatePet") UpdatePet

> Update an existing pet


```go
func (me *PET_IMPL) UpdatePet(body *models_pkg.Pet)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Pet object that needs to be added to the store |


#### Example Usage

```go
var body *models_pkg.Pet

var result 
result,_ = pet.UpdatePet(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid ID supplied |
| 404 | Pet not found |
| 405 | Validation exception |



### <a name="find_pets_by_status"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.FindPetsByStatus") FindPetsByStatus

> Multiple status values can be provided with comma separated strings


```go
func (me *PET_IMPL) FindPetsByStatus(status []models_pkg.Status2Enum)([]*models_pkg.Pet,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| status |  ``` Required ```  ``` Collection ```  | Status values that need to be considered for filter |


#### Example Usage

```go
status := []models_pkg.Status2Enum{ models_pkg.status2_AVAILABLE }

var result []*models_pkg.Pet
result,_ = pet.FindPetsByStatus(status)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid status value |



### <a name="find_pets_by_tags"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.FindPetsByTags") FindPetsByTags

> Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.


```go
func (me *PET_IMPL) FindPetsByTags(tags []string)([]*models_pkg.Pet,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| tags |  ``` Required ```  ``` Collection ```  | Tags to filter by |


#### Example Usage

```go
tags := []string{"tags"}

var result []*models_pkg.Pet
result,_ = pet.FindPetsByTags(tags)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid tag value |



### <a name="get_pet_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.GetPetById") GetPetById

> Returns a single pet


```go
func (me *PET_IMPL) GetPetById(petId int64)(*models_pkg.Pet,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| petId |  ``` Required ```  | ID of pet to return |


#### Example Usage

```go
petId,_ := strconv.ParseInt("252", 10, 8)

var result *models_pkg.Pet
result,_ = pet.GetPetById(petId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid ID supplied |
| 404 | Pet not found |



### <a name="update_pet_with_form"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.UpdatePetWithForm") UpdatePetWithForm

> Updates a pet in the store with form data


```go
func (me *PET_IMPL) UpdatePetWithForm(
            petId int64,
            name *string,
            status *string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| petId |  ``` Required ```  | ID of pet that needs to be updated |
| name |  ``` Optional ```  | Updated name of the pet |
| status |  ``` Optional ```  | Updated status of the pet |


#### Example Usage

```go
petId,_ := strconv.ParseInt("252", 10, 8)
name := "name"
status := "status"

var result 
result,_ = pet.UpdatePetWithForm(petId, name, status)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 405 | Invalid input |



### <a name="delete_pet"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.DeletePet") DeletePet

> Deletes a pet


```go
func (me *PET_IMPL) DeletePet(
            petId int64,
            apiKey *string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| petId |  ``` Required ```  | Pet id to delete |
| apiKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
petId,_ := strconv.ParseInt("252", 10, 8)
apiKey := "api_key"

var result 
result,_ = pet.DeletePet(petId, apiKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid ID supplied |
| 404 | Pet not found |



### <a name="upload_file"></a>![Method: ](https://apidocs.io/img/method.png ".pet_pkg.UploadFile") UploadFile

> uploads an image


```go
func (me *PET_IMPL) UploadFile(
            petId int64,
            additionalMetadata *string,
            file []byte)(*models_pkg.ApiResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| petId |  ``` Required ```  | ID of pet to update |
| additionalMetadata |  ``` Optional ```  | Additional data to pass to server |
| file |  ``` Optional ```  | file to upload |


#### Example Usage

```go
petId,_ := strconv.ParseInt("252", 10, 8)
additionalMetadata := "additionalMetadata"
file :=  []byte("")

var result *models_pkg.ApiResponse
result,_ = pet.UploadFile(petId, additionalMetadata, file)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="store_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".store_pkg") store_pkg

### Get instance

Factory for the ``` STORE ``` interface can be accessed from the package store_pkg.

```go
store := store_pkg.NewSTORE()
```

### <a name="get_inventory"></a>![Method: ](https://apidocs.io/img/method.png ".store_pkg.GetInventory") GetInventory

> Returns a map of status codes to quantities


```go
func (me *STORE_IMPL) GetInventory()(map[string]int64,error)
```

#### Example Usage

```go

var result map[string]int64
result,_ = store.GetInventory()

```


### <a name="create_place_order"></a>![Method: ](https://apidocs.io/img/method.png ".store_pkg.CreatePlaceOrder") CreatePlaceOrder

> Place an order for a pet


```go
func (me *STORE_IMPL) CreatePlaceOrder(body *models_pkg.Order)(*models_pkg.Order,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | order placed for purchasing the pet |


#### Example Usage

```go
var body *models_pkg.Order

var result *models_pkg.Order
result,_ = store.CreatePlaceOrder(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid Order |



### <a name="get_order_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".store_pkg.GetOrderById") GetOrderById

> For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions


```go
func (me *STORE_IMPL) GetOrderById(orderId int64)(*models_pkg.Order,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | ID of pet that needs to be fetched |


#### Example Usage

```go
orderId,_ := strconv.ParseInt("252", 10, 8)

var result *models_pkg.Order
result,_ = store.GetOrderById(orderId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid ID supplied |
| 404 | Order not found |



### <a name="delete_order"></a>![Method: ](https://apidocs.io/img/method.png ".store_pkg.DeleteOrder") DeleteOrder

> For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors


```go
func (me *STORE_IMPL) DeleteOrder(orderId int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | ID of the order that needs to be deleted |


#### Example Usage

```go
orderId,_ := strconv.ParseInt("252", 10, 8)

var result 
result,_ = store.DeleteOrder(orderId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid ID supplied |
| 404 | Order not found |



[Back to List of Controllers](#list_of_controllers)

## <a name="user_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".user_pkg") user_pkg

### Get instance

Factory for the ``` USER ``` interface can be accessed from the package user_pkg.

```go
user := user_pkg.NewUSER()
```

### <a name="create_user"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.CreateUser") CreateUser

> This can only be done by the logged in user.


```go
func (me *USER_IMPL) CreateUser(body *models_pkg.User)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Created user object |


#### Example Usage

```go
var body *models_pkg.User

var result 
result,_ = user.CreateUser(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 0 | successful operation |



### <a name="create_users_with_array_input"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.CreateUsersWithArrayInput") CreateUsersWithArrayInput

> Creates list of users with given input array


```go
func (me *USER_IMPL) CreateUsersWithArrayInput(body []*models_pkg.User)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  ``` Collection ```  | List of user object |


#### Example Usage

```go
var body []*models_pkg.User

var result 
result,_ = user.CreateUsersWithArrayInput(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 0 | successful operation |



### <a name="create_users_with_list_input"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.CreateUsersWithListInput") CreateUsersWithListInput

> Creates list of users with given input array


```go
func (me *USER_IMPL) CreateUsersWithListInput(body []*models_pkg.User)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  ``` Collection ```  | List of user object |


#### Example Usage

```go
var body []*models_pkg.User

var result 
result,_ = user.CreateUsersWithListInput(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 0 | successful operation |



### <a name="get_login_user"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.GetLoginUser") GetLoginUser

> Logs user into the system


```go
func (me *USER_IMPL) GetLoginUser(
            username string,
            password string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| username |  ``` Required ```  | The user name for login |
| password |  ``` Required ```  | The password for login in clear text |


#### Example Usage

```go
username := "username"
password := "password"

var result string
result,_ = user.GetLoginUser(username, password)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid username/password supplied |



### <a name="get_logout_user"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.GetLogoutUser") GetLogoutUser

> Logs out current logged in user session


```go
func (me *USER_IMPL) GetLogoutUser()(,error)
```

#### Example Usage

```go

var result 
result,_ = user.GetLogoutUser()

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 0 | successful operation |



### <a name="get_user_by_name"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.GetUserByName") GetUserByName

> Get user by user name


```go
func (me *USER_IMPL) GetUserByName(username string)(*models_pkg.User,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| username |  ``` Required ```  | The name that needs to be fetched. Use user1 for testing. |


#### Example Usage

```go
username := "username"

var result *models_pkg.User
result,_ = user.GetUserByName(username)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid username supplied |
| 404 | User not found |



### <a name="update_user"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.UpdateUser") UpdateUser

> This can only be done by the logged in user.


```go
func (me *USER_IMPL) UpdateUser(
            username string,
            body *models_pkg.User)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| username |  ``` Required ```  | name that need to be updated |
| body |  ``` Required ```  | Updated user object |


#### Example Usage

```go
username := "username"
var body *models_pkg.User

var result 
result,_ = user.UpdateUser(username, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid user supplied |
| 404 | User not found |



### <a name="delete_user"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.DeleteUser") DeleteUser

> This can only be done by the logged in user.


```go
func (me *USER_IMPL) DeleteUser(username string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| username |  ``` Required ```  | The name that needs to be deleted |


#### Example Usage

```go
username := "username"

var result 
result,_ = user.DeleteUser(username)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid username supplied |
| 404 | User not found |



[Back to List of Controllers](#list_of_controllers)



