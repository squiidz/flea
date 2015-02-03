flea
====

## What is flea ? 

Flea is a tiny and simple project builder for Go web server.

![alt tag](http://upload.wikimedia.org/wikipedia/commons/6/68/Punaise.jpg)

## Example

1. ``` flea new projectName ``` create the project folder and config.json file. 
2. ``` flea install ``` install the dependencies found in the config.json file.
3. ``` flea save yourCommitMessage ``` add all file who have change, commit them and push it to your git server.


## Project Structure 
```
-- ProjectFolder
           	|-- public
          		|-- template
                	|-- index.html
                |-- js
                |-- css
                |-- fonts
          |-- app.go
          |-- config.json
```
## TODO

- Add fonctionality 
- Debug
- Refactoring

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## License

MIT
