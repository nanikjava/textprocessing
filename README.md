### CAUTION

The docker build is taking longer than normal because the application is statically compiled, along with sqlite3 which is needed for the application.

### High level design

Following are some information about the design:

* The docker run command will mount the local directory containing the files as /app/test-files folder. All files will be read and if the file content cannot be 
processed it will be skipped.
* The files are processed in the background when the application starts up. This is design this way to ensure that the application starts up
as soon as possible without delay and able to process user's request. 
* Another reason for the file processing done as background task is because there is no guarantee when the file processing will be complete 
as it can process a lot of big files, so doing it on-the-fly does not make sense.
* The data are stored inside Sqlite in-memory, this is to make it easier to query data and since it is memory it is faster.
* The table are designed with all the fields including the filename that the data was obtained from, to make it easy to query based
on user's parameter.

### Improvement

* The Sqlite database is in-memory which means there is a possibility of consuming lots of memory if the data stored are big, this
require further testing to ensure how big can it goes before it create issue.
* External storage can be introduced for the data by using Redis or NoSql, as both provide query based on dates.
* Database error handling could be improved and reported better.
* Docker build could be optimised by doing multi-docker build, where the application can be hosted inside distroless or scratch image
* More test cases must be added to test database and happy days and non-happy-days handler use case
