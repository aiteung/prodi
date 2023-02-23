# prodi
kepentingan borang akreditasi prodi


## Release A Package

Set Environtment variabel:

GOPROXY=proxy.golang.org

commit and push your code first

```sh
git tag                                 #check current version
git tag v0.0.3                          #set tag version
git push origin --tags                  #push tag version to repo
go list -m github.com/ORG/REPO@v0.0.3   #publish to pkg dev, replace ORG/URL with your repo URL
```