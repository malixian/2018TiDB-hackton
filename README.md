Description:
    《TiDB任务执行计划与监测可视化平台》，实现三大功能：TiDB执行计划可视化、TiDB执行任务监测、TiDB执行任务订制。解决三大问题：集群执行状态获取难，运维故障检测难，管理调优判断难。

backend: 
    技术栈：golang
    
    功能:包含collect-server与client两个部分，client负责收集TiDB中执行sql的任务执行信息与监控信息，并上传到server。server暴漏出HTTP接口以供查询与展示。
    
    Designed by github.com/malixian

frontend：
    技术栈：vue
    
    功能：可以从输入框中输入sql语句，与server进行HTTP交互，通过返回的json数据，生成任务执行计划树图、监控信息图。
    
    Designed by github.com/ZoeShaw101
