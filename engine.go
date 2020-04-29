Package engine

Import (
       “fetcher”
       “crawl”
       “log”
)

func Run(seed string, linkLimit int) {
       if len(seed) > 0 {
              log.Printf(“Fetching %s”, seed)
       
              body, err := fetcher.Fectch(seed)
              if err != nil {
                     log.Printf(“Fetch error: url %s,%v”, seed, err)
                     log.Fatal(err)
              }
              err = crawl.ParseUrlListAndContent(body, linkLimit)
              if err != nil {
       	       log.Fatal(err)
              }
       }
}
  
