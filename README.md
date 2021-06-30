Golang brainfuck machine (GBFM)
==============================================

[Brainfuck](http://esolangs.org/wiki/brainfuck) is an esoteric programming language with 8 very simple instructions.

The purpose of GBFM is to interpret the brainfuck code or transpile it into other target programming languages.

Limitations:

* :wink: now only Javascript target language is supported!
* :smirk: Javascript transpiled program don't support ',' (read char from stdin) brainfuck operator.

How to use
-----

1. Run GBFM in interpretation mode use cmd:

    ```bash
    ./gbfm.sh run ./test_data/hello_world.bf
    ```

2. Run GBFM in transpiler mode use cmd:

    ```bash
    ./gbfm.sh translate ./test_data/hello_world.bf
    ```

    Generate ./test_data/hello_world.bf.js in result
    To run *.bf.js file use nodejs

    ```bash
    node ./test_data/hello_world.bf.js
    ```

Testing
-----

Prerequisites for tests running:

* local nodejs installation!

To run all tests use cmd:

```bash
go test -v ./...
```

Licence
----

Code is released under the MIT [license](/LICENSE).

Contribute
----

Pull requests are very welcome!
