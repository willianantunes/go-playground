# An honest code challenge to get to know Go!

When I'm about to learn a new programming language, I usually build a system that can be easily compared to real ones. For instance, I like to start with a Tic Tac Toe game playable through the command line and evolve it to an API with a database. This time it will be a bit different.

## The challenge itself

Please consider the file [sample.log](./sample.log) for the following details:

- The service reads an access log file every 5 minutes (it can be customized).
- The access log file should be parsed so that each line can be easily identifiable by someone who doesn't understand it.
- Suppose the HTTP method associated with the line is either GET or POST. In that case, it should be saved in a database to be consulted later.
- A line is considered unique if the following fields are equal: http_cf_connection_ip, time_local, request, and body_bytes_sent. If this situation happens, then it should be informed in any way.
- The service exposes an HTTP API where you can consult what is in the database with pagination.
- The service exposes an HTTP API where you can get the latest lines that have been inserted either with GET or POST. The user chooses which one they want.

## Credits

I wasn't the one who created this challenge. Actually, I saw it on a YouTube channel. It can even be used by some companies out there, but I'm not aware of it.
