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
: 它允许不兼容的接口之间进行协作。适配器模式的核心思想是将一个类的接口转换成客户端期望的另一个接口，从而使原本不兼容的类能够一起工作。

Bridge（桥接模式）</br>
: 它的核心思想是将抽象部分与实现部分分离，使它们可以独立变化。通过这种方式，桥接模式能够避免类的数量爆炸（即类的组合呈指数增长），同时提高代码的可扩展性和可维护性。

Composite（组合模式）</br>
: 它允许你将对象组合成树形结构来表示“部分-整体”的层次关系。组合模式让客户端可以统一地处理单个对象和对象组合。

Decorator（装饰者模式）</br>
: 它允许你动态地为对象添加行为或职责，而不需要修改对象的原始类。通过引入装饰者类，可以在运行时灵活地组合不同的功能，而不需要创建大量的子类。装饰者模式的核心思想是将对象包装在一个或多个装饰者中，每个装饰者都可以在调用被装饰对象的方法之前或之后添加额外的行为。

Facade（外观模式）</br>
: 它提供了一个统一的接口，用于访问子系统中的一组接口。外观模式的核心思想是简化复杂系统的使用，通过提供一个高层接口，隐藏系统的复杂性，使客户端更容易使用。

Flyweight（享元模式）</br>
: 它通过共享对象来减少内存使用和提高性能。享元模式的核心思想是将对象的共享部分（内部状态）与不可共享部分（外部状态）分离，从而减少重复对象的创建。

Proxy（代理模式）</br>
: 它通过提供一个代理对象来控制对另一个对象的访问。代理模式的核心思想是在不改变原始对象的情况下，通过代理对象来增强或限制对原始对象的访问。

### 3. Behavioral Patterns（行为型模式）: 这些模式与对象之间的交互和职责划分相关。

Chain of Responsibility（职责链模式）</br>
: 是一种行为设计模式，它允许多个对象有机会处理请求，从而避免请求的发送者与接收者之间的耦合。责任链模式将这些对象连成一条链，并沿着这条链传递请求，直到有对象处理它为止。

Command（命令模式）</br>
: 它将请求封装为一个对象，从而使你可以用不同的请求对客户进行参数化，并且支持请求的排队、记录日志、撤销操作等功能。

Interpreter（解释器模式）</br>
: 它定义了一种语言的语法表示，并提供了一个解释器来解释这种语法。解释器模式通常用于处理类似编程语言、查询语言、规则引擎等场景。

Iterator（迭代器模式）</br>
: 它提供了一种方法顺序访问一个聚合对象中的各个元素，而又不需要暴露该对象的内部表示。迭代器模式的核心思想是将遍历逻辑从聚合对象中分离出来，使得聚合对象和遍历逻辑可以独立变化。

Mediator（中介者模式）</br>
: 

Memento（备忘录模式）</br>
: 

Observer（观察者模式）</br>
: 

State（状态模式）</br>
: 

Strategy（策略模式）</br>
: 

Template Method（模板方法模式）</br>
: 

Visitor（访问者模式）</br>
: 

