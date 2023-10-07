实现一个rpc服务:
1. 从文件/示例/数据库等数据源中获取特征数据集, 每个特征包含了一个经纬度位置和该位置的名称信息;
2. 服务支持获取特征(GetFeature)、列出所有特征(ListFeatures)、记录特征(RecordFeature)、路由聊天(RouteChat)四个方法.