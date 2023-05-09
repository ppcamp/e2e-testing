# playwright-cucumber e2e test

Doing the way that I've did for this repo, you'll be able to spawn a new playwright for each
scenario without impact your current tests.

> Why to initialize a Playwright window instead of reusing the older one?
>> R.: We want to each scenario be completely independent, avoiding possible errors with cookies
>> and caches.


Unfortunately, this approach crop some part of the godog implementation. In the original godog
step function, the closure can have several parameters, which one of those can be of a different type. Of course, you can always use reflection and Generic to keep the original features, although, I'm not going to do that.

```go
// [ORIGINAL]
// ctx.Step(`^I have ([\d]+) ([\w]+) at ([0-9.]+)`)
func iDoSomething(ctx context.Context, groupA int, groupB string, groupC float32) error {}

// meanwhile, this package allows to have just a single group
```


## How to run?

Type `make` or `make help` to see the available commands.


## Dependencies

Installing everything

```bash
make install
```

If you have some problem, take a look in the [playwright issues](./Playwright.md)


## TODOs

-  check for the regexes to allow passing a variable from the feature file and using the callbackClosure


## Output

As you may notice below, one problem of this approach is the unsync stdio, if that output matter, you should sync this using a mutex, which may slow down a little bit.

<div align="center">
  <img src="https://user-images.githubusercontent.com/38117637/236939205-76ca907e-d054-4dda-a38e-196c31d08848.png" align="center" alt="comparative" aria="simple image with the demonstrative"/>
</div>
