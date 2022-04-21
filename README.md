# test-cloudtrail
## Example 

1. Cloud Trail中所记录的所有userIdentity有哪些？各产生了多少次事件？ 请列举userIdentity的username，或invokedBy, 或principalId 以及相关事件次数。
```console
go run main.go principalId
```


2. Cloud Trail中所记录的有哪些AWS资源被创建？ 请列举相关资源的arn或id。 如 （arn:aws:elasticloadbalancing:us-west-1:479841282623:listener/app/pki-sld-lb-sdf/e2cecddb95a881d3/68e28fcbce005947， eni-09948011cdc73810a）
```console
go run main.go resource
```

