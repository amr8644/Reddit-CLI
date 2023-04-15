
# Reddit OSINT

Search any user or subreddit using the terminal


## Authors

- [@amr_ashebo](https://twitter.com/ashebo_amr)


## Run locally

Clone it from the repo

```bash
 git clone https://github.com/amr8644/Reddit-CLI.git
```

To build the project
```
make build
````
    
## Usage/Examples

```
Usage of ./main:
  -limit string
        the maximum number of items desired (default: 25, maximum: 100) (default "25")
  -sort string
        one of (hot, new, top, controversial) (default "hot")
  -subreddit string
        Subreddit you want search (default "golang")
  -time string
        one of (hour, day, week, month, year, all) (default "week")
  -type string
        one of(user,subreddit) (default "user")
  -user string
        Username you want search (default "username")
  -where string
        one of (upvoted,downvoted,overview,submitted) (default "about")
```


## License

[MIT](https://choosealicense.com/licenses/mit/)


