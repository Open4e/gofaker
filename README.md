# gofaker

A port of Perl's Data::Faker library that generates fake data.

***Note: I have only just started working on this (driven by some work requirements). Only the names stuff is implemented so far.***

## Installation

    go get github.com/4eek/gofaker

or

    cd $GOPATH/src
    git clone git://github.com/gofaker/gofaker.git
    cd gofaker
    go install

## Usage

    import "github.com/4eek/gofaker"

###Names:

    gofaker.Name()      // Kevin Fourie
    gofaker.FirstName() // John
    gofaker.LastName()  // Smith
    gofaker.Prefix()    // Miss
    gofaker.LastName()  // PhD

More coming soon.

## Run the tests and benchmarks

Just the tests...

    go test

...and with benchmarks...

    go test -test.bench=.

## Performance

    10000 Usernames generated in 19.675ms (0.019675)

## Examples

    cd examples
    go run gofakeit.go

## Contributors

* Kevin Fourie ( http://github.com/4eek/gofaker ).

## TODO

* Finish the actual porting ;)
* Docs

## Note on Patches/Pull Requests

* Fork the project.
* Make your feature addition or bug fix.
* Add tests for it. This is important so I don't break it in a future version unintentionally.
* Commit, do not mess with makefiles, version, or history.
  (if you want to have your own version, that is fine but bump version in a commit by itself I can ignore when I pull)
* Send me a pull request. Bonus points for topic branches.

## Copyright

Copyright (c) 2012 Kevin Fourie. See LICENSE for details.

