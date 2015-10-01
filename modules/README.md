# Modules

Modules are how this proxy works. The proxy registers modules, and each module implemtents two functions:

* filterRequest() - this filters outgoing requests from hitting the intended target and can have any effect. For example, dropping the request or sending a different one
* filterResponse() - this filters incoming responses before they hit the app and can have any effect. For example, sending an empty response, or a carefully crafted different one.

## Writing a module

To write a module, you create a struct in a new package and implement the Module interface (package main, modules.go). The module itself will be a struct with one field, MetaStuct which contains the metadata. For the sake of consistency, please call it Metadta (case sensitive).
It also helps to create a constructor for the module (i.e. func NewMyModule() \*MyModule).

## Registering modules

To register the module for use within the proxy, call the RegisterModule function in main.go, function main (i.e. RegisterModule(mymodule.NewMyModule()).

