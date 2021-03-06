# Go Testing

# Conventions
-  All of our tests must have the format <filename>_test.go
- Test files should be in the same package as your code.
- In Go all tests go inside a file named similar to the file you are testing. So if you want to test code in main.go, then you write the tests in main_test.go
- In the test file we will always need to import the testing package. (Build-in package)
- There is the following signature for every test you write, func TestXxx(*testing.T), where the first X MUST be Uppercase.
- Calls t.Error or t.Fail to indicate a failure (I called t.Errorf to provide more details)
- t.Log can be used to provide non-failing debug information

# Some useful commands 
## Test commands 
  ```go
- go test
- go test –v 
- go test –run  "TestName"
  ```

## Coverage commands &  Visualization 
  ```go
- go test –cover
- go test –cover –coverprofile=c.out
- tool cover –html=c.out –o coverage.html
  ```

## Unit Testing
### Basic Test
- Log or Logf :  Write documentation out into the log output.
- Error or Errorf: Write documentation and also say that this test is failed but we are continue moving forward to execute code in this test.
- Fatal or Fatalf : Similarly, document that this test is failed but we are done. We move on to the next test function.
- ##### "Given, When, Should" format
- Given: Why are we writing this test?
- When: What data are we using for this test?
- Should: When are we expected to see it happen or not happen?

### Table Test
- Set up a data structure of input to expected output.
- This way we don't need a separate function for each one of these. We just have 1 test function.
- As we go along, we just add more to the table.

### Sub Test 
- Sub test helps us streamline our test functions, filters out command-line level big tests into smaller sub tests.

### Assertion/Mocking
- stretchr/testify  one of the most popular go testing library (link: https://github.com/stretchr/testify)
- to install : 
  ```go
    go get github.com/stretchr/testify
  ```
- easy assertion
    ```go
    -> assert.Equal(t, Calculate(2), 4)
    -> assert.NotEqual(t, Calculate(2), 5)
    -> assert.Nil(status)
    -> assert.NotNil(object)
    ```
- mockery provides the ability to easily generate mocks for golang interfaces using the stretchr/testify/mock package. 
    -> to install : 
    ```go
     go get github.com/vektra/mockery/.../
    ```
    -> to generate mock : 
    ```bash
     mockery -name "<interfaceToMock>"
    ```

### Ginkgo
- Ginkgo builds on Go's testing package, allowing expressive Behavior-Driven Development ("BDD") style tests. 
- It is typically (and optionally) paired with the Gomega matcher library.
    -> to install ginkgo & gomega: 
    ```go
     go get github.com/onsi/ginkgo/ginkgo
     go get github.com/onsi/gomega/...
    ```
- General format:
```go
        Describe("the strings package", func() {
        Context("strings.Contains()", func() {
            When("the string contains the substring in the middle", func() {
            It("returns `true`", func() {
                Expect(strings.Contains("Ginkgo is awesome", "is")).To(BeTrue())
            })
            })
        })
        })
```