## 简述
对wukongIM进行必要的二开，适配我们的数据单独存储的需求

1. 把对pebbledb的读写抽到单独的服务中，通过rest api调用。
2. 通过环境变量配置 数据服务地址 WK_ADAPTER_HOST,
    > 例如 http://localhost:8080
3. TODO 新增环境变量，用来控制 adapter 是否执行。
4. TODO 新增环境变量，用来控制 apdater 的返回值是否真正生效。
