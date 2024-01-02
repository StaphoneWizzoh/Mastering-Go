_Generics_

They allow functions to process several data types without the need to write special code.

The `any` keyword tells the compiler that there are no constraints about a particular data type.

`comparable` constraint is just a predefined interface that includes all data types that can be compared with == or !=. This limits the use of only acceptable data types for a particular function.

Slices cannot be compared directly, meaning that additional implementation should be done.

Generics are a better option to use if one has to support multiple data types, than using interfaces.

Generics sacrifices speed for flexibility.
