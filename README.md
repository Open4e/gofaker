> # PLEASE NOTE: Incomplete and unmaintained.
> Should you want to revive this project please message me (@4eek).

# gofaker [![Build Status](https://travis-ci.org/4eek/gofaker.svg?branch=master)](https://travis-ci.org/4eek/gofaker) [![Build Status](https://drone.io/github.com/4eek/gofaker/status.png)](https://drone.io/github.com/4eek/gofaker/latest) [![Coverage Status](https://coveralls.io/repos/4eek/gofaker/badge.svg)](https://coveralls.io/r/4eek/gofaker)

A port of Perl's Data::Faker library that generates fake data.
Heavily inspired by https://github.com/EmmanuelOga/ffaker and https://github.com/fzaninotto/Faker

***Note: I have only just started working on this (driven by some work requirements). Only the names and lorem stuff is implemented so far.***

## Installation

    go get github.com/4eek/gofaker

or

    cd $GOPATH/src
    git clone git://github.com/gofaker/gofaker.git
    cd gofaker
    go install

## Usage

    import "github.com/4eek/gofaker/person"
    import "github.com/4eek/gofaker/lorem"

###Names:

    person.FullName()  // Kevin Fourie
    person.FirstName() // John
    person.LastName()  // Smith
    person.Prefix()    // Miss
    person.LastName()  // PhD

###Lorem:

    lorem.Word()
    // culpa
    lorem.Words(5)
    // autem sint sit quo aspernatur
    lorem.Sentence(10)
    // Eum culpa neque iste modi aut adipisci ipsam quam rem. 
    lorem.Sentences(10)
    // Omnis pariatur commodi voluptatum aliquid animi. Sunt qui asperiores quia nemo nesciunt. Iste dict
    // a aperiam numquam natus ea. Nemo qui et earum quibusdam sit. Molestiae eius accusantium repellat v
    // eritatis vero. Commodi commodi rem accusantium qui deserunt. Sint ea minima consectetur eius aliqu
    // am. Exercitationem et voluptatibus ut voluptatem quia. A accusamus fugit exercitationem dolor quo.
    //  Laboriosam earum sint quis debitis est.
    lorem.Paragraph(10)
    // Dolorum eos omnis sed possimus impedit. Odit nihil veniam eos blanditiis molestiae. Ratione ex per
    // ferendis mollitia reiciendis eum. Voluptate neque tenetur temporibus et sit. Quos rerum qui exerci
    // tationem quo repellat. Voluptatem magnam repudiandae repellendus blanditiis esse. Nihil similique 
    // consequatur est aliquam accusantium. Non reiciendis dolor et voluptas hic. Asperiores rerum provid
    // ent expedita esse dignissimos. Et aut et quis voluptatem aut.
    lorem.Paragraphs(3)
    // Quisquam fugit quia et corrupti sed. Perferendis sunt autem totam delectus rem. Totam reprehenderi
    // t occaecati quo omnis incidunt. Dolorum et expedita eos nihil reiciendis. Quia voluptas ad quas te
    // mporibus velit. Error natus qui vel deleniti quis.
    //
    // Quod quidem neque exercitationem nostrum recusandae. Est dignissimos nulla possimus qui impedit. N
    // ihil corrupti quo aspernatur rerum deserunt. Minus eos suscipit repudiandae facilis nostrum. Volup
    // tates autem nobis et perspiciatis facilis. Qui facere quos veritatis repellendus dolorem.
    //
    // Facilis voluptates odio id alias et. Corporis quas aspernatur nobis dolorem cumque. Voluptatem vit
    // ae ut quia qui eius. Et quo tenetur hic magni blanditiis. Vel impedit et nihil numquam veritatis. 
    // Quia soluta et aut error ab.

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

