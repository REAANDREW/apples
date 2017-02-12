# Something Continuous


## References

[https://ariejan.net/2015/10/03/a-makefile-for-golang-cli-tools/](https://ariejan.net/2015/10/03/a-makefile-for-golang-cli-tools/)

Followed the tutorial for the Makefile

## IDEA

Meta endpoints

```go
metaApi := router.Group("/_")
{
    infoApi := metaApi.Group("/info")
    {
        infoApi.GET("/application", func(ctx *gin.Context) {
            ctx.JSON(200, map[string]string{
                "Version":   Version,
                "BuildTime": BuildTime,
            })
        })

        hostInfo, err := host.Info()
        if err != nil {
            panic(ErrCannotGetHostInfo)
        }
        infoApi.GET("/host", func(ctx *gin.Context) {
            ctx.JSON(200, hostInfo)
        })
    }
}
```

