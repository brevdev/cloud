# Contributing to Lambda Labs

## Setup

- Create a `.env` file in the root of the project with the following:

```
LAMBDALABS_API_KEY=secret_my-api-key_**********
```

## Running Tests

Use the vscode "run tests" task to run the tests.


## Useful Prompts
```
can you take a look at the file structure in @v1 this is supposed to be the reference/interface for providers. can you replicate the file structure in @lambdalabs ? I just want the file structure and maybe some stubs if they make sense.
```

```
Please analyze the Lambda Labs @lambdalabs  implementation in and research their actual API capabilities to determine which features are supported vs unsupported.

**Required Research Steps:**
1. **API Documentation Analysis**: Find and review Lambda Labs' official API documentation at @https://cloud.lambda.ai/api/v1/openapi.json  to identify supported endpoints and features
2. **Feature-by-Feature Verification**: For each CloudClient interface method, verify if Lambda Labs actually supports it by checking their API docs
3. **Evidence-Based Decisions**: Only mark features as "supported" if you find concrete evidence in their documentation

**Implementation Approach:**
- Use the existing NotImplCloudClient pattern for unsupported features
- Return ErrNotImplemented for features that Lambda Labs doesn't support
- Maintain full CloudClient interface compliance

**Key Questions to Answer:**
- What instance management operations does Lambda Labs actually support? (create, terminate, list, stop/start, reboot?)
- Do they support volume resizing, instance type changing, or machine images?
- What networking features do they provide? (firewall rules, security groups?)
- Do they have quota management APIs?
- What authentication and location management do they support?

**Deliverables:**
1. Evidence-based list of supported vs unsupported features
2. Refactored implementation using NotImplCloudClient
3. Updated documentation reflecting actual capabilities
4. Remove any files/methods that are completely unnecessary

Please provide specific citations from Lambda Labs' API documentation for any features you mark as supported. Make sure we list features that the client implements NOT the API.
```