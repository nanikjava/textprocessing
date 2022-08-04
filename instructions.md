Rokt Backend Take Home
===

The Rokt Backend Take Home is intended to test your abilities at building portable,
correct, and readable code within a modern development ecosystem. The project will
describe a simple business problem and expect you to solve it to your best ability.
Great solutions are easily built, include necessary tests and documentation, solve
the problem accurately, and solve the problem performantly.

## Overview

A company has a business process that requires processing of text files from external
customers. Some information about the text file:

- Each line of the text file is an individual record and can be processed independently of the other records in the text file
- The size of the file can range from a few KB to a few GB
- The file will be ordered by iso8601 time ascending
- Sample files can be found here: https://github.com/ROKT/backend-technical-test

Here is the format of each line:

```
[Date in YYYY-MM-DDThh:mm:ssZ Format][space][Email Address][space][Session Id in GUID format]
```

For example:

```
2020-12-04T11:14:23Z jane.doe@email.com 2f31eb2c-a735-4c91-a122-b3851bc87355
```

## Requirements

Please write an application to parse the text files above and expose an HTTP API to serve the files.

### Input

The application *must* follow the following standard *exactly*.

- The program must run on port 8279
- The program must expose an HTTP POST route on the `/` path
- The from/to options must support iso8601 UTC timestamps *only*
- The program must accept the POST body as defined below
- The files must be loaded from `/app/test-files`

```bash
curl -XPOST localhost:8279/ -d '{"filename":"sample1.txt", "from":"2020-07-06T23:00:00Z", "to": "2021-07-06T23:00:00Z"}'
```

### Output

Your output must be parsed entries within the date time range *inclusively*, in JSON
format.

A successful response should return content-type `application/json`, an HTTP status
code of 200, and a response body of the following format:

```JSON
[
  {
    "eventTime":"2000-01-01T03:05:58Z",
    "email":"test123@test.com",
    "sessionId":"97994694-ea5c-4da7-a4bb-d6423321ccd0"
  },
  {
    "eventTime":"2000-01-01T04:05:58Z",
    "email":"test456@test.com",
    "sessionId":"97994694-ea5c-4da7-a4bb-d6423321ccd1"
  }
]
```
- The array *must* be ordered by eventTime from earliest to latest
- The ordering of the keys within the JSON objects does not matter

If the input is invalid, or there are no valid entries, your application should return
a 200 HTTP status response code, a content-type of `application/json`, and the following
response body:

```JSON
[]
```

### Building and Running your Application

Included in this zip file is a `Dockerfile`, a `build.sh`, and a `run.sh`.

Please update the `Dockerfile` as necessary to build your application. You are free
to use any parent image. The expectation is that your application will start an HTTP
server on port 8279. The `run.sh` script will expose this on the host machine.

Prior to testing your application each script will be run once in order.

The first argument to `run.sh` will always be the absolute path of the folder containing
the test files. This folder will be mounted into your container.  Files in this folder
will be past to your program _by filename only_.

The `run.sh` script mounts the given filepath into `/app/test-files`. This folder must exist
for the mount to work. The given sample `Dockerfile` ensures this.

### README.md

In this zip file is an empty README.md. Please fill this out describing your project,
your test methodology, embellishments, and places where you can improve.

## Submission

Please submit your take home as a compressed zip file to this google form:

[Rokt Backend Takehome Submission](https://docs.google.com/forms/d/1FdP41PWZ8cJzNQoU7HEonWUrtR9Ef12iXtIhbnXnnb4/)

## Notes and Hints

Some notes about the challenge:

- Please write the application in a language that you are most comfortable with.
- The submitted code should be production-grade code. Your submission will be judged on code quality as well as correctness
- You should include testing code
- The solution should handle failure cases
- Package the final code in a zip file. Please do not include any binaries files in the zip, e.g. libraries, compiled code...

Couple of hints on how to proceed (those are just hints, please go as far as you like):

- Make it work first!
- Think about edge cases
- Think about how to write the application with constant memory consumption regardless of file size
- Run through your application with the provided test file