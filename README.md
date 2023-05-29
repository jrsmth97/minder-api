
# Minder API

![Minder](src/server/static/images/minder.jpg "minder")

Minder is a mini and simple dating apps where you can match and find your life partner. 

## Minimum Requirements

Minder Api shares the same [minimum requirements][] as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

[minimum requirements]: https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements

### Run The Project

Minder requires Go version 1.19 or newer and the Makefile requires GNU make.

On Windows, the makefile requires the use of a bash terminal to support all makefile targets.
An easy option to get bash for windows is using the version that comes with [git for windows](https://gitforwindows.org/).

1. [Install Go](https://golang.org/doc/install)
2. Clone the Telegraf repository:

   ```shell
   git clone https://github.com/jrsmth97/minder-api.git
   ```

3. Rename `.env.example` file to `.env` in root directory
4. Run `make run` from the source directory to run this project

   ```shell
   cd minder-api
   make run
   ```

5. Testing the project

   ```shell
   make test
   ```

## List Of Endpoints

	purchase := r.router.Group("/purchases")
	purchase.POST("/", r.middleware.Auth, r.purchase.CreatePurchase)
	purchase.GET("/:purchaseId", r.middleware.Auth, r.purchase.GetPurchase)
	purchase.POST("/cancel/:purchaseId", r.middleware.Auth, r.purchase.CancelPurchase)
	purchase.POST("/sync", r.middleware.Auth, r.middleware.AdminOnly(r.purchase.SyncPurchase))

Endpoint | Method | Description | Auth | Restrict |
|---|---|---|---|---|---|
| [/auth/register] | POST | User registration | No | - | 
| [/auth/login] | POST | User login | No | - | 
| [/users/me] | GET | Get current login account profie | Yes | - | 
| [/users/explore] | GET | Explore list of account | Yes | - | 
| [/users/me] | PUT | Update current login account profile | Yes | - | 
| [/users/delete] | DELETE | Delete current login account | Yes | - | 
| [/users/count] | GET | Count all registered accounts in system (for testing) | Yes | Admin Only | 
| [/memberships] | GET | Get list of available memberships plan | Yes | - | 
| [/memberships] | POST | Create a membership plan | Yes | Admin Only | 
| [/memberships/:membershipId] | GET | Get membership detail | Yes | - | 
| [/memberships/:membershipId] | PUT | Update membership detail | Yes | Admin Only | 
| [/memberships/:membershipId] | DELETE | Delete a membership | Yes | Admin Only | 
| [/locations] | GET | Get list of available locations | Yes | - | 
| [/locations] | POST | Create a new location | Yes | Admin Only | 
| [/locations/:locationId] | GET | Get location detail | Yes | - | 
| [/locations/:locationId] | PUT | Update location detail | Yes | Admin Only | 
| [/locations/:locationId] | DELETE | Delete a location | Yes | Admin Only | 
| [/media/] | POST | Upload image | Yes | - | 
| [/media/:mediaId] | GET | Retrieve image | Yes | - | 
| [/media/:mediaId] | DELETE | Delete image | Yes | - | 
| [/swipes/like/:targetId] | POST | Perform right swipe (like) | Yes | - | 
| [/swipes/pass/:targetId] | POST | Perform left swipe (pass) | Yes | - | 
| [/swipes/favourite/:targetId] | POST | Perform top swipe (favourite) | Yes | - | 
| [/purchases] | POST | Create a purchase | Yes | - | 
| [/purchases/:purchaseId] | GET | Get purchase detail | Yes | - | 
| [/purchases/cancel/:purchaseId] | POST | Cancel a purchase | Yes | - | 
| [/purchases/sync] | POST | Sync all pending purchases to payment gateway provider | Yes | Admin Only | 

## Project Structure

```
minder-api
├── uploads/     # uploaded files
├── src/         # src files
│   ├── db/          # gorm lib database init
│   ├── helper/      # helper functions
│   ├── seed/        # data seeding for database
│   ├── static/      # static assets file
│   ├── test/        # feature testing files
│   └── server/      # core files
│       ├── controller/     # list of controllers
│       ├── enums/          # list of enums & constants
│       ├── middleware/     # list of middlewares
│       ├── model/          # list of table entities
│       ├── param/          # list of request dto / parameters define
│       ├── pkg/            # helper package services 
│       ├── service/        # list of feature services
│       ├── view/           # list of service response define
│       └── repository/     # repository and it implementations
│           ├── repo_impl/       # list of repository implementations
│           └── repo.go          # defined repositories
├── ... 
└── go.mod
```

## Author

**Fajar Permadi**
* <https://github.com/jrsmth97>

## Copyright and License

copyright 2023 jrsmth97.   
