# Topic: *Creational Design Patterns*
### Author: *Vladimir Leadavschi*
------
## Objectives :
__1. Study the Creational Design Patterns__

__2. Implement them in real projects__

## Theory :

Creational patterns provide object creation mechanisms that increase flexibility and reuse of existing code.

Some examples of this kind of design patterns are :

   * Singleton
   * Builder
   * Prototype
   * Object Pooling
   * Factory Method
   * Abstract Factory
   
## Implementation :
   * Creational (lab 1.Short description:)
    During the first laboratory work I've implemented only three creational design pattern which was Factory, Builder and prototype;
    As a theme I chose something like a custom fighterBuilder )))
    
    
 Builder:
      Creates objects of diferent types based on object name;
      
      //Here I create builders of type king and magician based on the input string;
      kingBuilder := builder.GetBuilder("king")
    	
      magicianBuilder := builder.GetBuilder("magician")
      
      
      //Here our players are in instantiation process:
      newPlayerKing := builder.NewPlayer(kingBuilder)
	    
      newPlayerMagician := builder.NewPlayer(magicianBuilder)
      
      
      //Here we build our proper object:
       kingFighter := newPlayerKing.BuildFighter("Arthur", "King", "sword")
	      
       magicianFighter := newPlayerMagician.BuildFighter("Gearalt", "Magician", "theurgy")
      
      
      
 Factory:
    My factoy method is implemented also in builder, because every fighter(player) have a specific weapon, which is also dinamycally generated based on input                   
    which we provide for it.
    
    //Implementation in Builder:
        func (k *King) GetWeapon(w string) (err error) {
	          k.Weapon, err = factory.GetWeapon(w)
	          return nil
        }
    /// Factory of weapons here
      func GetWeapon(weaponType string) (IWeapon, error) {
	        if weaponType == "sword" {
		          return NewSword(), nil
	         }

	        if weaponType == "spear" {
		          return NewSpear(), nil
	         }

	        if weaponType == "theurgy" {
		          return NewTheurgy(), nil
	        }

	        if weaponType == "dagger" {
		          return NewDagger(), nil
	        }

	        if weaponType == "bowandarrow" {
		        return NewBowAndArrow(), nil
	        }

	        return nil, fmt.Errorf("wrong weapon type")
      }
      
 
 
 Prototype:
     Also is a part of builder( The composite object (player) which is created after we instatiate a builder:
        
        func (f *Fighter) Clone() prototype.ClonePrototyper {
	          return &Fighter{
                  Name: f.Name + "_clone",
		              Role:   f.Role + "_clone",
		              Weapon: f.Weapon,
	          }
        }
    
 After we specify character we wanna create, our playerBuilder builds this type of object(which is of type Fighter) fighter implements method Clone, which mean           
 which mean that every simple object also extend method clone.
           
      
        
    
