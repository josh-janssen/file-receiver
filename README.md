# File Receiver

**File Receiver** is a small http server that accepts incoming files and saves them to a directory.

## Test with Curl

```bash
$ curl http://localhost:8080/upload -F "file=@text.txt" -vvv
```

## License

This project is under the GPLv3 license. See the LICENSE file for the full license text.