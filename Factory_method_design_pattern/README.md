## Factory Method Design Pattern

Abstract factory design pattern hides out the complex implementations providing classes that can return required implementations. It is not just useful for end users also makes developers life easier by having the structure code implementations. 

Lets fist understand the terms we are going to use

### Factory
- It is a Suppurate Entity that has its own set of functionalities, Principles and use cases. It is consider to have different implementations for different factories. A factory code is always Abstracted from the client and made easy by defining abstract functions in Abstract Interface. 

### Concrete class
- the entry of the design structure, returns the factory methods according to the end user needs. Hides all the complexity. In here we define set of rules so that user gets desired output without any higher Archie understanding of entire code structure.

**Example** - Imagine the concept class like a waiter in the Hotels. Waiter Serves food by taking orders without involving you in the complexity of searching for it.

### Abstract Factory  
- It has all the abstract methods (a function without body) that every factory class must override. It gives commonality between factory classes so that if any other factory class are created in the future so there is a way to match the functionalities.

**Example** - We can consider it like the chef rooms. One chef room prepare cheese products and anther prepare chickens. where in there might be different sections for different kind of food items but at the end they must return the food that a waiter can serve. If it is burger waiter might have different arrangements no matter the kind of berger same case for pizza. Preparing food is common functionality over here. 

Abstract factory defines set of functionalities ( preparing food ) that every factory class ( Chef sections ) must do , programmatically it is done by extending the class and overriding the abstract methods.


### Factory class 
- there can be different factory classes each have there own way of implementing Abstract factory methods. Every factory classes returns the set of product classes.

**Example** - Let say the hotel has different sections for different types of burgers or pizzas ( Cheese Burger, Chicken Burger, ....). We can consider the place where cheese products are prepared as a different section (cheese Factory class) and chicken products ( chicken Factory class).At the end Every section must return a burger with two breads. This what the factory class defines there can be different implementations but returns same thing that abstract factory expects. The abstract factory expects Abstract product classes.



### Abstract Products 
- This are the interface that defines the structure of product class. Every product class must implement this abstract product interfaces and define the functions defined in it.

**Example** Let say it may be cheese burger or a chicken burger they implement the same burger structure. Same case for piza wether it is cheese one or chickens they must implement the pizza structure.

### Product classes : 
- this is where actual logic present. It is like abstract product interfaces defines the structure and product classes actually implement that.


### All Together
If we go all together from the customer to the chef. First customer orders the food to the waiter( client calling the concrete class) they waiter decides route the order to cheese or chicken factories ( concrete class has set of rules that is where request routing to the appropriate factory class is defined) . As request reaches the cheese factory or chicken factory, it is then decide what to make  burger or pizzas, if it is a burger then burger structure needs to be implemented  or if it is pizza , pizza structure needs to be implemented. So every factory returns products, that products must be implemented using the product interfaces. One important thing is that where it is a cheese factory or a chicken factory they must have burger and pizza if not the request is denied by the waiter ( every factory class must have the methods defined in the abstract class, also the products returned by the factory class methods must implement the abstract product interfaces so that there can be commonality between different factory classes, every method in the factory class returns a product class that product class implement the abstract product interface). To be more clear each product class has a abstract product interface. as if in the case of our example pizza and burger are the product classes and the factory decode weather to make cheese or chicken. So factory decides the logic and product classes must be implementing Abstract product interfaces.

Client → Concrete Factory (Waiter) → Factory Class (Chef Section) → Abstract Product (Food Structure) → Product (Actual Dish)
### Lets us have a small real time uncase to understand **Factory Method Design Pattern**

Let say we have two Different OS **Mac** and **windows** they decided to show the UI on the screen to the users. Obviously Windows has its own way of presenting Images and Battens and mac has its own way. To have a synchronization, Implementation needs to be same. Such that if a person have to see a webpage they need not to interact in different ways in different operating systems right. To maintain this we are now going to look at how Factory method design pattern solves this problem.

![Alt Text](assets/Factory.png "Architecture")

### Concrete Class (ConcretClass.go)

Lets us consider main.go has a Client, FMDP concept is very simple it just exposes the concert class which return the actual class according to the user requirements. This ConcretClass in our case is responsible for checking what particular OS the Client is using and based on that it return us a factory.

### Abstract Factory ( interfaces/FactoryInterface.go)

Abstract Factory lets client understand what all Functions we can perform using the Factory Class returned by the Concrete Class. It contain the set of predefined abstract functions, that must be performed by the Factory Class. In our case our theme is all about rendering ui and functions are RenderImage and RenderBatten there can be more but for simplicity i have defined this two. tomorrow a new factory comes like ubuntu or any other they must look at the abstract Factory interface and implement those functions. Here the work of developers became simple ans also the Clients.

### Factory Class (windows/WindowsFactory.go, mac/MacFactory.go)

Factory Classes is the place where all the functions that are defined in the Abstract Factory are implemented. In our case we have two factory's Windows and Mac. In GoLang we can use struct as a class. So now we have WindowsFactory struct and MacFactory struct. Both of this have RenderImage and RenderBatten functions.

### Product Interfaces (interfaces/ProductInterfaces.go)

As we just discussed weather it is windows or mac a batten must click and image must view. So Product interfaces defines that structure. For simplicity we just made the output as string but in real scenarios there can be something like onClick, hover, color, height, width etc.. and images interface is different from batten interface. product interface gives a set of functions that needs to be implemented by products(Battens, Images etc..)   

### Concrete Factory Products

If you look at the Return types of the abstract functions defined in the Factory Class, this are factory products interfaces which is actually implemented by product classes. Everything Function that is defined under Product Interface must be implemented by Factory Products. In case of Windows we have WindowsImage and WindowsBatten returns product interfaces like Images and Battens. 

```go
func main() {
	WindowsUI := UIrendering("windows")
	MacUI := UIrendering("Mac")
	fmt.Println(WindowsUI.RenderImage().Render())
	fmt.Println(WindowsUI.RenderBatten().Render())
	fmt.Println(MacUI.RenderImage().Render())
	fmt.Println(MacUI.RenderBatten().Render())
}
```
The flow is like UIrendering is a concrete class that returns a Factory interface (return Type ui), this UI has set of Functions like RenderImage and RenderBatten. Render gives factory specific implication.  


