# GitGo
A minimal CLI app to query GitHub users.

This project was conceived to introduce GO to beginners.

## Installation
- `go get github.com/pmbenjamin/gitgo`

## Usage
- Help: 
    ```
    $ gitgo
    Usage: gitgo [options]
    Options:
        -u, --user string
            Search Users
    ```
- Single User Query:
    ```
    $ gitgo -u pmbenjamin
    Searching user(s): [pmbenjamin]
    Username:        pmbenjamin
    Name:            Peter Benjamin
    Email:           petermbenjamin@gmail.com
    Bio:             Software Engineer. Hacker. Passionate about coding, IoT, security, and teaching kids how to code.
    ```
- Multi-User Query:
    ```
    $ gitgo -u pmbenjamin,defunkt                                                                                                                                    
    Searching user(s): [pmbenjamin defunkt]
    Username:        pmbenjamin
    Name:            Peter Benjamin
    Email:           petermbenjamin@gmail.com
    Bio:             Software Engineer. Hacker. Passionate about coding, IoT, security, and teaching kids how to code.
    
    Username:        defunkt
    Name:            Chris Wanstrath
    Email:           chris@github.com
    Bio:             🍔 
    ```

## Contributing
Feel free to open a GitHub Issue or submit a PR with your features.