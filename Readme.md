## Creational : patterns in creations of objects
#### Singleton :
Sometimes we need to have only one instance of our class for example a single DB connection shared by multiple objects as creating a separate DB connection for every object may be costly. Similarly, there can be a single configuration manager or error manager in an application that handles all problems instead of creating multiple managers

#### Builder : 
Here we are dealing with repetitive building algorithms which is same for creation of a set of objects.
The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic they require. 


#### Factory : 
In Factory pattern, we create object without exposing the creation logic to client and the client use the same common interface to create new type of object.<br>
With the Factory method, we delegate the creation of families of objects to a different package or object to abstract us from the knowledge of the pool of possible objects we could use.
t- Gives particular kind of object implementing same interface based on user requirement. Here we are simplifying the creation of many types of objects.
<br>t- ex- sql drivers where sql.open gets the required vendor db object based on driver.

#### Abstract Factory :

#### Prototype : 
use an already created instance of some type to clone it and complete it with the particular needs of each context.
The aim of the Prototype pattern is to have an object or a set of objects that is already created at compilation time, but which you can clone as many times as you want at runtime.

## Structural : patterns related to shape commonly used structures and relationships
#### Composite : 
Forming hierarchies and trees of objects. Objects have different objects with their own fields and methods inside them. This approach is very powerful and solves many problems of inheritance and multiple inheritances.
Golang encourages use of composition almost exclusively by its lack of inheritance
<br>t- objects, methods, fields inside objects instead of inheriting the other type.
We explicitly declare other types inside types so that its behaviour can be used.

####  Adapter : 
Ex - an interface gets outdated and it's not possible to replace it easily or fast. Instead, you create a new interface to deal with the current needs of your application, which, under the hood, uses implementations of the old interface.<br>
Instead of modifying your old source code (something which could not be possible in some situations), you have created a way to use the old functionality with a new signature.
<br>t- |z - OR adding new features on top of old one.
Can it be used to add features in third party libraries which we can't modify???

####  Bridge : 
change what an object does as much as we want. It also allows us to change the abstracted object while reusing the same implementation.
<br>t- change what the actual object was doing.

#### Proxy : 
The Proxy pattern usually wraps an object to hide some of its characteristics. These characteristics could be the fact that it is a remote object (remote proxy), a very heavy object.
<br>t- Hiding the heavy object with wrapping it in light proxy object.

#### Decorator :
The Decorator design pattern allows you to decorate an already existing type with more functional features without actually touching it.

#### Facade :
You use Facade when you want to hide the complexity of some tasks, especially when most of them share utilities (such as authentication in an API. t- ex. we can create a func with user n pass and the func will hide user from actual implementation of authentication). A library is a form of facade, where someone has to provide some methods for a developer to do certain things in a friendly way

#### Flyweight: 
Flyweight is a pattern which allows sharing the state of a heavy object between many instances of some type. Thanks to the Flyweight pattern, we can share all possible states of objects in a single common object, and thus minimize object creation by using pointers to already created objects.
<br>So, if we face a million users trying to access information about a match, we will actually just have two teams in memory with a million pointers to the same memory direction.

## Behavioural Patterns : encapsulate behaviors. t-About the methods patterns

#### Strategy : 
All types achieve the same functionality in a different way but the client of the strategy isn't affected.
<br>t- focus on changing algorithms based on requirement. Different types implement same interface methods. same interface object holds the particular type object and call methods of the interface.(Z-looks same like factory)

#### Chain Of Responsibility :  
t- chaining the calls of method of multiple object by just calling the one object method.
The 1st object holds ref to a object. In the 1st object method we call the referenced object method and the chain can be go on like that. ex- logging to multiple logging format.

#### Command :
#### Template :
#### Momento :
#### Interpreter :
#### Visitor :
This design tries to separate the logic needed to work with a specific object outside of the object itself. So we could have many different visitors that do some things to specific types.<br>
For example, imagine that we have a log writer that writes to console. We could make the logger "visitable" so that you can prepend any text to each log

#### State : 
in very simple terms, is something that has one or more states and travels between them to execute some behaviors.
We have a State interface and an implementation of each state we want to achieve. There is also usually a context that holds cross-information between the states.
t- -|z - almost similar with chain of responsibility.


#### Mediator :
#### Observer :
Observer pattern, also known as publish/subscriber or publish/listener.
Observer pattern is simple--to subscribe to some event that will trigger some behavior on many subscribed types.
The Observer pattern is especially useful to achieve many actions that are triggered on one event.

t- publisher has all the observer interface objects and it notifies each observer by calling its notify method. Subscribers will implement the observer interface and add itself to subscriber observer list.


<br><br>
book - https://learning.oreilly.com/library/view/go-design-patterns/9781786466204/ <br>
TODO:
Try adding main method in each pattern.
Fix code navigation in IDE and for github

