评论回复 , 昵称(管理员) 不能使用, - ok
评论回复邮件发送        ok


评论分页 - ok

404页面 - ok

sql优化,连接池 - ok

错误日志 - ok

首页友链 - 最后做|不做


静态页面        v.1.0.1

404页面优化     v.1.0.1

内存泄露 - ok


    后台管理
        登录 - 账号密码写死  ok
        文章 - 增删改       ok
        评论 - 改,回复 ok
        友链管理 - 待定 v.1.0.1
    
    系统
        重启       v.1.0.1
        配置文件    - ok
        中间件  v.1.0.1

    
    
代码整理

daemonize

    Golang 在 Mac、Linux、Windows 下如何交叉编译
        **Mac 下编译 Linux 和 Windows 64位可执行程序**
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
        CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
        
        **Linux 下编译 Mac 和 Windows 64位可执行程序**
        CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
        CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

        **Windows 下编译 Mac 和 Linux 64位可执行程序**
        SET CGO_ENABLED=0
        SET GOOS=darwin
        SET GOARCH=amd64
        go build main.go
        
        SET CGO_ENABLED=0
        SET GOOS=linux
        SET GOARCH=amd64
        go build main.go
        
        GOOS：目标平台的操作系统（darwin、freebsd、linux、windows） 
        GOARCH：目标平台的体系架构（386、amd64、arm） 
        交叉编译不支持 CGO 所以要禁用它
        
        上面的命令编译 64 位可执行程序，你当然应该也会使用 386 编译 32 位可执行程序 
        很多博客都提到要先增加对其它平台的支持，但是我跳过那一步，上面所列的命令也都能成功，
        且得到我想要的结果，可见那一步应该是非必须的，或是我所使用的 Go 版本已默认支持所有平台。

