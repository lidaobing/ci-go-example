package intpkg //same package name as source file

import (
    "testing" //import go package for testing related functionality
    )

func Test_Add2Ints_1(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
    if (Add2Ints(3, 4) != 7) { //try a unit test on function
        t.Error("Add2Ints did not work as expected.") // log error if it did not work as expected
    } else {
        t.Log("one test passed.") // log some info if you want
    }
}

func Test_Add2Ints_2(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
    t.Log("this is passed") //Indicate that this test failed and log the string as info
}
