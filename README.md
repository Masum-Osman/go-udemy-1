> In Go test, no need to import the module to access it’s functions. Because it is in the same dir and same module, and furthermore a TEST file, it has access to all the function by nature.
> 

> The Test function must named under some rules.
> 
1. The Func should start with Test key word with the Capital T in Test

If the test is passed, then you will get only notified. But if the function failed, then you will the function name in return.