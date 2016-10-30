# Go JSON File Handling

This package provides simple wrapping around saving and reading JSON files (streams) in GO. It exists because way too many sample/demo code are doing it the wrong way. Wrong in the sense of efficiency -- when reading JSON from file for example, the right is to read from stream, instead of reading the whole file via `ioutil.ReadAll` into a big buffer first then decode the buffer next.

Every single hit I searched, for "JSON files reading writing", is doing reading like this. They are actually not wrong, but just not ideal, especially when it comes to big JSON files.

Hence, this simple JSON files reading writing wrapping package. Here is a [simple demo code](https://github.com/suntong/lang/blob/master/lang/Go/src/ds/jsonfile.go).

