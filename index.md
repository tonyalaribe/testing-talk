## Levels of Testing in Go; From unit to end to end
Or, how I think about testing in 2021.

It goes without saying these days, that it is important to test our applications. We could manually test the features we build, but there is increasing consensus amongst developers, that these tests should be written out in code. There are many advantages to actually writing tests:
1. The tests can serve as documentation for features we create. A user looking at the tests can see how a given function or logic is useds.
2. The tests can help us know that the feature we built is actually working as expected.
3. The tests serve as a beacon to let us know when we break existing functionality. This commonly happens when refactoring or extending existing code. We might easily change something which was important, and the tests would be expected to break in this scenario, drawing our attention to this issue.

In general, there are multiple types of tests, usually with overlapping definitions, depending on who you ask. But in the context of "myself", working in small teams building microservices or even just building complete end to end apps in Go, I would limit the types of tests I discuss in this article, to the tests I and my teams actually write. So this article does not assume to cover every kind of tests. Most notably tests like acceptance tests and others.

Kinds of tests I actually write:
- Unit tests
- Integration tests
- End to end tests

### Unit tests:
I love unit tests. They are usually self contained tests, that test singular functionality separately from anything else. In Go, these are fast to run, and pretty much should rarely change after they have been written. If you assume the tests as a pyramid, then unit tests would be at the bottom and the largest group. 

When I build apps, I think of my applications as a composition of multiple little/tiny units. It is these tiny units that I test, individually.

For the purpose of this article, we would imagine a geocoding service. This geocoding service recieves an authenticated request, then

Request -> 
Auth(req) ?-> 
RequestToContextObj(req) -> 
GetUserLocFromDB(ctxObj) -> 
GetLocationFromGeocodeProvider(ctxObj, userLocData) -> 
BuildResponseObj(GeocodeResult) -> 
ResponsetoClient 

Let's imagine this chain of function calls as the entirity of our
service/endpoint. 

For unit test purposes, I could unit test every part of the code except those
which call the database or third party providers. 
I would unit test the request
to context object, where I parse and validate a request to build an internal
object which represents the request. 
I would unit test the BuildResponseObject(GeocodeResult), which constructs the
geocode provider result into the format which i return to the clients
I would also unit test the building of the geocode provider request, to be sure
that these requests are performed correctly. 

Go provides really easy to use tools for all of these kinds of testing, and that
is precisely what i would be showing next.



