# slack-secops-challenge

## Feature Requests
### Other languages
- It'd be nice to support other alphabets outside english, e.g. French / Greek / Vietnamese alphabets.
- Might need to exclude Eastern-Asian characters since it's more complicated to decide what forms a word

## Design Decisions

### Use Golang native libraries
- As for a simple API, native libraries are sufficient
- Also lowers the risks (not eliminate) of vulnerabilities coming from third-parties

### Use Basic Auth instead of Digest Auth
- For Digest Auth, I could go with either using a Golang package such as "go-http-auth" or use Nginx with Digest auth. However, both of them lack popularity and seem somewhat obsecure and could potentially be vulnerable in production.
- On the other hand, I'm already using HTTPS, Basic auth is widely used compare to Digest and would be sufficient.

### Not using a CI/CD tool
- I wasn't sure if I'm authorized to provide the SSH key of the server to a third-party (e.g. Bamboo/CircleCI) as this could be considered unsecure. However that can be done if needed.

### Using Docker
- Because it's Portable, Consistent, Clean, and Friendly to automation tools.
- Easily setup autostart containers when the server reboots or crashes
