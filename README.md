# slack-secops-challenge

## API Documentation

### /getwordfreq

#### Authentication
- Basic Authentication which requires username and password

#### Input:
- A POST JSON request with format 
        {
            "input" : "[YOUR SENTENCE HERE]"
        }

#### Output:
- A JSON response with format 
        {
            "count" :   [int: NUMBER OF UNIQUE WORDS IN THE SENTENCE],
            "words" :   {
                "[UNIQUE WORDS #1]" :  [int: NUMBER OF TIMES APPEARED IN THE SENTENCE],
                "[UNIQUE WORDS #2]" :  [int: NUMBER OF TIMES APPEARED IN THE SENTENCE],
                ...
            }
        }

-----

## Feature Requests

### Other languages
- It'd be nice to support other alphabets outside english, e.g. French / Greek / Vietnamese alphabets.
- Might need to exclude Eastern-Asian characters since it's more complicated to decide what forms a word

-----

## Design Decisions

### Using Golang native net/http libraries
- As for a simple API, native libraries are sufficient at this stage
- Also lowers the risks (not eliminate) of vulnerabilities coming from third-parties

### Choosing Basic Auth instead of Digest Auth
- For Digest Auth, I could go with either using a Golang package such as "go-http-auth" or use Nginx with Digest auth. However, both of them lack popularity and seem somewhat obsecure and could potentially be vulnerable in production. 
- Hence my best option for Digest auth would be to implement one myself. However, given that the time limit for this challenge is only one week and I will probably not be able to work on this much during weekdays, using Basic auth for now and migrating the project to Digest auth in the future would be a better option.
- On the other hand, I'm already using HTTPS, Basic auth would be sufficient. It's also more widely supported compare to Digest auth. 

### Not using a CI/CD tool
- I wasn't sure if I'm authorized to provide the SSH key of the server to a third-party (e.g. Bamboo/CircleCI) as this could be considered unsecure. However, that can be done if needed in the future.

### Using Docker
- Because it's Portable, Consistent, Clean, and Friendly to automation tools.
- Can easily setup autostart containers when the server reboots or crashes.
