/*Package parse provides a parser for go files that produces a swagger spec document.

You give it a main file and it will parse all the files that are required by that main
package to produce a swagger specification.  This parser doesn't function standalone though; it needs you to provide it
with a template swagger document.

It seems weird to me to have to write doc comments to produce JSON when the only purpose of those comments is to provide that JSON.
So instead of duplicating work it can take a swagger spec to start from and then
synchronize the content of the specs definitions and parameters, with whatever is in your code.
This tool takes the stance that if you want to use it in the swagger spec then you should be able to define it as a struct.
Whenever types are involved I want the compiler to actually agree that it is that type and use that information in the produced json.

Swagger 2.0 is a lot larger in scope than any of the previous versions and as such the documentation would require quite
a bit of cross-referencing but all of this would be done with plain strings in doc comments.
The opportunities for a misspelled identifier are plenty and the resulting JSON would still mostly look ok, so the errors would be fairly hard to spot.

Feel free to submit an issue if you disagree with any of this.

The parser supports filters to limit the packages to scan for rest operations
There are also filters for models so you can specify which packages
to include when scanning for models.
When a model has a filtered model as field then that filtered model will be
included transitively.

To use you can add a go:generate comment to your main file for example:

		// go:generate swagger generate spec .

The following annotations exist:

+swagger:meta

The +swagger:meta annotation flags a file as source for metadata about the API.
This is typically a doc.go file with your package documentation.

You can specify a Consumes and Produces key which has a new content type on each line
Schemes is a tag that is required and allows for a comma separated string composed of:
http, https, ws or wss

Host and BasePath can be specified but those values will be defaults,
they should get substituted when serving the swagger spec.

Default parameters and responses are not supported at this stage, for those you can edit the template json.

+swagger:strfmt [name]

A +swagger:strfmt annotation names a type as a string formatter. The name is mandatory and that is
what will be used as format name for this particular string format.
String formats should only be used for very well known formats.

+swagger:model [?model name]

A +swagger:model annotation optionally gets a model name as extra data on the line.
when this appears anywhere in a comment for a struct, then that struct becomes a schema
in the definitions object of swagger.

The struct gets analyzed and all the collected models are added to the tree.
The refs are tracked separately so that they can be renamed later on.

+swagger:route [method] [path pattern] [operation id] [?tag1 tag2 tag3]

A +swagger:route annotation links a path to a method.
This operation gets a unique id, which is used in various places as method name.
One such usage is in method names for client generation for example.

Because there are many routers available, this tool does not try to parse the paths
you provided to your routing library of choice. So you have to specify your path pattern
yourself in valid swagger syntax.

+swagger:params [operationid1 operationid2]

Links a struct to one or more operations. The params in the resulting swagger spec can be composed of several structs.
There are no guarantees given on how property name overlaps are resolved when several structs apply to the same operation.
This tag works very similar to the swagger:model tag except that it produces valid parameter objects instead of schema
objects.

+swagger:response [?response name]

Reads a struct decorated with +swagger:response and uses that information to fill up the headers and the schema for a response.
A +swagger:route can specify a response name for a status code and then the matching response will be used for that operation in the swagger definition.
*/
package parse
