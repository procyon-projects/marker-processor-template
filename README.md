# marker-processor-template
In order to create a new marker processor, either clone this project or use **Marker CLI**.

The simple rules that you need to follow if you create a new marker processor by cloning this repository are:
* Rename **processor-name** directory with your marker name.
* Replace **AppName** constant value with your marker name. 
* Replace **Module Name** in go.mod file with your module name.
## Marker CLI Installation
You can install Marker CLI by using the following command.

```shell
go get -u github.com/procyon-projects/marker/...
```

Then use the following command to initialize a new marker processor project.

```shell
marker init [name] [module]
```

## How to develop a new marker processor
To learn how to develop a new processor, please check out [marker project](https://github.com/procyon-projects/marker).

## Third-Party Libraries
We use some third-party libraries while creating marker processor template. Here is the list we’ve used :
*  [cobra](https://github.com/spf13/cobra)

## License
Marker Processor Template is released under version 2.0 of the Apache License
