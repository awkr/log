## Note

* logging to file is not compatible with logging to folder.

  logging to file has a higher priority and if setted, logging to folder will be ignored

## TODO

1. [x] color support
2. [x] strip color when log to other destinations, not stdout / stderr
3. [x] log to file
4. [ ] log to folder
	* [ ] format of file name
	* [ ] log to different files according to log level
5. [ ] file rotate
	* [ ] by day
	* [ ] by hour
6. [ ] file & line number
	* [ ] file name with line number
	* [ ] full file path is **not considered**
