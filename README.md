# goDesignPatterns
go设计模式的学习

### 1. Creational Patterns（创建型模式）: 这些模式与对象创建机制有关，目的是提高创建对象的灵活性和可扩展性。

Singleton（单例模式）</br>
: 确保一个类只有一个实例，并提供一个全局访问点

Simple Factory（简单工厂模式）</br>
: 一个工厂类负责创建所有产品，通过条件判断决定创建哪种产品

Factory Method（工厂方法模式）</br>
: 每个产品对应一个工厂类，符合开闭原则

Abstract Factory（抽象工厂模式）</br>
: 每个工厂类可以创建一组相关产品，强调产品族的概念

Builder（建造者模式）</br>
: 它用于分步构建复杂对象。建造者模式的核心思想是将一个复杂对象的构建过程与其表示分离，使得同样的构建过程可以创建不同的表示。

Prototype（原型模式）</br>
: 它通过复制现有对象来创建新对象，而不是通过新建类的方式。原型模式的核心思想是利用对象的克隆能力，避免重复初始化，特别适用于创建成本较高的对象。

: 使用原型模式，如果有引用类型，则需要考虑深拷贝和浅拷贝的问题
浅拷贝只复制对象本身而不复制其引用的对象，深拷贝则会递归地复制整个对象图。
这需要根据需求选择适当的拷贝方式

### 2. Structural Patterns（结构型模式）: 这些模式关注类或对象如何组成更大的结构。

Adapter（适配器模式）</br>
: 

Bridge（桥接模式）</br>
: 

Composite（组合模式）</br>
: 

Decorator（装饰者模式）</br>
: 

Facade（外观模式）</br>
: 

Flyweight（享元模式）</br>
: 

Proxy（代理模式）</br>
: 它通过提供一个代理对象来控制对另一个对象的访问。代理模式的核心思想是在不改变原始对象的情况下，通过代理对象来增强或限制对原始对象的访问。

### 3. Behavioral Patterns（行为型模式）: 这些模式与对象之间的交互和职责划分相关。

Chain of Responsibility（职责链模式）</br>
Command（命令模式）</br>
Interpreter（解释器模式）</br>
Iterator（迭代器模式）</br>
Mediator（中介者模式）</br>
Memento（备忘录模式）</br>
Observer（观察者模式）</br>
State（状态模式）</br>
Strategy（策略模式）</br>
Template Method（模板方法模式）</br>
Visitor（访问者模式）</br>
