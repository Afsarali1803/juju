# Test Suite

The following package contains the integration test suite for Juju.

Tests are structured into test suites. Each suite contains a root task (akin
to a package test) that will setup and run each individual test.

To help break tests down, each test can have a number of subtests. Subtests
are meant for indivdual units of work, without having to bootstrap a controller
for every test. Each subtest will just `ensure` that it does have one, failure
to find a suitable controller, it will create one for you.

### Example of a test suite

```sh
/suites/deploy/                    # test-suite
               task.sh             # root/package test (setup)
               deploy_bundles.sh   # tests
```

### Example of a test

```sh
run_deploy_bundles() {             # Subtest

}

test_deploy_bundles() {            # Test
    run "run_deploy_bundles"       # Run subtest
}
```

## Exit codes / Success

All tests will run through until the end of a test/subtest, unless it encounters
a none zero exit code. In otherwards if you want to assert something passes,
ensure that the command returns `exit 0`. Failure can then be detected of the
inverse.

```sh
echo "passes" | grep -q "passes"   # passes
echo "failed" | grep -q "passes"   # fails
```

## Getting started

To get started, it's best to quickly look at the help command from the runner.

```sh
cd tests && ./main.sh -h
```

Running a full sweep of the integration tests (which will take a long time), can
be done by running:

```sh
cd test && ./main.sh
```

### Running tests

To run the tests, they can be broken down into steps:

```sh
./main.sh deploy                     # Runs deploy test suite
./main.sh deploy test_deploy_bundles # Runs test (and all of the subtests)
./main.sh deploy run_deploy_bundle   # Runs subtest
```

Note: running subtests, will also invoke the parent test to ensure that it has
everything setup correctly.

Running `./main.sh deploy run_deploy_bundle` will also run `test_deploy_bundles`,
but no other subtests, just `run_deploy_bundle`.