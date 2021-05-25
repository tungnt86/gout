# gout - A Go unit test kit
gout integrates go-sqlmock from DATA-DOG to mock database activities.
This kit is helpful to test commit/rollback of transactions, which spread on several repositories in sevice/usecase layer.
## Usecases
- Test your transactions, which spread on several repositories in service/usecase layer to make sure a transaction is rollbacked if there is an error.
- Mock database errors, EX: connection timeout.
## Example code
Example code is in `example` folder.