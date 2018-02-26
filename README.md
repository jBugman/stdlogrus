## stdlogrus

Package provides some glue for using logrus from std logger.
Main use case is ability to use `logrus` with `net/http.ReverseProxy` which accepts hardcoded std `Logger` struct.

It is based on code from https://github.com/sirupsen/logrus/issues/118#issuecomment-345475880

Basic setup is just call in `main`:

    stdlogrus.Setup()
