今天没有去继续更新开源10年的项目，因为昨天接触到几个新的知识点，所以今天加强一下对他们的认识，下面是本节的一个知识点。

1）Gin Web框架的认识

2）Gin加载静态资源

3）Gin加载动态资源

在学习Gin的过程中动手搞了一个小Demo，把Gin的简单用发都跑了一下，强烈建议各位读者不要只看文章，自己动手写一下效果更佳。

项目在GitHub上的地址：https://github.com/kaiyuan10nian/GinDemo

在开始这一节前，希望有兴趣的同学可以先去看下这个开源10年项目前面的故事，虽不会直接提升你的开发能力，但看后确实能了解这个项目的灵魂。

1、开源10年项目-项目成立，欢迎大家的加入

2、开源10年项目-准备工作，语言的选择和第一阶段的目标

3、开源10年项目-Go项目实战之博客功能开发

下面开始我们今天的知识点分享。

一、创建一个Go web项目，这里我命名为：GinDemo，方便我们的学习。

项目结构
![image](https://user-images.githubusercontent.com/101968429/159395580-1babec22-7416-4f89-84de-b308a44e8f17.png)


Gin安装：go get -u http://github.com/gin-gonic/gin

直接在Goland下面的Terminal中输入就可以了。看到下图就表示你安装成功了
![image](https://user-images.githubusercontent.com/101968429/159395764-6a91459a-6a9b-4a44-b7d0-53fb0378df4f.png)


二、先写一个小案例，并在Postman或者浏览器中打开看下效果。

在GinDemo下面创建一个main.go文件，输入以下内容：
![image](https://user-images.githubusercontent.com/101968429/159395782-c3ee62b5-8add-4ee0-a84b-e8c07a1013cb.png)


然后在下面的Termonal中输入以下指令运行该项目：

Go run main.go

看到下图即表示你已经运行起来了。
![image](https://user-images.githubusercontent.com/101968429/159395792-6da0659a-05bd-4ab7-b7dd-0446c4e07a85.png)


这个时候，打开你的浏览器，输入：localhost:8080/ping,将显示以下内容：
![image](https://user-images.githubusercontent.com/101968429/159395815-f3b73898-7598-40aa-8c25-4275af3d5eee.png)


整个Demo中的注释还是比较清楚的，每一行代码是什么意思，有什么作用等都比较简单，运行到这里基本上Gin的精髓就已经学到了。

下面两个知识点是可以解决我们在真正的实际项目中经常会遇到的需求的。所以我这里单独拉了出来写一下。当然，它还有其它别的更多用法，我们就不一一说了，至少掌握了这两个对于普通的开发工作就足以应对了。

三、加载静态资源

在实际项目中我们经常会用到很多静态资源，比如：图片、文件等。那Gin是怎么处理的呢？还是看案例，下面我们对上面的main.go进行一下简单修改：
![image](https://user-images.githubusercontent.com/101968429/159395861-8b9d61b7-16d0-40b3-a9d4-7986fafaa0ef.png)


主要加了两行代码，用到两个函数。注释中已经描述的很清楚了，说一下代码中未描述的内容，非常重要，这两个函数的第一个参数就是相对地址，也就是说是用户端访问的时候访问的地址，第二个参数是本地服务器的地址，也就是引用的地址。

比如我们运行上面项目后，若想访问/Users/fu/GolandProjects/GinDemo/web/static里面的图片，那么直接如下操作即可：
![image](https://user-images.githubusercontent.com/101968429/159395888-5503501e-ef89-4c5f-86fc-ac6a55055603.png)


若想加载第二种加载形式的，则直接这么访问：
![image](https://user-images.githubusercontent.com/101968429/159395915-88c4840c-177e-4271-bfe8-d1d0b890c0a5.png)


其实，这两种访问形式访问的都是同一张图片。

这就是静态加载资源的方式方法，那么我们做一个web站点肯定不是仅仅有静态资源，肯定还有动态资源，动态资源的加载怎么实现？

四、加载动态资源

首先，在GinDemo-web下新建templete并在其中新增一个index.html文件，文件内容很简单，这里制作演示所以就没有去做接口互动。
![image](https://user-images.githubusercontent.com/101968429/159396125-35955cc2-4e84-4f08-978d-3cd985c06c3a.png)


然后，在根目录下创建了controller文件夹并创建con.go，主要用来存放逻辑层的操作，受JAVA开发的影响我这么去做了，你可以随意哈，怎么高兴怎么来。

![image](https://user-images.githubusercontent.com/101968429/159396145-cdfabc18-e82d-4906-9156-a69304f87368.png)

还是在main.go中做修改,并加载trmplete下面的资源，然后做了一个web分组，分组中仅有一个接口，其处理放到了con.go文件中的IndexController函数中。
![image](https://user-images.githubusercontent.com/101968429/159396166-05efc486-359d-4fbf-8fe2-c4fbc6a0a222.png)


直接运行，然后浏览器中访问index.html看看有什么效果？
![image](https://user-images.githubusercontent.com/101968429/159396179-4917b31c-28de-4f01-b029-ea5f87cf400d.png)


到这里，其实这个项目就算完成了，其中涉及到的知识点我们也都了解了，就这些内容已经完全够现阶段的我们使用了。

拜拜。。。See you next.

开源十年项目的更新首发于公众号：计算机自学平台，有兴趣的小伙伴可以持续关注，并欢迎各位加我的微信跟我一起完成并推动项目的发展。
