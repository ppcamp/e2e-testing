# playwright-cucumber e2e test

Doing the way that I've did for this repo, you'll be able to spawn a new playwright for each
scenario without impact your current tests.

- with concurrency and no headless: 5.967s
- without concurrency and no headless: 9.103s

> Why to initialize a Playwright window instead of reusing the older one?
>> R.: We want to each scenario be completely independent, avoiding possible errors with cookies
>> and caches.

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