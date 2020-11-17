# weakdi

## from 
If use golang develop web, you maybe use martini. Unfortunately, this frame isn't used generally. Because DI is really really really adapt in golang.

We can use DI do something but only something.This is why this library named weak-di.

## How
```
type infos struct {
    job: string
}

func main() {
    s := di.MakeStorage()
    s.Set("info", func(name string, detail infos) {
        fmt.Printf(strings.Join([]string{name,":",detail.job}))
    })
    s.Provide("qwe")
    s.ProvideType(infos{"student"}, (*infos)(nil))
    s.Invoke("info")
}
```