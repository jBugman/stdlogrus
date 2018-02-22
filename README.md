## stdlogrus

Package provides some glue for using logrus from std logger.

It is based on code from https://github.com/sirupsen/logrus/issues/118#issuecomment-345475880

Simplest use would be to just call in `main`:

    stdlogrus.Setup()
