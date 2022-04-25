# tran翻译使用说明

## 1.start

打开main.exe

![image-20220425160717797](C:%5CUsers%5Cwindows%5CAppData%5CRoaming%5CTypora%5Ctypora-user-images%5Cimage-20220425160717797.png)

## 2.复制地址

输入需要翻译文件夹的最外层

如：我们需要翻译数据库DMKD

![image-20220425160810301](C:%5CUsers%5Cwindows%5CAppData%5CRoaming%5CTypora%5Ctypora-user-images%5Cimage-20220425160810301.png)

我们只需要打开这个文件夹复制**D:\GoProject\tran\数据库DMKD**![image-20220425160834433](C:%5CUsers%5Cwindows%5CAppData%5CRoaming%5CTypora%5Ctypora-user-images%5Cimage-20220425160834433.png)

## 3.输入

输入这个地址到main.exe中，按下回车即可开始。

![image-20220425160931612](C:%5CUsers%5Cwindows%5CAppData%5CRoaming%5CTypora%5Ctypora-user-images%5Cimage-20220425160931612.png)

## 4.查看反馈数据

查看反馈数据，哪些文件夹已经开始翻译

![image-20220425161023259](C:%5CUsers%5Cwindows%5CAppData%5CRoaming%5CTypora%5Ctypora-user-images%5Cimage-20220425161023259.png)

生成的默认文件夹为文件夹名字加上new

比如这里生成的

![image-20220425161205907](C:%5CUsers%5Cwindows%5CAppData%5CRoaming%5CTypora%5Ctypora-user-images%5Cimage-20220425161205907.png)

## 错误处理

当main.exe文件中出现红色的文字，就关闭main.exe这个程序，然后删除每个**新建翻译目录**中最后一个翻译的文件，再重新打开即可（毕竟不是所有人电脑上都有redis来标记，直接续传。。这样是最方便的处理了）

比如如果中途关闭或者出现错误了，就给上面示例中2013new，2014new......中这些新生成的文件夹的最后一个生成的文件删除（**注意按修改日期排序**，右键文件夹可以调整排序方式）

比如这里就是删除，箭头指向这个文件。

![image-20220425161550744](C:%5CUsers%5Cwindows%5CAppData%5CRoaming%5CTypora%5Ctypora-user-images%5Cimage-20220425161550744.png)

删除后重新打开main.exe，输入地址即可（很少出现报错的情况，注意翻译期间，别断网，别关闭就行。如果关闭就需要进行错误处理的操作）