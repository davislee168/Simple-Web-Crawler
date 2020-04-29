WebCrawler structure (single process or distributed)
1.	Using single process for now.
2.	Build an engine() function which takes seed (url) as input.
3.	Fetch() function is to take seed’s URl as input and returns body content.
4.	ParseUrlListAndContent() function is to take body content from Fetch() function and returns 
requested number of URLs and content.  Export each link with associated content to text files.

