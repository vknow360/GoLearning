package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://user:password@www.example.com:8080/path?query=123#fragment"

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	rawUrl1WithQuery := "https://www.example.com/search?q=golang&sort=asc"
	parsedUrl1, err := url.Parse(rawUrl1WithQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	queryParams := parsedUrl1.Query()
	fmt.Println("Query Parameters:")
	for key, values := range queryParams {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	//Building URL from components
	urlComponents := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/path",
		RawQuery: "query=456",
		Fragment: "section1",
	}
	builtUrl := urlComponents.String()
	fmt.Println("Built URL:", builtUrl)

}
