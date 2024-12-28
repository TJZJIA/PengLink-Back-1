1. 四级管理
2. 获取数据
3. 数据展示
4. 设备管理

## 记住

"context"：

这个包提供了一个可以在goroutine之间传递请求范围的值、取消信号和截止时间的机制。它用于控制goroutine的生命周期，比如取消一个正在进行的操作。
"fmt"：

这是Go语言的标准库中的格式化I/O包，用于输入和输出，包括打印到控制台和字符串格式化。
"time"：

这个包提供了基本的时间功能，包括当前时间、时间的解析、时间的格式化和解析等。
"go.mongodb.org/mongo-driver/mongo"：

这是MongoDB的官方Go语言驱动程序，用于与MongoDB数据库进行交互。通过这个包，你可以连接到MongoDB数据库，执行CRUD操作（创建、读取、更新、删除）等。
"go.mongodb.org/mongo-driver/mongo/options"：

这个包提供了设置MongoDB客户端和集合操作的配置选项。它允许你自定义连接设置、读取和写入选项等。
var DB *mongo.Database 这行代码声明了一个全局变量DB，它是一个指向mongo.Database类型的指针，用于存储MongoDB数据库的引用，以便在程序的其他部分进行数据库操作。这个变量被声明但未初始化，通常在程序的初始化阶段，你会使用连接字符串和可能的配置选项来初始化这个变量。