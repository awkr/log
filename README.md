## Note

* logging to file is not compatible with logging to folder.

  logging to file has a higher priority and if setted, logging to folder will be ignored

## TODO

1. [x] color support
2. [x] strip color when log to other destinations, not stdout / stderr
3. [x] log to file
4. [x] log to folder
	* [x] format of file name
	* [x] log to different files according to log level
	* [ ] refine file name
5. [ ] file rotate
	* [ ] by day
	* [ ] by hour
	* [ ] format of file name
6. [ ] file & line number
	* [ ] file name with line number
	* [ ] full file path is **not considered**
