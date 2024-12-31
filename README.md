# Organization Service

The Leagueify organization service is responsible for managing user authentication,
authorization, and profiles. The Leagueify organization service uses 
[Go][go-download] using version `1.23.0`.

## Key Functions

- Create and manage organizations.
- Assign administrators and roles within organizations.
- Manage leagues associated with an organization.

## Development

### Prerequisites

- [Docker][docker-download] is installed and running.

### Getting Started

Local development of the Leagueify organization service uses docker for a
consistent development environment. Running the Leagueify organization service
locally can be accomplished using commands found in the
[Makefile][repo-makefile]. During local development changes will trigger a live
reload, eliminating the requirement to restart the docker image manually. This
is accomplished by using the wonderful tool [air][github-air]. The most common
commands have been outlined below:

```bash
# start leagueify docker image
make start

# stop leagueify docker image removing attached volumes
make clean
```

The Leagueify organization service is ready for development once the banner output is
visible within the terminal. The banner blelow was created using the
[Text to ASCII Art Generator by Patorjk][patorjk-taag].

```
leagueify-organization-1  |
leagueify-organization-1  | '||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|  '||' '||''''| '||' '|'
leagueify-organization-1  |  ||        ||  .       |||    .|'     '   ||    |   ||  .     ||   ||  .     || |
leagueify-organization-1  |  ||        ||''|      |  ||   ||    ....  ||    |   ||''|     ||   ||''|      ||
leagueify-organization-1  |  ||        ||        .''''|.  '|.    ||   ||    |   ||        ||   ||         ||
leagueify-organization-1  | .||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....| .||. .||.       .||.
leagueify-organization-1  |
leagueify-organization-1  |  ..|''||   '||''|.    ..|'''.|      |     '|.   '|' '||' |'''''||      |     |''||''| '||'  ..|''||   '|.   '|'
leagueify-organization-1  | .|'    ||   ||   ||  .|'     '     |||     |'|   |   ||      .|'      |||       ||     ||  .|'    ||   |'|   |
leagueify-organization-1  | ||      ||  ||''|'   ||    ....   |  ||    | '|. |   ||     ||       |  ||      ||     ||  ||      ||  | '|. |
leagueify-organization-1  | '|.     ||  ||   |.  '|.    ||   .''''|.   |   |||   ||   .|'       .''''|.     ||     ||  '|.     ||  |   |||
leagueify-organization-1  |  ''|...|'  .||.  '|'  ''|...'|  .|.  .||. .|.   '|  .||. ||......| .|.  .||.   .||.   .||.  ''|...|'  .|.   '|
leagueify-organization-1  |
```

[docker-download]: https://www.docker.com/get-started/
[github-air]: https://github.com/air-verse/air
[go-download]: https://go.dev/dl/
[patorjk-taag]: https://patorjk.com/software/taag/#p=display&f=Kban&t=LEAGUEIFY%0AORGANIZATION
[repo-makefile]: ./Makefile
